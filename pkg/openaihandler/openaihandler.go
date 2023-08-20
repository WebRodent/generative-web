package openaihandler

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"

	"generative-web/internal/config"
)

func RequestSimple(conf config.Config, prompt []string, engine string, max_tokens int32, temperature float32, top_p float32, frequency_penalty float32, presence_penalty float32, stop []string) (string, error) {
	// check if api key is not set
	if conf.OpenAI.ApiKey == "" {
		return "", fmt.Errorf("API key not set")
	}

	// create credential
	keyCredential, err := azopenai.NewKeyCredential(conf.OpenAI.ApiKey)
	if err != nil {
		return "", err
	}

	// establish client for pure openai
	client, err := azopenai.NewClientForOpenAI("https://api.openai.com/v1", keyCredential, nil)
	if err != nil {
		return "", err
	}

	// setup context
	ctx := context.Background()

	// setup options
	options := &azopenai.CompletionsOptions{
		Prompt:           prompt,
		MaxTokens:        &max_tokens,
		Temperature:      &temperature,
		TopP:             &top_p,
		FrequencyPenalty: &frequency_penalty,
		PresencePenalty:  &presence_penalty,
		Stop:             stop,
	}

	// make request
	response, err := client.GetCompletions(ctx, *options, nil)
	if err != nil {
		return "", err
	}

	// print response
	fmt.Println(response)

	// return response
	return "", nil
}
