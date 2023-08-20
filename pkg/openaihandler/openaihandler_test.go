package openaihandler

import (
	"fmt"
	"generative-web/internal/config"
	"testing"
)

// test simple request
func TestRequestSimple(t *testing.T) {
	conf, err := config.Load("../../config.yml")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	var prompt []string
	prompt = []string{"This is a test"}
	var engine string = "davinci"
	var max_tokens int32 = 5
	var temperature float32 = 0.5
	var top_p float32 = 1
	var frequency_penalty float32 = 0
	var presence_penalty float32 = 0
	var stop []string
	response, err := RequestSimple(conf, prompt, engine, max_tokens, temperature, top_p, frequency_penalty, presence_penalty, stop)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	fmt.Println(response)

}
