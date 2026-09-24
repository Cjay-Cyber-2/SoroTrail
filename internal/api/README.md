# api

## Purpose
The `internal/api` package exposes the queryable HTTP API and GraphQL server for SoroTrail. It handles routing, authentication, rate limiting, request compression, SSE / WebSocket streaming subscriptions, and filtered event queries long after Stellar RPC forgets them.

## Entry Points
- `Server`: HTTP server wrapper containing the Chi router and middleware stack.
- `NewServer(...) *Server`: Factory function to instantiate the API server.
- `graphql.NewExecutableSchema`: GraphQL query/subscription execution engine.

## Invariants & Design Decisions
1. **Read-Only Separation**: HTTP endpoints are strictly read-only unless authenticated with API keys or multi-tenancy headers.
2. **Independent Tenant Scoping**: Subscriptions and queries enforce tenant grants via the `Scope` mechanism, preventing cross-tenant data leaks or privilege escalation through custom filters.
3. **Optimized Caching & Compression**: Built-in response caching and compression middleware minimize database load during heavy analytics queries.

## Architecture Cross-Links
- See [`../../docs/architecture-overview.md`](../../docs/architecture-overview.md) for the API seam placement.
- See [`../../docs/multi-tenancy.md`](../../docs/multi-tenancy.md) for scoping and multi-tenant authorization rules.
