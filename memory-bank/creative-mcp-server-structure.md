📌 CREATIVE PHASE START: MCP Server Structure
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

1️⃣ PROBLEM
   Description: Design MCP server structure to support dual operation modes (SSE and stdio)
   Requirements: 
   - Support both SSE (HTTP-based) and stdio (direct process) communication
   - Clean separation between communication layer and business logic
   - Easy mode switching via configuration
   - Maintain MCP protocol compliance in both modes
   Constraints: 
   - Must use mcp-go framework (v0.29.0)
   - Single binary deployment
   - Minimal configuration complexity

2️⃣ OPTIONS
   Option A: Unified Server with Mode Adapter - Single server with pluggable transport layer
   Option B: Separate Server Implementations - Distinct SSE and stdio server implementations  
   Option C: Factory Pattern with Shared Core - Factory creates mode-specific servers sharing core logic

3️⃣ ANALYSIS
   | Criterion | Unified Server | Separate Servers | Factory Pattern |
   |-----|-----|-----|-----|
   | Code Reuse | ⭐⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐ |
   | Maintainability | ⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐⭐ |
   | Performance | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
   | Complexity | ⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐ |
   | Testability | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
   
   Key Insights:
   - Unified server maximizes code reuse but adds abstraction complexity
   - Separate servers offer best performance but duplicate business logic
   - Factory pattern provides good balance with clear separation of concerns

4️⃣ DECISION
   Selected: Option A: Unified Server with Mode Adapter
   Rationale: Best balance of code reuse and maintainability while keeping complexity manageable. The mcp-go framework likely provides transport abstractions that make this approach natural.
   
5️⃣ IMPLEMENTATION NOTES
   - Create Server interface with Start(), Stop(), RegisterTool() methods
   - Implement transport adapters: SSETransport, StdioTransport
   - Use dependency injection to provide transport to unified server
   - Configuration determines which transport adapter to instantiate
   - Business logic (tool handlers) remains transport-agnostic

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
�� CREATIVE PHASE END 