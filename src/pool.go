package src

import "sync"

type ConnectionPool struct {
	connections map[string]*PersistentConnection
	maxConns    int
	mu          sync.RWMutex
}

func NewConnectionPool(maxConns int) *ConnectionPool {
	return &ConnectionPool{
		connections: make(map[string]*PersistentConnection),
		maxConns:    maxConns,
	}
}

// Add into the conneciton pool
func (cp *ConnectionPool) Add(addr string, conn *PersistentConnection) bool {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	if len(cp.connections) >= cp.maxConns {
		return false
	}

	cp.connections[addr] = conn
	return true
}

func (cp *ConnectionPool) Remove(addr string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	delete(cp.connections, addr)
}

func (cp *ConnectionPool) Get(addr string) *PersistentConnection {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	return cp.connections[addr]
}

func (cp *ConnectionPool) Count() int {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return len(cp.connections)
}
