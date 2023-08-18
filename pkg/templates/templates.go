package templates

import (
	"fmt"
	"os"
	"strings"
)

type ContentBlock struct {
	BlockType string `json:"block_type"`
}

type File struct {
	Path          string         `json:"path"`
	Type          string         `json:"type"`
	ContentBlocks []ContentBlock `json:"content_blocks"`
}

type Template struct {
	Name  string
	Files []File `json:"files"`
}

func LoadTemplate(name string) Template {
	var t Template
	_, err := os.ReadDir("templates/" + name)
	if err != nil {
		fmt.Errorf("error reading template directory: %v", err)
	}
	t.Name = name
	t.Update()
	return t
}

func (t Template) Update() {
	dir, err := os.ReadDir("templates/" + t.Name)
	if err != nil {
		fmt.Errorf("error reading template directory: %v", err)
	}
	for _, file := range dir {
		fmt.Println("File found: %v", file.Name())
		// get the file extension, by splitting the file name at the dot
		// and getting the last element of the resulting slice
		ext := strings.Split(file.Name(), ".")[1]
		contentBlockTarget := "GENWEB"
		f := LoadFile(file.Name(), ext, contentBlockTarget)
		t.Files = append(t.Files, f)
	}
}

func LoadFile(name string, ext string, contentBlockTarget string) File {
	var f File
	f.Path = name
	f.Type = ext
	f.ContentBlocks = LoadContentBlocks(name, contentBlockTarget)
	return f
}

func LoadContentBlocks(name string, contentBlockTarget string) []ContentBlock {
	var contentBlocks []ContentBlock
	// open file
	_, err := os.Open("templates/" + name)
	if err != nil {
		fmt.Errorf("error opening file: %v", err)
	}
	// read file

	return contentBlocks
}
