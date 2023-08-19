package templates

import (
	"fmt"
	"testing"
)

func TestLoadTemplateSingle(t *testing.T) {
	var template = LoadTemplate("single")
	fmt.Println(template)
}

func TestUpdateSingle(t *testing.T) {
	var template = LoadTemplate("single")
	template.Update()
	fmt.Println(template)
}

func TestLoadFileSingle(t *testing.T) {
	var file = LoadFile("../../templates/single/index.html", "html", "GENWEB")
	fmt.Println(file)
}

func TestParseContentBlocks(t *testing.T) {
	var cblock ContentBlock
	var rawBlocks [][]byte
	rawBlocks = [][]byte{
		[]byte("{  name:test, something:else   , foo:  bar}"),
		[]byte("{  type:  text, content:  \"Hello World!\"   }"),
		[]byte("{  type:html, content:  \"<h1>Hello World!</h1>\"   }"),
	}
	for _, rawBlock := range rawBlocks {
		cblock.Parse(rawBlock)
		fmt.Println(cblock)
	}
}
