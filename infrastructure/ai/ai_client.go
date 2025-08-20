package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitcommit/domain/entity"
	"io"
	"net/http"
	"gitcommit/dto"
)

type GrokClient struct {
	APIKey     string
	BaseURL    string
	Model      string
	httpClient *http.Client
}

func NewGrokClient(apiKey string) *GrokClient {
	return &GrokClient{
		APIKey:     apiKey,
		BaseURL:    "https://api.groq.com/openai/v1",
		Model:      "llama3-70b-8192",
		httpClient: &http.Client{},
	}
}

func (c *GrokClient) GenerateChat(messages []entity.Message, maxTokens int) (entity.ChatResponse, error) {
	reqBody := dto.GrokRequestDto{
		Model:       c.Model,
		Messages:    dto.FromDomainMessages(messages),
		MaxTokens:   maxTokens,
		Temperature: 0.7,
		TopP:        1.0,
	}

	// Marshal to JSON
	body, err := json.Marshal(reqBody)
	if err != nil {
		return entity.ChatResponse{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Build request
	req, err := http.NewRequest("POST", c.BaseURL+"/chat/completions", bytes.NewBuffer(body))
	if err != nil {
		return entity.ChatResponse{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return entity.ChatResponse{}, fmt.Errorf("failed to call Grok API: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return entity.ChatResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	// Handle non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return entity.ChatResponse{}, fmt.Errorf("Grok API error: %s", string(raw))
	}

	var res dto.GrokResponseDto
	if err := json.Unmarshal(raw, &res); err != nil {
		return entity.ChatResponse{}, fmt.Errorf("failed to decode Grok response: %w", err)
	}

	return dto.ToDomainResponse(res), nil
}
