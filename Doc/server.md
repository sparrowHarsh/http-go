Client
  ↓
Network Listener
  ↓
Request Parser
  ↓
Middleware (auth, logging, validation)
  ↓
Router
  ↓
Handler (business logic)
  ↓
Response Builder
  ↓
Client


## listener, err := net.Listen("tcp", s.cfg.Address)
Purpose: It represents an open TCP socket that listens for incoming client connections on a specific address (e.g., :8080). Without it, the server can't accept connections—it's the "entry point" for clients.

Key Role:
Binding to Address: net.Listen("tcp", s.cfg.Address) binds the server to the address and starts listening. This creates the listener.
Accepting Connections: The listener's Accept() method waits for and accepts new client connections, returning a net.Conn for each one.

Persistence: By storing it in the Server struct (s.listener), it remains accessible across methods (e.g., in ListenAndServe for the accept loop). If not stored, it goes out of scope and is lost, as discussed earlier.

Lifecycle: The listener should be created once (e.g., in ListenConnection) and closed when the server shuts down (e.g., via defer listener.Close() or in a shutdown method).