package openaiclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

// OpenAICreateRun создает новый run для потока
func (client *OpenAIClient) OpenAICreateRun(threadID *string, assistant *string) (*openai.Run, error) {
	if client.openaiClient == nil {
		return nil, fmt.Errorf("OpenAI client is not initialized")
	}
	if threadID == nil {
		return nil, fmt.Errorf("thread id is nil")
	}

	var assist *openai.Assistant
	var err error

	if assistant == nil {
		assist, err = client.OpenAIGetAssistant(nil)
		if err != nil {
			return nil, err
		}
		if assist == nil {
			return nil, errors.New("openAI assistant not found")
		}
	} else {
		assist, err = client.OpenAIGetAssistant(assistant)
		if err != nil {
			return nil, err
		}
		if assist == nil {
			return nil, errors.New("openAI assistant not found")
		}
	}

	requestRun := openai.RunRequest{
		AssistantID:  assist.ID,
		Model:        assist.Model,
		Instructions: *assist.Instructions,
	}
	run, err := client.openaiClient.CreateRun(context.Background(), *threadID, requestRun)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create thread: %s", err))
	}

	return &run, nil
}

// OpenAIGetRun получает объект существующего run
func (client *OpenAIClient) OpenAIGetRun(threadID, runID *string) (*openai.Run, error) {
	if client.openaiClient == nil {
		return nil, fmt.Errorf("OpenAI client is not initialized")
	}
	if threadID == nil || runID == nil {
		return nil, fmt.Errorf("thread or run id is nil")
	}

	run, err := client.openaiClient.RetrieveRun(context.Background(), *threadID, *runID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create thread: %s", err))
	}

	return &run, nil
}
