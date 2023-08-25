package openaihandler

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"

	"generative-web/internal/config"
)

type Options struct {
	engine            string   `default:"ada"`
	max_tokens        int32    `default:"1"`
	temperature       float32  `default:"0.5"`
	top_p             float32  `default:"1"`
	frequency_penalty float32  `default:"0"`
	presence_penalty  float32  `default:"0"`
	stop              []string `default:"[\"ENDBLOCK\"]"`
	n                 int32    `default:"1"`
}

func RequestSimple(conf config.Config, prompt []string, options Options) (string, error) {
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
	completionOptions := &azopenai.CompletionsOptions{
		Prompt:           prompt,
		DeploymentID:     options.engine,
		MaxTokens:        &options.max_tokens,
		Temperature:      &options.temperature,
		TopP:             &options.top_p,
		FrequencyPenalty: &options.frequency_penalty,
		PresencePenalty:  &options.presence_penalty,
		Stop:             options.stop,
		N:                &options.n,
	}

	// make request
	response, err := client.GetCompletions(ctx, *completionOptions, nil)
	if err != nil {
		return "", err
	}

	for _, choice := range response.Choices {
		if choice.Text != nil {
			fmt.Println("Content of the response:")
			fmt.Println(*choice.Text)
		}
		if choice.Index != nil {
			fmt.Println("Index of the response:")
			fmt.Println(*choice.Index)
		}
		if choice.LogProbs != nil {
			fmt.Println("LogProbs of the response:")
			fmt.Println(*choice.LogProbs)
		}
		if choice.FinishReason != nil {
			fmt.Println("FinishReason of the response:")
			fmt.Println(*choice.FinishReason)
		}
	}
	// print response
	fmt.Println(response)

	// return response
	return "", nil
}
