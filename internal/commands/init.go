package commands

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func createFile(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func Init(name string) error {
	curVer := runtime.Version()
	curVer = strings.TrimPrefix(curVer, "go")
	fmt.Println("Detected Go version: go", curVer)

	fmt.Println("Initializing project:", name)
	err := os.Mkdir(name, 0755)

	if err != nil {
		return err
	}

	err = createFile(name+"/main.go", `package main

import "fmt"

func main() {
	fmt.Println("Hello from my Go project!")
}
`)

	if err != nil {
		return err
	}

	err = createFile(name+"/go.mod", `module `+name+`

go `+curVer+`
`)

	if err != nil {
		return err
	}

	readmeContent := `# ` + name + `
A go project made by GoCraft!!
	`

	err = createFile(name+"/README.md", readmeContent)
	if err != nil {
		return err
	}

	gitignoreContent := `# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary
*.test

# Coverage
*.out

# Build directory
bin/
`
	err = createFile(name+"/.gitignore", gitignoreContent)
	if err != nil {
		return err
	}
	return nil
}
