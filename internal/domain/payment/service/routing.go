package service

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/rs/zerolog/log"

	paymentmodel "github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	paymentrepository "github.com/nuriansyah/lokatra-payment/internal/domain/payment/repository"
	routingmodel "github.com/nuriansyah/lokatra-payment/internal/domain/routing/model"
	routingrepository "github.com/nuriansyah/lokatra-payment/internal/domain/routing/repository"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

// RoutingEngine performs config-first resolution with database fallback.
type RoutingEngine struct {
	config          PaymentServiceConfig
	paymentRepo     paymentrepository.Repository
	routingRepo     routingrepository.Repository
	gatewayRegistry GatewayRegistry
}

func NewRoutingEngine(config PaymentServiceConfig, paymentRepo paymentrepository.Repository, routingRepo routingrepository.Repository, gatewayRegistry GatewayRegistry) *RoutingEngine {
	return &RoutingEngine{
		config:          config,
		paymentRepo:     paymentRepo,
		routingRepo:     routingRepo,
		gatewayRegistry: gatewayRegistry,
	}
}

func (re *RoutingEngine) Resolve(ctx context.Context, request PaymentFlowRequest) (RoutingDecision, error) {
	policyDecision, ok := re.resolveFromPolicies(request)
	if ok {
		return policyDecision, nil
	}

	if re.config.UseDatabaseFallback {
		decision, err := re.resolveFromDatabase(ctx, request)
		if err == nil {
			return decision, nil
		}
		log.Debug().Err(err).Msg("[routing] database fallback failed")
	}

	return re.resolveDefault(request)
}

func (re *RoutingEngine) resolveFromPolicies(request PaymentFlowRequest) (RoutingDecision, bool) {
	matching := make([]RoutingPolicy, 0, len(re.config.Policies))
	for _, policy := range re.config.Policies {
		if !policy.Enabled {
			continue
		}
		if policy.matches(request) {
			matching = append(matching, policy)
		}
	}

	if len(matching) == 0 {
		return RoutingDecision{}, false
	}

	sort.SliceStable(matching, func(i, j int) bool {
		if matching[i].Priority == matching[j].Priority {
			return matching[i].Name < matching[j].Name
		}
		return matching[i].Priority < matching[j].Priority
	})

	for _, policy := range matching {
		if decision, ok := re.buildPolicyDecision(policy, request); ok {
			return decision, true
		}
	}

	return RoutingDecision{}, false
}

func (re *RoutingEngine) buildPolicyDecision(policy RoutingPolicy, request PaymentFlowRequest) (RoutingDecision, bool) {
	candidates := re.collectCandidates(policy, request)
	if len(candidates) == 0 {
		return RoutingDecision{}, false
	}

	sort.SliceStable(candidates, func(i, j int) bool {
		if candidates[i].Score == candidates[j].Score {
			return candidates[i].AccountLabel < candidates[j].AccountLabel
		}
		return candidates[i].Score > candidates[j].Score
	})

	selected := candidates[0]
	return RoutingDecision{
		Strategy:        policy.strategyOrDefault(re.config.DefaultStrategy),
		PolicyName:      policy.Name,
		PSP:             selected.PSP,
		PSPAccountID:    selected.AccountID,
		PSPAccountLabel: selected.AccountLabel,
		Reason:          selected.Reason,
		Candidates:      candidates,
	}, true
}

