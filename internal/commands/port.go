package commands

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	gnet "github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
)

func Port(port string) {

	portNumber, err := strconv.Atoi(port)

	if err != nil || portNumber < 1 || portNumber > 65535 {
		fmt.Printf("Error: invalid port \"%s\"\n", port)
		return
	}

	address := ":" + strconv.Itoa(portNumber)
	listner, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Port %s already in use\n", port)
		return
	}
	listner.Close()
	fmt.Println("Port", port, "is available")

}

func WatchPorts() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		fmt.Print("\033[H\033[2J\033[3J")

		ListPorts()

		fmt.Println()
		fmt.Println("Refreshing every 2s... Press Ctrl+C to exit.")

		<-ticker.C
	}
}

func ListPorts() {
	fmt.Println("GoCraft Port Monitor")
	fmt.Println("--------------------")
	fmt.Println()
	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)

	fmt.Fprintln(writer, "ADDRESS\tPORT\tPID\tPROCESS\tSTATUS")

	connections, err := gnet.Connections("all")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	seenPorts := make(map[uint32]bool)
	for _, conn := range connections {
		if conn.Status != "LISTEN" {
			continue
		}

		if seenPorts[conn.Laddr.Port] {
			continue
		}

		seenPorts[conn.Laddr.Port] = true

		proc, err := process.NewProcess(conn.Pid)

		name := "-"

		if err == nil {
			name, err = proc.Name()

			if err != nil {
				name = "-"
			}
		}

		fmt.Fprintf(
			writer,
			"%s\t%d\t%d\t%s\t%s\n",
			conn.Laddr.IP,
			conn.Laddr.Port,
			conn.Pid,
			name,
			conn.Status,
		)
	}

	writer.Flush()
}
