# PRODUCT CONTEXT

## Product Purpose
This MCP server provides time-related services to client applications that need accurate timestamp and formatted time data.

## Target Users
- Developers integrating time services into their applications
- Applications requiring timezone-aware time data
- Systems needing both unix timestamps and human-readable time formats

## Core Value Proposition
- **Precision**: Accurate time data with timezone support
- **Flexibility**: Multiple output formats for different use cases
- **Integration**: Standard MCP protocol for easy integration
- **Dual Modes**: Both server and stdio operation modes for different deployment scenarios

## Use Cases
1. **Unix Timestamp Retrieval**: Applications needing standardized time markers
2. **Timezone Conversion**: Converting time to specific regional formats
3. **Custom Formatting**: Displaying time in user-preferred formats
4. **Server Integration**: Backend services requiring time data
5. **CLI Tools**: Command-line applications needing time services

## Product Constraints
- Must follow MCP protocol standards
- Requires Go +1.24 runtime environment
- Dependent on mcp-go framework
- Must support both SSE and stdio modes

## Success Metrics
- Accurate time delivery across timezones
- Reliable operation in both server modes
- Easy integration with MCP-compatible clients
- Minimal latency for time requests 