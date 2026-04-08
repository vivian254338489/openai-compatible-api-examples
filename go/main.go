package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Config holds API configuration
type Config struct {
	APIKey   string
	BaseURL  string
	Model    string
}

// Message represents a chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Request body for chat completions
type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// Response from chat completions
type ChatResponse struct {
	ID      string `json:"id"`
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
	Usage struct {
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`
}

func main() {
	config := loadConfig()

	if config.APIKey == "" {
		fmt.Println("Error: API_KEY environment variable is not set")
		fmt.Println("Set it with: export API_KEY=your_api_key_here")
		fmt.Println("Or create a .env file based on .env.example")
		os.Exit(1)
	}

	fmt.Printf("API: %s\n", config.BaseURL)
	fmt.Printf("Model: %s\n", config.Model)
	fmt.Println(strings.Repeat("-", 40))

	response, err := sendRequest(config)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Success!")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("Response: %s\n", response.Choices[0].Message.Content)
	fmt.Printf("Tokens: %d\n", response.Usage.TotalTokens)
}

func loadConfig() Config {
	return Config{
		APIKey:  os.Getenv("API_KEY"),
		BaseURL: getEnvOrDefault("BASE_URL", "https://tken.shop/v1"),
		Model:   getEnvOrDefault("MODEL", "gpt-4o-mini"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func sendRequest(config Config) (*ChatResponse, error) {
	requestBody := ChatRequest{
		Model: config.Model,
		Messages: []Message{
			{Role: "user", Content: "Say hello in one sentence."},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	url := config.BaseURL + "/chat/completions"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}
