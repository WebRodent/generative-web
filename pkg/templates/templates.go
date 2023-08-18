package templates

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type ContentBlock struct {
	BlockOptions map[string]string `json:"options"`
}

type ContentBlockInterface interface {
	Parse()
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

func (t *Template) Update() {
	dir, err := os.ReadDir("templates/" + t.Name)
	if err != nil {
		fmt.Errorf("error reading template directory: %v", err)
	}
	for i, file := range dir {
		fmt.Println(fmt.Sprintf("File %[1]v found: %[2]v", i+1, file.Name()))
		// get the file extension, by splitting the file name at the dot
		// and getting the last element of the resulting slice
		splitName := strings.Split(file.Name(), ".")
		if len(splitName) < 2 {
			// Handle this case or skip the file if needed
			continue
		}
		ext := splitName[1]
		contentBlockTarget := "GENWEB"
		filePath := fmt.Sprintf("templates/%[1]v/%[2]v", t.Name, file.Name())
		fmt.Println(fmt.Sprintf("Path: %v", filePath))
		f := LoadFile(filePath, ext, contentBlockTarget)
		t.Files = append(t.Files, f)
	}
}

func LoadFile(filePath string, ext string, contentBlockTarget string) File {
	fmt.Println("Loading file...")
	var f File
	f.Path = filePath
	f.Type = ext
	cblocks := LoadContentBlocks(filePath, contentBlockTarget)
	f.ContentBlocks = cblocks
	fmt.Println(fmt.Sprintf("Loaded content blocks: %v", cblocks))
	fmt.Println(fmt.Sprintf("Loaded file: %v", f))
	return f
}

func LoadContentBlocks(filePath string, contentBlockTarget string) []ContentBlock {
	var contentBlocks []ContentBlock
	// make content block target into byte slice
	// first check if the content block target is 6 bytes long
	var cbt []byte
	cbt = make([]byte, len(contentBlockTarget))
	for i := 0; i < len(contentBlockTarget); i++ {
		cbt[i] = contentBlockTarget[i]
	}
	// open file
	fmt.Println(fmt.Sprintf("Opening file: %v", filePath))
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(fmt.Sprintf("error opening file: %v", err))
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	fmt.Println(fmt.Sprintf("File opened: %v", file.Name()))
	var content []byte
	var nBytes int
	content = make([]byte, 10000)
	nBytes, err = file.Read(content)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	fmt.Println(fmt.Sprintf("Read %v bytes from file", nBytes))
	// read file
	for i := 0; i < nBytes; i++ {
		//fmt.Println(fmt.Sprintf("Byte %v: %v", i, content[i]))
		if content[i] == cbt[0] {
			//fmt.Println("Found first byte of content block target")
			// check if the next 5 bytes match the content block target
			for j := 1; j < len(cbt); j++ {
				if content[i+j] == cbt[j] {
					//fmt.Println(fmt.Sprintf("Found byte %v of content block target", j+1))
					if j == len(cbt)-1 {
						//fmt.Println("Found content block target!")
						// extract content block from { to }
						var contentBlockRaw []byte
						for k := i; k < nBytes; k++ {
							if content[k] == 125 {
								//fmt.Println("Found end of content block")
								contentBlockRaw = content[i+len(cbt) : k+1]
								//fmt.Println(fmt.Sprintf("Content block raw: %v", string(contentBlockRaw)))
								break
							}
							// if end of file is reached before end of content block
							if k == nBytes-1 {
								//fmt.Println("End of file reached before end of content block")
								break
							}
						}
						// parse content block
						var contentBlock ContentBlock
						contentBlock.Parse(contentBlockRaw)
						contentBlocks = append(contentBlocks, contentBlock)

					}
				} else {
					fmt.Println("no content block targets found")
					break
				}
			}
		}
	}

	return contentBlocks
}

func (cblock *ContentBlock) Parse(raw []byte) {
	fmt.Println("Parsing content block...")
	fmt.Println(fmt.Sprintf("Raw: %v", string(raw)))
	// Extract content block values for each key by this format: { key:"value"    ,key2 :  "value2",  ... }
	// first search raw for a : character, then keep the characters before it as the key
	// then search for two " characters, then keep the characters between them as the value
	// then search for a , character, then repeat the process

	// first, remove all whitespace from raw, and brackets
	var rawClean []byte
	var keys [][]byte
	var values [][]byte
	for i := 0; i < len(raw); i++ {
		// if character is not whitespace or a bracket, or a " character, keep it
		if raw[i] != 32 && raw[i] != 123 && raw[i] != 34 {
			rawClean = append(rawClean, raw[i])
			// if : is found, extract key
			if raw[i] == 58 {
				//fmt.Println(fmt.Sprintf("Found %v", string(raw[i])))
				var key []byte
				key = rawClean
				//fmt.Println(fmt.Sprintf("Key: %v", string(key)))
				keys = append(keys, key)
				rawClean = make([]byte, 0)
			}
			// if , is found, extract value
			if raw[i] == 44 || raw[i] == 125 {
				//fmt.Println(fmt.Sprintf("Found %v", string(raw[i])))
				var value []byte
				value = rawClean
				//fmt.Println(fmt.Sprintf("Value: %v", string(value)))
				values = append(values, value)
				rawClean = make([]byte, 0)
			}

		}
	}
	//fmt.Println(fmt.Sprintf("Raw clean: %v", string(rawClean)))
	//fmt.Println(fmt.Sprintf("Keys: %v", len(keys)))
	//fmt.Println(fmt.Sprintf("Values: %v", len(values)))
	// for each key, clean the : character from the end
	for i := 0; i < len(keys); i++ {
		keys[i] = keys[i][:len(keys[i])-1]
	}
	//fmt.Println(fmt.Sprintf("Keys: %v", len(keys)))
	// for each value, clean the , character from the end
	for i := 0; i < len(values); i++ {
		values[i] = values[i][:len(values[i])-1]
	}
	//fmt.Println(fmt.Sprintf("Values: %v", len(values)))

	cblock.BlockOptions = make(map[string]string)
	for i := 0; i < len(keys); i++ {
		cblock.BlockOptions[string(keys[i])] = string(values[i])
	}
	//fmt.Println(fmt.Sprintf("Content block: %v", cblock))

}
