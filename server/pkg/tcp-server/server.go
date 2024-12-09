package tcpserver

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
)

// Server представляет TCP-сервер.
type Server struct {
	address     string
	handlers    map[string]Handler
	clients     map[string]*Client
	broadcastCh chan []byte
	handlersMu  sync.RWMutex
	clientsMu   sync.RWMutex
}

// NewServer создает новый сервер.
func NewServer(address string) *Server {
	return &Server{
		address:     address,
		handlers:    make(map[string]Handler),
		clients:     make(map[string]*Client),
		broadcastCh: make(chan []byte, 100),
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

	go s.handleBroadcastChannel()
	defer close(s.broadcastCh)

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

		client := NewClient(conn)

		s.clientsMu.Lock()
		s.clients[conn.RemoteAddr().String()] = client
		s.clientsMu.Unlock()

		go s.handleConnection(conn)
		go s.handleIncomingChannel(client)
	}
}

// handleConnection обрабатывает соединение.
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	ctx := NewContext(s)

	defer func() {
		s.clientsMu.Lock()
		if client, exists := s.clients[conn.RemoteAddr().String()]; exists {
			close(client.IncomingChan)
		}
		delete(s.clients, conn.RemoteAddr().String())
		s.clientsMu.Unlock()
		fmt.Printf("Client %s disconnected.\n", conn.RemoteAddr().String())
	}()

	for {
		ctx.SetSender(conn.RemoteAddr().String())

		// Чтение сообщения
		header := make([]byte, 4)
		_, err := io.ReadFull(reader, header)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Client %s disconnected.\n", conn.RemoteAddr().String())
			} else {
				fmt.Printf("Error reading from %s: %v\n", conn.RemoteAddr().String(), err)
			}
			return
		}

		messageLen := binary.BigEndian.Uint32(header)
		if messageLen == 0 {
			fmt.Printf("Invalid message length from %s.\n", conn.RemoteAddr().String())
			return
		}

		message := make([]byte, messageLen)
		_, err = io.ReadFull(reader, message)
		if err != nil {
			fmt.Printf("Error reading message from %s: %v\n", conn.RemoteAddr().String(), err)
			return
		}

		// Обработка сообщения
		messageType, payload, err := parseMessage(message)
		if err != nil {
			fmt.Printf("Error parsing message from %s: %v\n", conn.RemoteAddr().String(), err)
			return
		}

		handler, err := s.getHandler(messageType)
		if err != nil {
			fmt.Printf("Handler error for %s: %v\n", conn.RemoteAddr().String(), err)
			ctx.Write(Response{
				Code: ResponseCodeNotFound,
				Data: map[string]string{
					"error": err.Error(),
				},
			})
			return
		}

		ctx.SetMessage(payload)
		err = handler(ctx)
		if err != nil {
			fmt.Printf("Error handling message from %s: %v\n", conn.RemoteAddr().String(), err)
			ctx.Write(Response{
				Code: ResponseCodeServerError,
				Data: map[string]string{
					"error": err.Error(),
				},
			})
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

func (s *Server) getClient(ip string) (*Client, error) {
	s.clientsMu.RLock()
	defer s.clientsMu.RUnlock()
	client, exists := s.clients[ip]
	if !exists {
		return nil, ErrClientNotFound
	}

	return client, nil
}

func (s *Server) handleIncomingChannel(client *Client) {
	for {
		select {
		case data, ok := <-client.IncomingChan:
			if !ok {
				return
			}

			responseLen := make([]byte, 4)
			binary.BigEndian.PutUint32(responseLen, uint32(len(data)))
			_, err := client.Conn.Write(append(responseLen, data...))
			if err != nil {
				fmt.Printf("ERROR: client.Conn.Write: %v\n", err)
			}
		}

	}
}

func (s *Server) handleBroadcastChannel() {
	for {
		select {
		case data, ok := <-s.broadcastCh:
			if !ok {
				for _, client := range s.clients {
					close(client.IncomingChan)
				}

				return
			}

			for _, client := range s.clients {
				client.IncomingChan <- data
			}
		}
	}
}
