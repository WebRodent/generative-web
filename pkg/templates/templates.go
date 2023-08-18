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
		filePath := fmt.Sprintf("templates/%[1]v/%[2]v", t.Name, file.Name())
		f := LoadFile(filePath, ext, contentBlockTarget)
		t.Files = append(t.Files, f)
	}
}

func LoadFile(filePath string, ext string, contentBlockTarget string) File {
	var f File
	f.Path = filePath
	f.Type = ext
	f.ContentBlocks = LoadContentBlocks(filePath, contentBlockTarget)
	return f
}

func LoadContentBlocks(filePath string, contentBlockTarget string) []ContentBlock {
	var contentBlocks []ContentBlock
	// open file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Errorf("error opening file: %v", err)
	}
	var content []byte
	var nBytes int
	nBytes, err = file.Read(content)
	if err != nil {
		fmt.Errorf("error reading file: %v", err)
	}
	// read file
	for i := 0; i < nBytes; i++ {
		fmt.Println(fmt.Sprintf("A byte in the file: %v", content[i]))
	}

	return contentBlocks
}
