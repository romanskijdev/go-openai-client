package openaiclient

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func (client *OpenAIClient) OpenAICreateDailyHoroscopeAssistant() (*openai.Assistant, error) {
	// Имя и описание для JSON Schema
	schemaName := "HoroscopeResponse"
	schemaDescription := "A JSON response format containing daily horoscopes for zodiac signs."

	// Инструкция для ассистента
	instructions := `You are an astrologer assistant. Generate daily horoscopes for all 12 zodiac signs. Always respond using JSON format with an array of horoscopes, each containing the "text" of the horoscope and the "sign" of the zodiac sign.`

	// Создаём объект ChatCompletionResponseFormat
	responseFormat := openai.ChatCompletionResponseFormat{
		Type: openai.ChatCompletionResponseFormatTypeJSONSchema, // Указываем тип
		JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
			Name:        schemaName,
			Description: schemaDescription,
			Schema:      HoroscopeSchema{}, // Используем созданную выше структуру для схемы
			Strict:      true,              // Используем строгий режим
		},
	}

	// Температура и другие параметры
	temperature := float32(0.7)
	topP := float32(1.0)

	// Создание настроек для ассистента
	settings := openai.AssistantRequest{
		Model:          "gpt-4o-mini",
		Name:           &schemaName,
		Description:    &schemaDescription,
		Instructions:   &instructions,
		ResponseFormat: responseFormat, // Используем правильный формат ответа
		Temperature:    &temperature,
		TopP:           &topP,
	}

	// Отправляем запрос на создание ассистента
	assistant, err := client.openaiClient.CreateAssistant(context.Background(), settings)
	if err != nil {
		return nil, fmt.Errorf("failed to create new assistant: %w", err)
	}

	return &assistant, nil
}

func (client *OpenAIClient) OpenAICreateCompatibilityAssistant() (*openai.Assistant, error) {
	// Имя и описание для JSON Schema
	schemaName := "CompatibilityResponse"
	schemaDescription := "A JSON response format containing compatibility analysis details between two entities."

	// Инструкция для ассистента
	instructions := `You are a compatibility assistant. Generate compatibility analysis between two people, including textual descriptions and percentage compatibility scores. Always respond using JSON format.`

	// Создаём объект ChatCompletionResponseFormat
	responseFormat := openai.ChatCompletionResponseFormat{
		Type: openai.ChatCompletionResponseFormatTypeJSONSchema, // Указываем тип
		JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
			Name:        schemaName,
			Description: schemaDescription,
			Schema:      CompatibilitySchema{}, // Используем созданную выше структуру для схемы
			Strict:      true,                  // Используем строгий режим
		},
	}

	// Температура и другие параметры
	temperature := float32(0.7)
	topP := float32(1.0)

	// Создание настроек для ассистента
	settings := openai.AssistantRequest{
		Model:          "gpt-4o-mini",
		Name:           &schemaName,
		Description:    &schemaDescription,
		Instructions:   &instructions,
		ResponseFormat: responseFormat, // Используем правильный формат ответа
		Temperature:    &temperature,
		TopP:           &topP,
	}

	// Отправляем запрос на создание ассистента
	assistant, err := client.openaiClient.CreateAssistant(context.Background(), settings)
	if err != nil {
		return nil, fmt.Errorf("failed to create new assistant: %w", err)
	}

	return &assistant, nil
}

func (client *OpenAIClient) OpenAICreatePersonalHoroscopeAssistant() (*openai.Assistant, error) {
	// Имя и описание для JSON Schema
	schemaName := "PersonalHoroscopeResponse"
	schemaDescription := "A JSON response format containing personal horoscope information."

	// Инструкция для ассистента
	instructions := `You are an astrologer assistant. Generate a personal horoscope with the main horoscope text, horoscopes for the first and second spheres, additional advice, and success percentages. Always respond using JSON format.`

	// Создаём объект ChatCompletionResponseFormat
	responseFormat := openai.ChatCompletionResponseFormat{
		Type: openai.ChatCompletionResponseFormatTypeJSONSchema, // Указываем тип
		JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
			Name:        schemaName,
			Description: schemaDescription,
			Schema:      PersonalHoroscopeSchema{}, // Используем созданную выше структуру для схемы
			Strict:      true,                      // Используем строгий режим
		},
	}

	// Температура и другие параметры
	temperature := float32(0.7)
	topP := float32(1.0)

	// Создание настроек для ассистента
	settings := openai.AssistantRequest{
		Model:          "gpt-4o-mini",
		Name:           &schemaName,
		Description:    &schemaDescription,
		Instructions:   &instructions,
		ResponseFormat: responseFormat, // Используем правильный формат ответа
		Temperature:    &temperature,
		TopP:           &topP,
	}

	// Отправляем запрос на создание ассистента
	assistant, err := client.openaiClient.CreateAssistant(context.Background(), settings)
	if err != nil {
		return nil, fmt.Errorf("failed to create new assistant: %w", err)
	}

	return &assistant, nil
}

func (client *OpenAIClient) OpenAICreateChattingAssistant(name, instructions *string, temperature, topP *float32) (*openai.Assistant, error) {
	// Имя и описание для JSON Schema
	schemaName := "ChattingResponse"
	schemaDescription := "A JSON response format containing message response information."

	// Создаём объект ChatCompletionResponseFormat
	responseFormat := openai.ChatCompletionResponseFormat{
		Type: openai.ChatCompletionResponseFormatTypeJSONSchema, // Указываем тип
		JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
			Name:        schemaName,
			Description: schemaDescription,
			Schema:      ChattingSchema{}, // Используем созданную выше структуру для схемы
			Strict:      true,             // Используем строгий режим
		},
	}

	// Создание настроек для ассистента
	settings := openai.AssistantRequest{
		Model:          "gpt-4o-mini",
		Name:           name,
		Description:    &schemaDescription,
		Instructions:   instructions,
		ResponseFormat: responseFormat, // Используем правильный формат ответа
		Temperature:    temperature,
		TopP:           topP,
	}

	// Отправляем запрос на создание ассистента
	assistant, err := client.openaiClient.CreateAssistant(context.Background(), settings)
	if err != nil {
		return nil, fmt.Errorf("failed to create new assistant: %w", err)
	}

	return &assistant, nil
}
