package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	completionsEndpoint = "/completions"
)

type CompletionRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
}

type CompletionResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

type CompletionOptions struct {
	Model       string
	MaxTokens   int
	Temperature float64
}

var DefaultCompletionOptions = CompletionOptions{
	Model:       "text-davinci-003",
	MaxTokens:   100,
	Temperature: 0.6,
}

func (c *Client) GetCompletion(prompt string, options ...CompletionOptions) (string, error) {
	options = append([]CompletionOptions{DefaultCompletionOptions}, options...)
	completion := CompletionRequest{
		Prompt: prompt,
	}
	for _, option := range options {
		completion.Model = or(option.Model, completion.Model)
		completion.MaxTokens = or(option.MaxTokens, completion.MaxTokens)
		completion.Temperature = or(option.Temperature, completion.Temperature)
	}

	data, err := json.Marshal(completion)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest(http.MethodPost, getURL(completionsEndpoint), bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	request.Header.Set("Authorization", "Bearer "+c.token)
	request.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(request)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("openai: %s %s", resp.Status, string(data))
	}

	var response CompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("openai: no choices")
	}

	return response.Choices[0].Text, nil
}
