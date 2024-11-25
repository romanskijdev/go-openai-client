package openaiclient

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

// OpenAICreateAssistant создает нового ассистента с заданным именем,
// используя ключ доступа к ассистенту (Assistant Key).
func (client *OpenAIClient) OpenAICreateAssistant(settings *openai.AssistantRequest) (*openai.Assistant, error) {
	if settings == nil {
		return nil, fmt.Errorf("assistant settings is nil")
	}

	// Отправляем запрос
	assistant, err := client.openaiClient.CreateAssistant(context.Background(), *settings)
	if err != nil {
		return nil, fmt.Errorf("failed to create new assistant: %w", err)
	}
	return &assistant, nil
}

func (client *OpenAIClient) OpenAIDeleteAssistant(assistantID *string) (*openai.AssistantDeleteResponse, error) {
	if assistantID == nil {
		return nil, fmt.Errorf("assistant settings is nil")
	}

	// Отправляем запрос
	assistant, err := client.openaiClient.DeleteAssistant(context.Background(), *assistantID)
	if err != nil {
		return nil, fmt.Errorf("failed to create new assistant: %w", err)
	}

	return &assistant, nil
}

// OpenAIUpdateAssistant Позволяет изменять имя, описание и инструкции для ассистента.
func (client *OpenAIClient) OpenAIUpdateAssistant(settings *openai.AssistantRequest, assistantID *string) (*openai.Assistant, error) {
	if settings == nil {
		return nil, fmt.Errorf("assistant settings is nil")
	}

	if assistantID == nil {
		assistantID = &client.config.AssistantID
	}

	updatedAssistant, err := client.openaiClient.ModifyAssistant(context.Background(), *assistantID, *settings)
	if err != nil {
		fmt.Printf("Failed to update assistant: %v\n", err)
		return nil, err
	}

	return &updatedAssistant, nil
}

// OpenAIGetAssistantList Позволяет получить список ассистентов
func (client *OpenAIClient) OpenAIGetAssistantList(limit *int, order *string, after *string, before *string) (*openai.AssistantsList, error) {
	assistList, err := client.openaiClient.ListAssistants(context.Background(), limit, order, after, before)
	if err != nil {
		fmt.Printf("failed to get assistant list: %s", err)
		return nil, err
	}

	return &assistList, nil
}

// OpenAIGetAssistant Позволяет получить список ассистентов
func (client *OpenAIClient) OpenAIGetAssistant(assistantID *string) (*openai.Assistant, error) {
	if assistantID == nil {
		assistantID = &client.config.AssistantID
	}

	assistList, err := client.openaiClient.RetrieveAssistant(context.Background(), *assistantID)
	if err != nil {
		fmt.Printf("failed to get assistant: %s", err)
		return nil, err
	}

	return &assistList, nil
}
