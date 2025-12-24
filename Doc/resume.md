## HTTP Server Framework (Go)

- Built a lightweight HTTP/1.1 server from scratch in Go using raw TCP sockets, implementing request parsing, response serialization, and concurrent connection handling with goroutines
- Designed an extensible router with support for dynamic route parameters (`/users/:id`), middleware chaining (logging, CORS, rate-limiting), and RESTful HTTP method handlers
- Implemented connection keep-alive, graceful shutdown with signal handling, and configurable read/write timeouts for production-grade reliability
- Achieved high throughput by leveraging Go's concurrency model, buffered I/O, and efficient memory management for parsing HTTP headers and request bodies
