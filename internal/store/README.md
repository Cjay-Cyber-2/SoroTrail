# store

## Purpose
The `internal/store` package provides durable persistence for contract events, ingestion state, watched contracts, replay tracking, and audit tables across PostgreSQL and SQLite. It implements the primary repository pattern for SoroTrail, ensuring clean storage isolation with zero ORM usage (plain SQL via `pgx` and `database/sql`).

## Entry Points
- `Store`: The primary interface defining all storage operations.
- `NewPostgres(pool *pgxpool.Pool) *Postgres`: Constructs a PostgreSQL store instance.
- `NewSQLite(db *sql.DB) *SQLite`: Constructs a SQLite store instance.
- `Migrate(ctx context.Context, dbConn any) error`: Executes pending database migrations.

## Invariants & Design Decisions
1. **Idempotent Upserts**: Inserting the same TOID event twice is a strict no-op (`DO NOTHING` on primary key conflict) that does not corrupt topics, values, or raw XDR fields.
2. **Multi-Network Isolation**: Every event and ingestion state is partitioned by network (e.g. `default`, `testnet`, `mainnet`), preventing cross-network data pollution.
3. **Fail-Closed Authorization Scope**: Queries require an explicit `Scope` object. A zero scope defaults to returning nothing rather than bypassing checks.
4. **Ascending Cursor Pagination**: Event pagination requires deterministic sorting (`id`, `ledger`, `created_at`) returning rows in strict ascending or descending order without gaps or skipped rows during concurrent insertions.

## Architecture Cross-Links
- See [`../../docs/architecture.md`](../../docs/architecture.md) for data flow diagrams and the complete persistence layout.
- See [`../../docs/archival.md`](../../docs/archival.md) for archival and pruning strategies.
