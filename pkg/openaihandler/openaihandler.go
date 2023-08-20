package openaihandler

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"

	"generative-web/internal/config"
)

func RequestSimple(conf config.Config, prompt []string, engine string, max_tokens int32, temperature float32, top_p float32, frequency_penalty float32, presence_penalty float32, stop []string, n int32) (string, error) {
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
		DeploymentID:     engine,
		MaxTokens:        &max_tokens,
		Temperature:      &temperature,
		TopP:             &top_p,
		FrequencyPenalty: &frequency_penalty,
		PresencePenalty:  &presence_penalty,
		N:                &n,
		Stop:             stop,
	}

	// make request
	response, err := client.GetCompletions(ctx, *options, nil)
	if err != nil {
		return "", err
	}

	for _, choice := range response.Choices {
		if choice.Text != nil {
			fmt.Println(*choice.Text)
		}
		if choice.Index != nil {
			fmt.Println(*choice.Index)
		}
		if choice.LogProbs != nil {
			fmt.Println(*choice.LogProbs)
		}
		if choice.FinishReason != nil {
			fmt.Println(*choice.FinishReason)
		}
	}
	// print response
	fmt.Println(response)

	// return response
	return "", nil
}
