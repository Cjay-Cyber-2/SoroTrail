# rpc

## Purpose
The `internal/rpc` package provides a robust JSON-RPC client for communicating with Stellar RPC endpoints (`getEvents`, `getLatestLedger`, etc.). It acts as the upstream bridge for the ingester.

## Entry Points
- `Client`: Interface defining supported RPC interactions.
- `NewClient(url string, opts ...ClientOption) Client`: Instantiates the production JSON-RPC client.

## Invariants & Design Decisions
1. **Targeted Surface Area**: The client is intentionally not a complete Stellar SDK; it implements only what the ingester and API lag monitors require.
2. **Resilience & Budgets**: Features built-in budget limiting, request retries with jittered exponential backoff, and client-side circuit breakers.

## Architecture Cross-Links
- See [`../../docs/architecture-overview.md`](../../docs/architecture-overview.md) for the RPC communication layer overview.
