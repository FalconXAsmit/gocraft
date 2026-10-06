package commands

import (
	"fmt"
	"os/exec"
	"runtime"
)

func checkTool(name string) bool {
	_, err := exec.LookPath(name)

	return err == nil
}

func Info() {
	fmt.Println("GoCraft")
	fmt.Println("=======")
	fmt.Println()

	fmt.Println("Version")
	fmt.Println("  0.1.0")
	fmt.Println()

	fmt.Println("System")
	fmt.Println("  OS:", runtime.GOOS)
	fmt.Println("  Architecture:", runtime.GOARCH)
	fmt.Println()

	fmt.Println("Go")
	fmt.Println("  Version:", runtime.Version())
	fmt.Println("  Root:", runtime.GOROOT())
	fmt.Println()

	fmt.Println("Tools")

	if checkTool("git") {
		fmt.Println("  Git: installed")
	} else {
		fmt.Println("  Git: not installed")
	}

	if checkTool("docker") {
		fmt.Println("  Docker: installed")
	} else {
		fmt.Println("  Docker: not installed")
	}
}
