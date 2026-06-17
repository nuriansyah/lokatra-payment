package paymentgateway

import "time"

type HTTPClientConfig struct {
	Timeout          time.Duration `json:"timeout"`
	RetryCount       int           `json:"retryCount"`
	RetryWaitTime    time.Duration `json:"retryWaitTime"`
	RetryMaxWaitTime time.Duration `json:"retryMaxWaitTime"`
	EnableDebug      bool          `json:"enableDebug"`
	UserAgent        string        `json:"userAgent"`
}

type ProviderConfig struct {
	Code            ProviderCode      `json:"code"`
	BaseURL         string            `json:"baseUrl"`
	APIKey          string            `json:"-"`
	SecretKey       string            `json:"-"`
	ServerKey       string            `json:"-"`
	ClientKey       string            `json:"-"`
	MerchantID      string            `json:"merchantId"`
	MerchantKey     string            `json:"-"`
	WebhookToken    string            `json:"-"`
	WebhookSecret   string            `json:"-"`
	Environment     string            `json:"environment"` // sandbox, production
	DefaultCurrency string            `json:"defaultCurrency"`
	HTTP            HTTPClientConfig  `json:"http"`
	Endpoints       map[string]string `json:"endpoints"`
	ExtraHeaders    map[string]string `json:"extraHeaders"`
	Extra           map[string]string `json:"extra"`
}

func (c ProviderConfig) Endpoint(key, fallback string) string {
	if c.Endpoints != nil {
		if v, ok := c.Endpoints[key]; ok && v != "" {
			return v
		}
	}
	return fallback
}

func (c ProviderConfig) Currency() string {
	if c.DefaultCurrency != "" {
		return c.DefaultCurrency
	}
	return "IDR"
}

type Config struct {
	HTTP      HTTPClientConfig                `json:"http"`
	Providers map[ProviderCode]ProviderConfig `json:"providers"`
}

func (c Config) Provider(code ProviderCode) (ProviderConfig, bool) {
	if c.Providers == nil {
		return ProviderConfig{}, false
	}
	p, ok := c.Providers[code]
	if !ok {
		return ProviderConfig{}, false
	}
	if p.Code == "" {
		p.Code = code
	}
	if p.HTTP.Timeout == 0 {
		p.HTTP = c.HTTP
	}
	return p, true
}
