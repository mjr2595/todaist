package config

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

type Config struct {
	TodoistAPIToken string
	GeminiAPIKey    string
	Timezone        string
	DateLang        string
}

type Clients struct {
	Gemini *genai.Client
	Config *Config
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	// Try to load .env file, but don't fail if it doesn't exist
	_ = godotenv.Load()

	config := &Config{
		TodoistAPIToken: os.Getenv("TODOIST_API_TOKEN"),
		GeminiAPIKey:    os.Getenv("GOOGLE_API_KEY"),
		Timezone:        "America/Chicago", // Default timezone
		DateLang:        "en",              // Default language
	}

	if config.TodoistAPIToken == "" {
		return nil, fmt.Errorf("TODOIST_API_TOKEN environment variable is required")
	}

	if config.GeminiAPIKey == "" {
		return nil, fmt.Errorf("GOOGLE_API_KEY environment variable is required")
	}

	return config, nil
}

// NewClients initializes API clients with the configuration
func NewClients(ctx context.Context, config *Config) (*Clients, error) {
	// Initialize Gemini client
	geminiClient, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  config.GeminiAPIKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	return &Clients{
		Gemini: geminiClient,
		Config: config,
	}, nil
}

// Close closes all API clients
func (c *Clients) Close() error {
	// Note: genai.Client doesn't have a Close method in the current version
	// This is a placeholder for future cleanup if needed
	return nil
}