package src

import (
	"net"
	"sync"
	"time"
)

type ConnectionState int

const (
	StateIdle ConnectionState = iota
	StateActive
	StateClosed
)

/* Persistant connection struct
- Fields needed

- conn, currstate, reqcount, maxReq

*/

type PersistentConnection struct {
	Conn         net.Conn
	State        ConnectionState
	RequestCount int
	MaxRequests  int
	CreatedAt    time.Time
	LastActiveAt time.Time
	IdleTimeout  time.Duration
	mu           sync.Mutex
}

// Create a New Persistent connection
func NewPersistentConnection(conn net.Conn, maxRequests int, idleTimeout time.Duration) *PersistentConnection {
	return &PersistentConnection{
		Conn:         conn,
		State:        StateIdle,
		MaxRequests:  maxRequests,
		CreatedAt:    time.Now(),
		LastActiveAt: time.Now(),
		IdleTimeout:  idleTimeout,
	}
}

// close underlying connections
func (pc *PersistentConnection) Close() error {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.State = StateClosed

	// closing currrent connection
	return pc.Conn.Close()
}

// Increment request count
func (pc *PersistentConnection) IncrementRequestCount() {
	pc.mu.Lock()
	defer pc.mu.Unlock()
	pc.RequestCount++
	pc.LastActiveAt = time.Now()
}

// CanServeMore return if connection can handle more request
func (pc *PersistentConnection) CanserveMore() bool {
	pc.mu.Lock()
	defer pc.mu.Unlock()

	return pc.RequestCount < pc.MaxRequests
}

//
