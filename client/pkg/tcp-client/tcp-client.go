package tcpclient

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

// Client представляет TCP-клиент.
type Client struct {
	address string
	conn    net.Conn
	mu      sync.Mutex
	pending map[string]chan *Message
}

// NewClient создает нового клиента.
func NewClient(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("net.Dial: %w", err)
	}

	client := &Client{
		address: address,
		conn:    conn,
		pending: make(map[string]chan *Message),
	}

	// Запускаем обработчик входящих сообщений
	go client.handleMessages()

	return client, nil
}

// SendMessage отправляет сообщение на сервер и возвращает ответ.
func (c *Client) SendMessage(msg Message) (*Message, error) {
	// Сериализуем сообщение.
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	// Добавляем заголовок с длиной сообщения.
	messageLen := uint32(len(data))
	header := make([]byte, 4)
	binary.BigEndian.PutUint32(header, messageLen)

	// Регистрируем ожидание ответа.
	c.mu.Lock()
	respChan := make(chan *Message, 1)
	c.pending[msg.ID] = respChan
	c.mu.Unlock()

	// Отправляем сообщение.
	_, err = c.conn.Write(append(header, data...))
	if err != nil {
		c.mu.Lock()
		delete(c.pending, msg.ID)
		c.mu.Unlock()
		return nil, fmt.Errorf("c.conn.Write: %w", err)
	}

	// Ждем ответа.
	select {
	case resp := <-respChan:
		return resp, nil
	}
}

// handleMessages обрабатывает входящие сообщения от сервера.
func (c *Client) handleMessages() {
	for {
		// Читаем входящее сообщение.
		rawMessage, err := c.readMessage()
		if err != nil {
			fmt.Println("c.readMessage:", err)
			return
		}

		var msg Message
		err = json.Unmarshal(rawMessage, &msg)
		if err != nil {
			fmt.Println("json.Unmarshal:", err)
			continue
		}

		// Сопоставляем сообщение с ожидающим запросом.
		c.mu.Lock()
		if ch, ok := c.pending[msg.ID]; ok {
			ch <- &msg
			close(ch)
			delete(c.pending, msg.ID)
		}
		c.mu.Unlock()
	}
}

// readMessage читает сообщение от сервера.
func (c *Client) readMessage() ([]byte, error) {
	// Читаем длину сообщения.
	header := make([]byte, 4)
	_, err := c.conn.Read(header)
	if err != nil {
		return nil, fmt.Errorf("c.conn.Read: %w", err)
	}

	messageLen := binary.BigEndian.Uint32(header)

	// Читаем само сообщение.
	message := make([]byte, messageLen)
	_, err = c.conn.Read(message)
	if err != nil {
		return nil, fmt.Errorf("c.conn.Read: %w", err)
	}

	return message, nil
}

// Close закрывает соединение.
func (c *Client) Close() error {
	return c.conn.Close()
}
