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
	prompt = []string{
		"Make a html file",
		"html file should start with <html> tag",
		"and end with </html>ENDFILE tag+ENDFILE",
		"like this:",
	}
	var engine string = "text-davinci-003"
	var max_tokens int32 = 1024
	var temperature float32 = 0.5
	var top_p float32 = 1
	var frequency_penalty float32 = 0
	var presence_penalty float32 = 0
	var stop []string = []string{"ENDFILE"}
	response, err := RequestSimple(conf, prompt, engine, max_tokens, temperature, top_p, frequency_penalty, presence_penalty, stop, 1)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	fmt.Println(response)
	// print response body

}