func (re *RoutingEngine) collectCandidates(policy RoutingPolicy, request PaymentFlowRequest) []RoutingCandidate {
	candidates := make([]RoutingCandidate, 0, len(re.config.Providers))
	for _, provider := range re.config.Providers {
		if !provider.Enabled {
			continue
		}
		if !provider.supports(request.PaymentMethodType, request.Currency) {
			continue
		}

		score := 0
		reasons := make([]string, 0, 3)
		if policy.PreferredPSP == provider.PSP {
			score += 100
			reasons = append(reasons, "preferred psp")
		}
		if policy.PreferredAccount != uuid.Nil && policy.PreferredAccount == provider.AccountID {
			score += 80
			reasons = append(reasons, "preferred account")
		}
		if len(policy.Countries) > 0 && containsString(policy.Countries, strings.ToUpper(request.CustomerCountry)) {
			score += 10
			reasons = append(reasons, "country match")
		}
		if len(policy.Currencies) == 0 || containsCurrency(policy.Currencies, request.Currency) {
			score += 10
		}
		if len(policy.PaymentMethods) == 0 || containsPaymentMethod(policy.PaymentMethods, request.PaymentMethodType) {
			score += 10
		}
		if len(reasons) == 0 {
			reasons = append(reasons, "provider compatible")
		}

		candidates = append(candidates, RoutingCandidate{
			PSP:          provider.PSP,
			AccountID:    provider.AccountID,
			AccountLabel: provider.AccountLabel,
			Score:        score,
			Reason:       strings.Join(reasons, ", "),
		})
	}

	return candidates
}

func (re *RoutingEngine) resolveFromDatabase(ctx context.Context, request PaymentFlowRequest) (RoutingDecision, error) {
	if re.routingRepo == nil {
		return RoutingDecision{}, failure.Unimplemented("routing database fallback")
	}

	profiles := []routingmodel.RoutingProfilesFilterResult{}
	var err error

	profiles = append(profiles, re.fetchProfilesByMerchant(ctx, request.MerchantID)...)
	if len(profiles) == 0 {
		profiles = append(profiles, re.fetchGlobalProfiles(ctx)...)
	}

	if len(profiles) == 0 {
		return RoutingDecision{}, failure.NotFound("routing profile")
	}

	sort.SliceStable(profiles, func(i, j int) bool {
		return profiles[i].RoutingProfiles.MetaCreatedAt.After(profiles[j].RoutingProfiles.MetaCreatedAt)
	})

	for _, profile := range profiles {
		decision, found, ruleErr := re.resolveProfileRules(ctx, profile.RoutingProfiles, request)
		if ruleErr != nil {
			err = ruleErr
			continue
		}
		if found {
			return decision, nil
		}
	}

	if err != nil {
		return RoutingDecision{}, err
	}
	return RoutingDecision{}, failure.NotFound("routing rule")
}

func (re *RoutingEngine) fetchProfilesByMerchant(ctx context.Context, merchantID uuid.UUID) []routingmodel.RoutingProfilesFilterResult {
	filter := routingmodel.Filter{
		FilterFields: []routingmodel.FilterField{{Field: string(routingmodel.RoutingProfilesDBFieldName.MerchantId), Operator: routingmodel.OperatorEqual, Value: merchantID}, {Field: string(routingmodel.RoutingProfilesDBFieldName.IsActive), Operator: routingmodel.OperatorEqual, Value: true}},
		Sorts:        []routingmodel.Sort{{Field: string(routingmodel.RoutingProfilesDBFieldName.MetaCreatedAt), Order: routingmodel.SortDesc}},
		Pagination:   routingmodel.Pagination{Page: 1, PageSize: 10},
	}
	profiles, err := re.routingRepo.ResolveRoutingProfilesByFilter(ctx, filter)
	if err != nil {
		log.Debug().Err(err).Msg("[routing] merchant profile lookup failed")
		return nil
	}
	return profiles
}

