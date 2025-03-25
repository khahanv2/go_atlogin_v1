package config

// Config represents application configuration
type Config struct {
	BaseURL string
	ProxyURL string
}

// DefaultBaseURL is the default base URL for API requests
const DefaultBaseURL = "https://example.com"

// NewConfig creates a new configuration with optional overrides
func NewConfig(baseURL, proxyURL string) *Config {
	cfg := &Config{
		BaseURL: DefaultBaseURL,
		ProxyURL: proxyURL,
	}

	// Override base URL if provided
	if baseURL != "" {
		cfg.BaseURL = baseURL
	}

	return cfg
}