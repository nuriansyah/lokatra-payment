package paymentgateway

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

func NewRestyClient(cfg ProviderConfig) *resty.Client {
	h := cfg.HTTP
	if h.Timeout == 0 {
		h.Timeout = 15 * time.Second
	}
	if h.RetryWaitTime == 0 {
		h.RetryWaitTime = 200 * time.Millisecond
	}
	if h.RetryMaxWaitTime == 0 {
		h.RetryMaxWaitTime = 2 * time.Second
	}
	if h.UserAgent == "" {
		h.UserAgent = "lokatra-payment/1.0"
	}

	r := resty.New().
		SetBaseURL(strings.TrimRight(cfg.BaseURL, "/")).
		SetTimeout(h.Timeout).
		SetRetryCount(h.RetryCount).
		SetRetryWaitTime(h.RetryWaitTime).
		SetRetryMaxWaitTime(h.RetryMaxWaitTime).
		SetHeader("User-Agent", h.UserAgent).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	for k, v := range cfg.ExtraHeaders {
		r.SetHeader(k, v)
	}
	if h.EnableDebug {
		r.SetDebug(true)
	}
	r.AddRetryCondition(func(resp *resty.Response, err error) bool {
		if err != nil {
			return true
		}
		if resp == nil {
			return true
		}
		code := resp.StatusCode()
		return code == http.StatusRequestTimeout || code == http.StatusTooManyRequests || code >= 500
	})
	return r
}

func BodyString(resp *resty.Response) string {
	if resp == nil {
		return ""
	}
	return string(resp.Body())
}
