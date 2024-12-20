package tcpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// Context представляет пользовательский контекст, совместимый с context.Context.
type Context struct {
	ctx     context.Context
	server  *Server
	message []byte
	sender  string
}

// NewContext создает новый пользовательский контекст.
func NewContext(server *Server) *Context {
	return &Context{
		ctx:    context.Background(),
		server: server,
	}
}

// SetMessage устанавливает сообщение в контекст.
func (c *Context) SetMessage(message []byte) {
	c.message = message
}

// GetMessage возвращает сообщение из контекста.
func (c *Context) GetMessage() []byte {
	return c.message
}

func (c *Context) SetSender(sender string) {
	c.sender = sender
}

func (c *Context) GetSender() string {
	return c.sender
}

// Write отправляет ответ клиенту.
func (c *Context) Write(data any) error {
	dataToSend, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	return c.SendToIP(c.sender, dataToSend)
}

func (c *Context) WriteAll(data any) error {
	dataToSend, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	return c.SendToAll(dataToSend)
}

// Value возвращает значение из контекста, реализуя интерфейс context.Context.
func (c *Context) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}

// Deadline, Done, Err - реализация интерфейса context.Context
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *Context) Err() error {
	return c.ctx.Err()
}

func (c *Context) SendToIP(ip string, data []byte) error {
	client, err := c.server.getClient(ip)
	if err != nil {
		return fmt.Errorf("c.server.getClient: %w", err)
	}

	client.IncomingChan <- data

	return nil
}

func (c *Context) SendToAll(data []byte) error {
	c.server.broadcastCh <- data

	return nil
}
