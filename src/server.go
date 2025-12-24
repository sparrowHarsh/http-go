package src

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"time"
)

/*
ServerConfig -> Basic thigns to start a server, just a desing pattern
*/
type ServerConfig struct {
	Address      string
	ReadTimeout  time.Duration // max time to read request
	WriteTimeout time.Duration // max time to write response
	Logger       *log.Logger
}

type Server struct {
	cfg ServerConfig

	listener       net.Listener
	ShutdownCtx    context.Context
	ShutdownCancel context.CancelFunc
}

// Create New connection
func NewServer(cfg ServerConfig) *Server {
	shutdownCtx, shutdownCancel := context.WithCancel(context.Background())

	cfg.Logger.Println("Create new server")
	return &Server{
		cfg:            cfg,
		ShutdownCtx:    shutdownCtx,
		ShutdownCancel: shutdownCancel,
	}
}

/*
Start Listening
*/

func (s *Server) ListenConnection() error {
	s.cfg.Logger.Println("Listen Connection function has been called")

	listener, err := net.Listen("tcp", s.cfg.Address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", s.cfg.Address, err)
	}

	s.listener = listener
	defer listener.Close()
	s.cfg.Logger.Println("Server started listening successfully at address", s.cfg.Address)

	/* Accept conncetion for multiple clients*/
	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-s.ShutdownCtx.Done():
				s.cfg.Logger.Println("Shutting down server")
				return nil

			default:
				s.cfg.Logger.Println("Error accepting connection:", err)
				continue
			}
		}
		// Now handle each connection in a seperate go routine
		go s.HandleConnection(conn)
	}
}

// Function for handeling conencton
func (s *Server) HandleConnection(conn net.Conn) error {
	s.cfg.Logger.Println("Handle connecton function called")

	// set read/write timeouts
	conn.SetReadDeadline(time.Now().Add(s.cfg.ReadTimeout))
	conn.SetWriteDeadline(time.Now().Add(s.cfg.WriteTimeout))
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Parse the complete HTTP request (headers + body)
	request, err := ParseHttpRequest(reader)
	if err != nil {
		return fmt.Errorf("error parsing request: %w", err)
	}

	s.cfg.Logger.Printf("Method: %s, Path: %s, Version: %s", request.Method, request.Path, request.Version)
	s.cfg.Logger.Printf("Headers: %v", request.Header)
	s.cfg.Logger.Printf("Body: %s", request.Body)

	// TODO: Route the request and send response
	// Send a simple HTTP response
	response := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 2\r\n\r\nOK"
	conn.Write([]byte(response))

	return nil
}
