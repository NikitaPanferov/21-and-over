package tcpserver

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"

	"github.com/NikitaPanferov/21-and-over/server/pkg/tcp-server/types"
)

// Handler представляет функцию-обработчик для сообщений.
type Handler func(ctx *Context) error

// Server представляет TCP-сервер.
type Server struct {
	address    string
	handlers   map[string]Handler
	conns      map[string]net.Conn
	handlersMu sync.RWMutex
	connsMu    sync.RWMutex
}

// NewServer создает новый сервер.
func NewServer(address string) *Server {
	return &Server{
		address:  address,
		handlers: make(map[string]Handler),
		conns:    make(map[string]net.Conn),
	}
}

// RegisterHandler регистрирует хендлер для определенного типа сообщений.
func (s *Server) RegisterHandler(messageType string, handler Handler) {
	s.handlersMu.Lock()
	defer s.handlersMu.Unlock()
	s.handlers[messageType] = handler
}

// Start запускает сервер.
func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	defer listener.Close()

	fmt.Printf("Server listening on %s\n", s.address)
	s.acceptConnections(listener)
	return nil
}

func (s *Server) acceptConnections(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		s.connsMu.Lock()
		s.conns[conn.RemoteAddr().String()] = conn
		s.connsMu.Unlock()

		go s.handleConnection(conn)
	}
}

// handleConnection обрабатывает соединение.
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		ctx := NewContext(context.Background(), s)
		ctx.SetSender(conn.RemoteAddr().String())
		// Читаем длину сообщения.
		header := make([]byte, 4)
		_, err := io.ReadFull(reader, header)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Client disconnected")
				s.connsMu.Lock()
				s.conns[conn.RemoteAddr().String()] = nil
				s.connsMu.Unlock()
			} else {
				fmt.Printf("Error reading header: %v\n", err)
			}
			return
		}

		messageLen := binary.BigEndian.Uint32(header)
		if messageLen == 0 {
			fmt.Println("Error: invalid message length")
			return
		}

		// Читаем само сообщение.
		message := make([]byte, messageLen)
		_, err = io.ReadFull(reader, message)
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			return
		}

		// Обрабатываем сообщение.
		messageType, payload, err := parseMessage(message)
		if err != nil {
			fmt.Printf("Error parsing message: %v\n", err)
			return
		}

		// Вызываем соответствующий хендлер.
		handler, err := s.getHandler(messageType)
		if err != nil {
			fmt.Printf("Error getting handler: %v\n", err)
			return
		}

		ctx.SetMessage(payload)

		err = handler(ctx)
		if err != nil {
			fmt.Printf("Error handling message: %v\n", err)
		}
	}
}

// parseMessage парсит сообщение и возвращает тип сообщения и полезные данные.
func parseMessage(message []byte) (string, []byte, error) {
	parts := strings.SplitN(string(message), ":", 2)
	if len(parts) < 2 {
		return "", nil, errors.New("invalid message format")
	}
	return parts[0], []byte(parts[1]), nil
}

// getHandler возвращает хендлер для указанного типа сообщения.
func (s *Server) getHandler(messageType string) (Handler, error) {
	s.handlersMu.RLock()
	defer s.handlersMu.RUnlock()
	handler, exists := s.handlers[messageType]
	if !exists {
		return nil, fmt.Errorf("no handler registered for message type: %s", messageType)
	}
	return handler, nil
}

func (s *Server) getConn(IPAddres string) (net.Conn, error) {
	s.connsMu.RLock()
	defer s.connsMu.RUnlock()
	conn, exists := s.conns[IPAddres]
	if !exists {
		return nil, fmt.Errorf("no connection found for IP address: %s", IPAddres)
	}

	if conn == nil {
		return nil, types.ErrConnectionRefused
	}

	return conn, nil
}

func (s *Server) SendToIP(IPAddress string, message []byte) error {
	conn, err := s.getConn(IPAddress)
	if err != nil {
		return fmt.Errorf("s.getConn: %w", err)
	}
	responseLen := make([]byte, 4)
	binary.BigEndian.PutUint32(responseLen, uint32(len(message)))
	_, err = conn.Write(append(responseLen, message...))
	if err != nil {
		return fmt.Errorf("conn.Write: %w", err)
	}
	return nil
}
