package openaiclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

// OpenAICreateThread создает новый поток в OpenAI API.
func (client *OpenAIClient) OpenAICreateThread(messages []openai.ThreadMessage) (*openai.Thread, error) {
	if client.openaiClient == nil {
		return nil, fmt.Errorf("OpenAI client is not initialized")
	}

	threadReq := openai.ThreadRequest{
		Messages: messages,
	}

	thread, err := client.openaiClient.CreateThread(context.Background(), threadReq)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create thread: %s", err))
	}

	return &thread, nil
}
