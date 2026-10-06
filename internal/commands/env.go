package commands

import (
	"fmt"
	"os/exec"
	"strings"
)

func Env() {
	cmd := exec.Command("go", "env", "GOROOT", "GOPATH", "GOMOD", "GOOS", "GOARCH")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	values := strings.Fields(string(output))

	if len(values) != 5 {
		fmt.Println("Error: unexpected Go environment output")
		return
	}

	fmt.Println("Go Environment")
	fmt.Println("--------------")
	fmt.Println("GOROOT:", values[0])
	fmt.Println("GOPATH:", values[1])
	fmt.Println("GOMOD:", values[2])
	fmt.Println("GOOS:", values[3])
	fmt.Println("GOARCH:", values[4])
}
