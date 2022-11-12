# OpenRFSense Common
This Go module contains common types and packages which are to be shared between the node and backend code.

- `id`: provides a random string generator seeded either with the current time or with an arbitrary byte slice. Used to generate various kinds of IDs internally (node hardware ID, campaign ID).
- `logging`: provides a single-output, leveled logger by wrapping `log.Logger` from the standard library. Uses a single allocation per log call.
- `stats`: contains a simple matrics/statistics manager for nodes, with arbitrary information provided by any object implementing the relevant interface.
- `types`: Go object representations for HTTP requests/responses between clients and backend, with validation.