# GoCraft

![Go Version](https://img.shields.io/badge/Go-1.25%2B-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-green)
![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20Linux%20%7C%20macOS-lightgrey)

**GoCraft** is a developer-focused CLI toolkit for Go. Inspect your Go environment, check ports, monitor active network ports, and scaffold new Go projects, all from one small binary.

## Features

- 🚀 Generate new Go projects with `init`
- 🖥️ Inspect system and Go info with `info`
- ⚙️ View important Go environment variables with `env`
- 🔌 Check whether a port is available
- 📡 List active listening ports and their processes
- 👀 Monitor ports in real time with `--watch`
- ℹ️ Display version information
- 📖 Built-in command help

## Installation

### Using Go

```bash
go install github.com/FalconXAsmit/gocraft@latest
```

Verify the install:

```bash
gocraft version
```

> Make sure your Go binary directory (`$GOPATH/bin` or `$HOME/go/bin`) is in your `PATH`.

### From Source

```bash
git clone https://github.com/FalconXAsmit/gocraft.git
cd gocraft
go install .
```

## Quick Start

```bash
gocraft help              # see all commands
gocraft init my-api       # scaffold a new project
gocraft port --watch      # live port monitor
```

## Commands

| Command                | Description                          |
| ---------------------- | ------------------------------------ |
| `gocraft init <name>`  | Create a new Go project              |
| `gocraft info`         | Show GoCraft and system information  |
| `gocraft env`          | Show Go environment                  |
| `gocraft port`         | Show active listening ports          |
| `gocraft port <port>`  | Check if a port is available         |
| `gocraft port --watch` | Monitor listening ports              |
| `gocraft version`      | Show GoCraft version                 |
| `gocraft help`         | Show available commands              |

`gocraft -h` and `gocraft --help` also work as aliases for `help`.

## Usage

### Create a Project

```bash
gocraft init my-api
```

Generates:

```
my-api/
├── main.go
├── go.mod
├── README.md
└── .gitignore
```

The generated `go.mod` uses the Go version detected from the environment running GoCraft.

### System Information

```bash
gocraft info
```

Shows:

- GoCraft version
- Operating system and architecture
- Go version and installation root
- Git availability
- Docker availability

### Go Environment

```bash
gocraft env
```

Shows `GOROOT`, `GOPATH`, `GOMOD`, `GOOS`, and `GOARCH`.

### Check a Port

```bash
gocraft port 8080
```

Checks whether a TCP port is available.

### List Active Ports

```bash
gocraft port
```

Example output:

```
GoCraft Port Monitor
--------------------

ADDRESS       PORT    PID    PROCESS       STATUS
127.0.0.1     8080    1234   my-api.exe    LISTEN
0.0.0.0       5432    2380   postgres.exe  LISTEN
```

### Watch Ports

```bash
gocraft port --watch
```

Refreshes the listening ports every seconds. Press `Ctrl+C` to stop.

## Requirements

- Go 1.25+
- Git
- Windows, Linux, or macOS

GoCraft uses [gopsutil](https://github.com/shirou/gopsutil) for network and process information.

## Project Structure

```
gocraft/
├── main.go
├── go.mod
├── go.sum
├── README.md
└── internal/
    └── commands/
        ├── env.go
        ├── help.go
        ├── info.go
        ├── init.go
        ├── port.go
        └── version.go
```

## Development

```bash
git clone https://github.com/FalconXAsmit/gocraft.git
cd gocraft

go run .              # run GoCraft
go run . info         # run a specific command
go run . port 8080

gofmt -w .            # format
go build .            # build
```

## Version 1

GoCraft v1 is a small, useful developer CLI. It includes:

- Project generation
- Go environment and system information
- Port availability checks
- Active port monitoring and live watching
- CLI help and version commands

**Not in v1** (planned for later): Docker project generation, additional project templates, and more developer workflow tools.

## Roadmap

- 🐳 Optional Docker support
- 📦 More project templates
- 🧰 Additional developer utilities
- ⚡ More project configuration options
- 🔧 Improved project generation
- 📊 More system and dev environment diagnostics

## Contributing

Contributions, ideas, and suggestions are welcome. Found a bug or have a feature idea? Open an issue or submit a pull request.

## License

MIT License. Copyright (c) 2026 Asmit.