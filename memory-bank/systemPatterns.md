# SYSTEM PATTERNS

## Architectural Patterns

### MCP Server Pattern
- **Protocol Compliance**: Strict adherence to MCP protocol specifications
- **Dual Mode Support**: Architecture supporting both SSE and stdio communication modes
- **Handler Registration**: Centralized tool/resource handler registration system

### Time Service Pattern
- **Immutable Time Operations**: Time calculations should be stateless and immutable
- **Timezone-Safe Operations**: All time operations must consider timezone context
- **Format Validation**: Input format strings validated before processing

## Code Organization Patterns

### Package Structure
```
go-time-mcp/
├── main.go                 # Application entry point
├── internal/
│   ├── server/            # MCP server implementation
│   ├── time/              # Time service logic
│   └── handlers/          # Request handlers
├── go.mod                 # Go module definition
└── go.sum                 # Dependency checksums
```

### Handler Pattern
- **Tool Handlers**: Implement MCP tool interface for time operations
- **Resource Handlers**: Manage time-related resources if needed
- **Error Handling**: Consistent error response formatting

### Configuration Pattern
- **Environment Variables**: Use environment variables for server configuration
- **Default Values**: Sensible defaults for all configuration options
- **Validation**: Configuration validation at startup

## Error Handling Patterns

### Error Types
- **Invalid Timezone**: Handle unknown or invalid timezone specifications
- **Invalid Format**: Handle malformed time format strings
- **System Errors**: Handle system-level time access errors

### Error Response Format
- Consistent JSON error responses following MCP protocol
- Meaningful error messages for debugging
- Error codes for programmatic handling

## Testing Patterns

### Unit Testing
- **Time Mocking**: Use time interfaces for testable time operations
- **Timezone Testing**: Test with various timezone scenarios
- **Format Testing**: Comprehensive format string validation tests

### Integration Testing
- **MCP Protocol Testing**: End-to-end protocol compliance testing
- **Mode Testing**: Test both SSE and stdio operation modes
- **Concurrent Testing**: Test concurrent request handling

## Security Patterns

### Input Validation
- **Whitelist Approach**: Validate timezone names against known list
- **Format Sanitization**: Sanitize format strings to prevent injection
- **Parameter Limits**: Impose reasonable limits on parameter lengths

### Error Information
- **Safe Error Messages**: Avoid exposing internal system information
- **Logging**: Comprehensive logging for debugging without information leakage
- **Audit Trail**: Track usage patterns for security analysis

## Performance Patterns

### Caching Strategy
- **Timezone Cache**: Cache loaded timezone information
- **Format Cache**: Cache compiled format patterns where applicable
- **Response Caching**: Consider caching for identical requests (time-dependent)

### Resource Management
- **Connection Pooling**: Efficient connection handling for SSE mode
- **Memory Management**: Proper cleanup of time-related objects
- **Goroutine Management**: Controlled concurrency for request handling 