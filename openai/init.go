package openaiclient

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
)

type OpenAIClient struct {
	openaiClient *openai.Client
	config       *OpenAIConfig
}

// NewOpenAIClient создает новый экземпляр Client с заданным API ключом.
func NewOpenAIClient(config *OpenAIConfig) (*OpenAIClient, error) {
	if config == nil {
		return nil, fmt.Errorf("failed to init OpenAIClient: Api key is nil")
	}
	if config.APIKey == "" {
		return nil, fmt.Errorf("failed to init OpenAIClient: Api key is nil")
	}

	return &OpenAIClient{
		openaiClient: openai.NewClient(config.APIKey),
		config:       config,
	}, nil
}
