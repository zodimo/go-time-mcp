# go-time-mcp

A Model Context Protocol (MCP) server that provides time-related operations with timezone support and custom formatting. Available to run as SSE server or in stdio mode.

## Features

- Get current Unix timestamp
- Get current time in any timezone (IANA, abbreviations, offsets)
- Custom time formatting support
- Dual operation modes: SSE (HTTP) and stdio
- Robust timezone validation and error handling
- Go 1.24+ with minimal dependencies

## Installation

### Using go install

```bash
go install github.com/zodimo/go-time-mcp@latest
```

### From source

```bash
git clone https://github.com/zodimo/go-time-mcp.git
cd go-time-mcp
go build -o go-time-mcp .
```

## Usage

### Command Line Options

```bash
# Run in stdio mode (default)
go-time-mcp

# Run in SSE mode on port 8080 (default)
go-time-mcp -mode sse

# Run in SSE mode on custom port
go-time-mcp -mode sse -port 3000

# Set request timeout
go-time-mcp -timeout 60s

# Set log level
go-time-mcp -log-level debug
```

### Environment Variables

You can also configure the server using environment variables:

```bash
export MCP_MODE=sse
export MCP_PORT=8080
export MCP_TIMEOUT=30s
export MCP_LOG_LEVEL=info
go-time-mcp
```

## MCP Configuration

### For Cursor.com

Create or update your MCP configuration file at:
- **Linux/macOS**: `~/.cursor/mcp.json`
- **Windows**: `%APPDATA%\Cursor\mcp.json`

#### Stdio Mode Configuration

```json
{
  "mcpServers": {
    "go-time-mcp": {
      "command": "go-time-mcp",
      "args": ["-mode", "stdio"],
      "env": {
        "MCP_LOG_LEVEL": "info"
      }
    }
  }
}
```

#### SSE Mode Configuration

```json
{
  "mcpServers": {
    "go-time-mcp": {
      "command": "go-time-mcp",
      "args": ["-mode", "sse", "-port", "8080"],
      "env": {
        "MCP_LOG_LEVEL": "info"
      }
    }
  }
}
```

### For Visual Studio Code

Create or update your MCP configuration file at:
- **Linux/macOS**: `~/.config/Code/User/mcp.json`
- **Windows**: `%APPDATA%\Code\User\mcp.json`

#### Stdio Mode Configuration

```json
{
  "mcpServers": {
    "go-time-mcp": {
      "command": "go-time-mcp",
      "args": ["-mode", "stdio", "-log-level", "info"]
    }
  }
}
```

#### SSE Mode Configuration

```json
{
  "mcpServers": {
    "go-time-mcp": {
      "command": "go-time-mcp",
      "args": ["-mode", "sse", "-port", "8080", "-log-level", "info"]
    }
  }
}
```

## Available MCP Tools

### getCurrentTime

Get the current time in a specified timezone with optional formatting.

**Parameters:**
- `timezone` (optional): IANA timezone (e.g., "America/New_York"), abbreviation (e.g., "EST"), or offset (e.g., "+05:00")
- `format` (optional): Time format string (supports common patterns like "YYYY-MM-DD HH:mm:ss")

**Example:**
```json
{
  "timezone": "America/New_York",
  "format": "YYYY-MM-DD HH:mm:ss"
}
```

### getUnixTimestamp

Get the current Unix timestamp (seconds since epoch).

**Parameters:** None

**Returns:** Current Unix timestamp as integer

## Supported Timezones

- **IANA Timezones**: `America/New_York`, `Europe/London`, `Asia/Tokyo`, etc.
- **Abbreviations**: `EST`, `CST`, `MST`, `PST`, `CET`, `JST`, etc.
- **UTC Offsets**: `+05:00`, `-08:00`, `+0530`, etc.
- **Special**: `UTC`, `GMT`

## Format Patterns

The server supports common time format patterns:

- `YYYY` or `yyyy` → 4-digit year (2024)
- `YY` or `yy` → 2-digit year (24)
- `MM` → Month with leading zero (01-12)
- `DD` or `dd` → Day with leading zero (01-31)
- `HH` → Hour in 24-hour format (00-23)
- `hh` → Hour in 12-hour format (01-12)
- `mm` → Minutes (00-59)
- `ss` → Seconds (00-59)
- `SSS` → Milliseconds (000-999)

**Example formats:**
- `YYYY-MM-DD HH:mm:ss` → `2024-01-15 14:30:45`
- `MM/DD/YYYY hh:mm:ss` → `01/15/2024 02:30:45`
- `DD.MM.YY HH:mm` → `15.01.24 14:30`

## Configuration Options

| Flag | Environment Variable | Default | Description |
|------|---------------------|---------|-------------|
| `-mode` | `MCP_MODE` | `stdio` | Server mode: `sse` or `stdio` |
| `-port` | `MCP_PORT` | `8080` | Port for SSE mode |
| `-timeout` | `MCP_TIMEOUT` | `30s` | Request timeout |
| `-log-level` | `MCP_LOG_LEVEL` | `info` | Log level: `debug`, `info`, `warn`, `error` |

## Development

### Requirements

- Go 1.24 or higher
- [mcp-go](https://github.com/mark3labs/mcp-go) v0.29.0

### Building

```bash
go build -o go-time-mcp .
```

### Testing

```bash
go test ./...
```

## License

This project is licensed under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
