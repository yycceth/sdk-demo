package sdkdemo

import "time"

// Config 定义配置
type Config struct {
	APIKey    string
	SecretKey string
	Endpoint  string
	Timeout   time.Duration
	Retries   int
}

// Option 定义配置函数类型
type Option func(*Config)

func WithTimeout(t time.Duration) Option {
	return func(c *Config) { c.Timeout = t }
}

func WithEndpoint(e string) Option {
	return func(c *Config) { c.Endpoint = e }
}
