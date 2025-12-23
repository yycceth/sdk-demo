package sdkdemo

import (
	"net/http"
	"time"
)

type Client struct {
	config *Config
	http   *http.Client
}

func NewClient(apiKey, secretKey string, opts ...Option) *Client {
	cfg := &Config{
		APIKey:    apiKey,
		SecretKey: secretKey,
		Endpoint:  "https://api.myservice.com",
		Timeout:   10 * time.Second,
	}

	// 应用用户定义的配置
	for _, opt := range opts {
		opt(cfg)
	}

	return &Client{
		config: cfg,
		http: &http.Client{
			Timeout: cfg.Timeout,
		},
	}
}