func (re *RoutingEngine) fetchGlobalProfiles(ctx context.Context) []routingmodel.RoutingProfilesFilterResult {
	filter := routingmodel.Filter{
		FilterFields: []routingmodel.FilterField{{Field: string(routingmodel.RoutingProfilesDBFieldName.MerchantId), Operator: routingmodel.OperatorIsNull, Value: true}, {Field: string(routingmodel.RoutingProfilesDBFieldName.IsActive), Operator: routingmodel.OperatorEqual, Value: true}},
		Sorts:        []routingmodel.Sort{{Field: string(routingmodel.RoutingProfilesDBFieldName.MetaCreatedAt), Order: routingmodel.SortDesc}},
		Pagination:   routingmodel.Pagination{Page: 1, PageSize: 10},
	}
	profiles, err := re.routingRepo.ResolveRoutingProfilesByFilter(ctx, filter)
	if err != nil {
		log.Debug().Err(err).Msg("[routing] global profile lookup failed")
		return nil
	}
	return profiles
}

func (re *RoutingEngine) resolveProfileRules(ctx context.Context, profile routingmodel.RoutingProfiles, request PaymentFlowRequest) (RoutingDecision, bool, error) {
	filter := routingmodel.Filter{
		FilterFields: []routingmodel.FilterField{{Field: string(routingmodel.RoutingRulesDBFieldName.ProfileId), Operator: routingmodel.OperatorEqual, Value: profile.Id}, {Field: string(routingmodel.RoutingRulesDBFieldName.IsActive), Operator: routingmodel.OperatorEqual, Value: true}},
		Sorts:        []routingmodel.Sort{{Field: string(routingmodel.RoutingRulesDBFieldName.Priority), Order: routingmodel.SortAsc}},
		Pagination:   routingmodel.Pagination{Page: 1, PageSize: 100},
	}
	rules, err := re.routingRepo.ResolveRoutingRulesByFilter(ctx, filter)
	if err != nil {
		return RoutingDecision{}, false, err
	}

	for _, rule := range rules {
		if !ruleMatchesRequest(rule.RoutingRules, request) {
			continue
		}
		if rule.RoutingRules.MatchCardBin.Valid && request.CardBIN != "" && !strings.HasPrefix(request.CardBIN, rule.RoutingRules.MatchCardBin.String) {
			continue
		}

		candidatePolicy := RoutingPolicy{
			Name:             profile.Name,
			Enabled:          true,
			Strategy:         routingmodel.RoutingStrategy(profile.Strategy),
			PreferredPSP:     paymentmodel.Psp(""),
			DatabaseFallback: true,
		}
		candidates := re.collectCandidates(candidatePolicy, request)
		if len(candidates) == 0 {
			continue
		}

		sort.SliceStable(candidates, func(i, j int) bool {
			if candidates[i].Score == candidates[j].Score {
				return candidates[i].AccountLabel < candidates[j].AccountLabel
			}
			return candidates[i].Score > candidates[j].Score
		})

		selected := candidates[0]
		return RoutingDecision{
			Strategy:        routingmodel.RoutingStrategy(profile.Strategy),
			PolicyName:      profile.Name,
			ProfileID:       profile.Id,
			ProfileMatched:  true,
			RuleID:          rule.RoutingRules.Id,
			RuleMatched:     true,
			PSP:             selected.PSP,
			PSPAccountID:    selected.AccountID,
			PSPAccountLabel: selected.AccountLabel,
			Reason:          fmt.Sprintf("matched routing rule %s; %s", rule.RoutingRules.Name, selected.Reason),
			FallbackUsed:    true,
			Candidates:      candidates,
		}, true, nil
	}

	if profile.FallbackProfileId.Valid {
		fallbackFilter := routingmodel.Filter{
			FilterFields: []routingmodel.FilterField{{Field: string(routingmodel.RoutingProfilesDBFieldName.Id), Operator: routingmodel.OperatorEqual, Value: profile.FallbackProfileId.UUID}, {Field: string(routingmodel.RoutingProfilesDBFieldName.IsActive), Operator: routingmodel.OperatorEqual, Value: true}},
		}
		fallbackProfiles, err := re.routingRepo.ResolveRoutingProfilesByFilter(ctx, fallbackFilter)
		if err != nil {
			return RoutingDecision{}, false, err
		}
		for _, fallbackProfile := range fallbackProfiles {
			return re.resolveProfileRules(ctx, fallbackProfile.RoutingProfiles, request)
		}
	}

	return RoutingDecision{}, false, nil
}

