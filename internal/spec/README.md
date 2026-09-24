# spec

## Purpose
The `internal/spec` package handles contract-spec enrichment, parsing Soroban contract definitions, and mapping XDR-encoded specifications into queryable metadata.

## Entry Points
- `Enricher`: Interface for contract spec enrichment.
- `NewEnricher(...) Enricher`: Constructs the spec enricher.

## Invariants & Design Decisions
1. **Lossless Fallback**: Unknown or unsupported spec types fall back to raw representations rather than failing ingestion or query parsing.
2. **Schema Alignment**: Seamlessly integrates with the decoder package to provide human-readable names and types for contract events.

## Architecture Cross-Links
- See [`../../docs/architecture-overview.md`](../../docs/architecture-overview.md) for the spec enrichment workflow.
