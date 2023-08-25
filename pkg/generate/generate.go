package generate

import (
	"fmt"
	"generative-web/pkg/templates"
	"generative-web/pkg/openaihandler"
)

type ContentFactory struct {
	Template templates.Template
	Themes   []string
}

func CreateContentFactory(template templates.Template, themes []string) (ContentFactory, error) {
	// check if template is a valid template
	if template.Name == "" {
		return ContentFactory{}, fmt.Errorf("empty template name")
	}
	if len(template.Files) == 0 {
		return ContentFactory{}, fmt.Errorf("template has no files")
	}
	var cf ContentFactory
	cf.Template = template
	cf.Themes = themes
	return cf, nil
}

func (cf *ContentFactory) Generate() (string, error) {
	var content string
	fmt.Println(fmt.Sprintf("generating instance of template: %v", cf.Template.Name))
	if cf.Themes != nil {
		fmt.Println(fmt.Sprintf("Themes: %v", cf.Themes))
		for _, file := range cf.Template.Files {
			fmt.Println(fmt.Sprintf("File: %v", file.Path))
			for _, contentBlock := range file.ContentBlocks {
				fmt.Println(fmt.Sprintf("Content Block: %v", contentBlock))
				prompt := contentBlock.GeneratePromptFromThemes(cf.Themes)
				fmt.Println(fmt.Sprintf("Prompt: %v", prompt))
				openaihandler.RequestSimple()

