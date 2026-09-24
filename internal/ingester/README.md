# ingester

## Purpose
The `internal/ingester` package manages the core polling, pagination, and retry loops that fetch contract events from Stellar RPC, decode them, and commit them durably to the store.

## Entry Points
- `Ingester`: Main background worker coordinating fetch cycles.
- `New(...) *Ingester`: Constructor for the ingester daemon.
- `Run(ctx context.Context)`: Starts the continuous polling loop.

## Invariants & Design Decisions
1. **Cursor Resume & Ledger Primitives**: On restart, the ingester checks persistent state for a saved cursor or falls back to `LastIngestedLedger + 1`, ensuring zero data loss and no redundant re-scans.
2. **Circuit Breaking & Backoff**: Automatically trips and backs off when Stellar RPC experiences transient downtime or rate limits.
3. **Dead-Letter Handling**: Unparseable or malformed events are isolated to prevent blocking the entire ingestion stream.

## Architecture Cross-Links
- See [`../../docs/architecture.md`](../../docs/architecture.md) for ingester polling state diagrams.
- See [`../../docs/sink.md`](../../docs/sink.md) for delivery and persistence guarantees.
