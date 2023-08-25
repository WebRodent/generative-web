package openaihandler

import (
	"fmt"
	"generative-web/internal/config"
	"testing"
)

// test simple request
// cant test this because of openai api key
func TestRequestSimple(t *testing.T) {
	conf, err := config.Load("../../config.yml")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	var topics []string
	topics = []string{
		"IaC",
		"AI",
		"Golang",
		"Development",
	}
	var prompt []string
	prompt = []string{
		`Make a html file
		html file should start with <html> tag
		and end with </html>ENDFILE tag+ENDFILE
		html file should follow a few content requirements
		html file is a blog post
		the blog discusses the following topics:
		`,
	}
	for _, topic := range topics {
		prompt[0] = prompt[0] + topic + ", "
	}
	prompt = append(prompt, "The html file:")
	var engine string = "ada"
	var max_tokens int32 = 1
	var temperature float32 = 0.5
	var top_p float32 = 1
	var frequency_penalty float32 = 0
	var presence_penalty float32 = 0
	var stop []string = []string{"ENDFILE"}
	var n int32 = 1
	response, err := RequestSimple(
		conf,
		prompt,
		engine,
		max_tokens,
		temperature,
		top_p,
		frequency_penalty,
		presence_penalty,
		stop,
		n,
	)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	fmt.Println(response)
	// print response body

}
