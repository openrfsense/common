# OpenRFSense Common
This Go module contains common types and packages which are to be shared between the node and backend code.

- `config`: manages configuration loading and value retrieval (dynamically) with generics support.
- `keystore`: provides an in-memory cache for Emitter channel keys with a set TTL.
- `logging`: provides a single-output, leveled logger by wrapping `log.Logger` from the standard library. Uses a single allocation per log call.
- `types`: Go object representations for HTTP requests/responses between nodes and backend.