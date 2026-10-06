package commands

import "fmt"

func Help() {
	fmt.Println("GoCraft - Developer toolkit for Go")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  gocraft <command>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  init <name>       Create a new Go project")
	fmt.Println("  info              Show GoCraft and system information")
	fmt.Println("  env               Show Go environment")
	fmt.Println("  port              Show active listening ports")
	fmt.Println("  port <port>       Check if a port is available")
	fmt.Println("  port --watch      Monitor listening ports")
	fmt.Println("  version           Show GoCraft version")
	fmt.Println("  help              Show this help message")
}
