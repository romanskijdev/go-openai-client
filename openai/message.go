package openaiclient

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"zod_backend_dev/core/utils"
)

// OpenAISendMessageToThread отправляет сообщение в указанный поток.
func (client *OpenAIClient) OpenAISendMessageToThread(threadID, message *string) (*openai.Message, error) {
	if threadID == nil {
		return nil, fmt.Errorf("threadID cannot be nil")
	}
	if message == nil {
		return nil, fmt.Errorf("message cannot be nil")
	}

	if client.openaiClient == nil {
		return nil, fmt.Errorf("OpenAI client is not initialized")
	}

	messageReq := openai.MessageRequest{
		Role:    openai.ChatMessageRoleUser, // Или "assistant", если сообщение от ассистента
		Content: *message,
	}

	resp, err := client.openaiClient.CreateMessage(context.Background(), *threadID, messageReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send message to thread: %w", err)
	}

	return &resp, nil
}

// OpenAIGetMessagesFromThread получает список сообщений из RUN
func (client *OpenAIClient) OpenAIGetMessagesFromThread(threadID, runID *string) (*openai.MessagesList, error) {
	if threadID == nil {
		return nil, fmt.Errorf("threadID cannot be nil")
	}

	if client.openaiClient == nil {
		return nil, fmt.Errorf("OpenAI client is not initialized")
	}

	resp, err := client.openaiClient.ListMessage(context.Background(), *threadID, utils.PointerToInt(1), nil, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to send message to thread: %w", err)
	}

	return &resp, nil
}
