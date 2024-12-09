package tcpclient

import (
	"encoding/binary"
	"fmt"
	"net"
)

// Client представляет TCP-клиент.
type Client struct {
	address string
	conn    net.Conn
}

// NewClient создает нового клиента.
func NewClient(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}
	return &Client{
		address: address,
		conn:    conn,
	}, nil
}

// SendMessage отправляет сообщение на сервер и возвращает ответ.
func (c *Client) SendMessage(messageType string, payload []byte) ([]byte, error) {
	// Формируем сообщение.
	message := []byte(fmt.Sprintf("%s:%s", messageType, payload))
	messageLen := uint32(len(message))

	// Создаем заголовок с длиной сообщения.
	header := make([]byte, 4)
	binary.BigEndian.PutUint32(header, messageLen)

	// Отправляем заголовок и сообщение.
	_, err := c.conn.Write(append(header, message...))
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	// Читаем ответ.
	return c.readResponse()
}

// readResponse читает ответ от сервера.
func (c *Client) readResponse() ([]byte, error) {
	// Читаем длину ответа.
	header := make([]byte, 4)
	_, err := c.conn.Read(header)
	if err != nil {
		return nil, fmt.Errorf("failed to read response header: %w", err)
	}

	responseLen := binary.BigEndian.Uint32(header)

	// Читаем само сообщение.
	response := make([]byte, responseLen)
	_, err = c.conn.Read(response)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return response, nil
}

// Close закрывает соединение.
func (c *Client) Close() error {
	return c.conn.Close()
}
