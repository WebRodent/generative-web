package generate

import (
	"fmt"

	"generative-web/pkg/templates"
)

type Prompt interface {
	Generate() (string, error)
	AddRequirements([]string) error
}

type GenerateContext struct {
	Prompt   Prompt
	Template templates.Template
	blocks   []int
}

type PreBuiltTags struct {
	PreText    string
	TagsList   []string
	ResultText string
}

func (pbt *PreBuiltTags) AddRequirements(reqs []string) error {
	pbt.TagsList = reqs
	return nil
}

func (pbt *PreBuiltTags) Generate() (string, error) {
	if pbt.TagsList == nil {
		return "", fmt.Errorf("No tags specified")
	}
	return pbt.PreText, nil
}
