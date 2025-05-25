# TECHNICAL CONTEXT

## Technology Stack

### Core Language
- **Go +1.24**: Modern Go language with latest features and performance improvements
- **Minimum Version**: Go 1.24 or higher required

### Dependencies
- **mcp-go**: Core MCP protocol implementation framework
  - Repository: https://github.com/mark3labs/mcp-go
  - Purpose: Provides MCP server protocol handling

### Runtime Requirements
- Go runtime environment
- Support for both SSE and stdio operation modes
- Timezone database access for timezone conversions
- Network capabilities for SSE server mode

## Technical Architecture

### Operation Modes
1. **SSE (Server-Sent Events) Mode**
   - HTTP-based server implementation
   - Real-time event streaming
   - Suitable for web-based integrations

2. **stdio Mode**
   - Standard input/output communication
   - Suitable for CLI tools and direct process communication
   - Lower overhead for simple integrations

### Core Functionality
- **Time Services**:
  - Unix timestamp generation
  - Timezone-aware time formatting
  - Custom format string support
  - Current time retrieval

### Parameters
- **timezone**: String parameter for timezone specification
- **time format**: String parameter for output format customization

## Technical Constraints
- Must comply with MCP protocol specifications
- Thread-safe time operations required
- Memory-efficient for continuous operation
- Error handling for invalid timezones and formats

## Performance Requirements
- Low latency for time requests
- Minimal memory footprint
- Efficient timezone data handling
- Reliable operation under concurrent requests

## Security Considerations
- Input validation for timezone and format parameters
- Prevention of format string attacks
- Safe timezone parsing
- Proper error messaging without information leakage 