func (re *RoutingEngine) resolveDefault(request PaymentFlowRequest) (RoutingDecision, error) {
	for _, provider := range re.config.Providers {
		if !provider.Enabled {
			continue
		}
		if provider.supports(request.PaymentMethodType, request.Currency) {
			return RoutingDecision{
				Strategy:        re.config.DefaultStrategy,
				PolicyName:      "default",
				PSP:             provider.PSP,
				PSPAccountID:    provider.AccountID,
				PSPAccountLabel: provider.AccountLabel,
				Reason:          "default provider selection",
				FallbackUsed:    true,
			}, nil
		}
	}

	return RoutingDecision{}, failure.NotFound("payment gateway")
}

func ruleMatchesRequest(rule routingmodel.RoutingRules, request PaymentFlowRequest) bool {
	if rule.MatchPaymentMethod != "" && rule.MatchPaymentMethod != routingmodel.PaymentMethodType(request.PaymentMethodType) {
		return false
	}
	if rule.MatchCurrency != "" && rule.MatchCurrency != routingmodel.PaymentCurrency(request.Currency) {
		return false
	}
	if rule.MatchAmountMin.Valid && request.Amount.LessThan(rule.MatchAmountMin.Decimal) {
		return false
	}
	if rule.MatchAmountMax.Valid && request.Amount.GreaterThan(rule.MatchAmountMax.Decimal) {
		return false
	}
	if rule.MatchUserCountry.Valid && !strings.EqualFold(rule.MatchUserCountry.String, request.CustomerCountry) {
		return false
	}
	if rule.MatchProductType.Valid && !strings.EqualFold(rule.MatchProductType.String, request.ProductType) {
		return false
	}
	return true
}

func containsString(values []string, target string) bool {
	for _, value := range values {
		if strings.EqualFold(value, target) {
			return true
		}
	}
	return false
}

func containsCurrency(values []paymentmodel.PaymentCurrency, target paymentmodel.PaymentCurrency) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func containsPaymentMethod(values []paymentmodel.PaymentMethodType, target paymentmodel.PaymentMethodType) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func (policy RoutingPolicy) matches(request PaymentFlowRequest) bool {
	if len(policy.UseCases) > 0 && !containsUseCase(policy.UseCases, request.UseCase) {
		return false
	}
	if len(policy.MerchantIDs) > 0 && !containsUUID(policy.MerchantIDs, request.MerchantID) {
		return false
	}
	if len(policy.OrderTypes) > 0 && !containsString(policy.OrderTypes, request.OrderType) {
		return false
	}
	if len(policy.ProductTypes) > 0 && !containsString(policy.ProductTypes, request.ProductType) {
		return false
	}
	if len(policy.PaymentMethods) > 0 && !containsPaymentMethod(policy.PaymentMethods, request.PaymentMethodType) {
		return false
	}
	if len(policy.Currencies) > 0 && !containsCurrency(policy.Currencies, request.Currency) {
		return false
	}
	if len(policy.Countries) > 0 && !containsString(policy.Countries, request.CustomerCountry) {
		return false
	}
	if policy.MinAmount != nil && request.Amount.LessThan(*policy.MinAmount) {
		return false
	}
	if policy.MaxAmount != nil && request.Amount.GreaterThan(*policy.MaxAmount) {
		return false
	}
	return true
}

func (policy RoutingPolicy) strategyOrDefault(defaultStrategy routingmodel.RoutingStrategy) routingmodel.RoutingStrategy {
	if policy.Strategy != "" {
		return policy.Strategy
	}
	return defaultStrategy
}

func containsUseCase(values []PaymentUseCase, target PaymentUseCase) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func containsUUID(values []uuid.UUID, target uuid.UUID) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}
