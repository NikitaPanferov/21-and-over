package tcpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NikitaPanferov/21-and-over/server/pkg/logger"
)

// Context представляет пользовательский контекст, совместимый с context.Context.
type Context struct {
	ctx      context.Context
	server   *Server
	message  *Message
	sender   string
	response *Message
	data     []byte
}

// NewContext создает новый пользовательский контекст.
func NewContext(server *Server) *Context {
	return &Context{
		ctx:    context.Background(),
		server: server,
	}
}

// SetMessage устанавливает сообщение в контекст.
func (c *Context) SetMessage(message *Message) {
	c.message = message
}

// GetMessage возвращает сообщение из контекста.
func (c *Context) GetMessage() *Message {
	return c.message
}

func (c *Context) SetSender(sender string) {
	c.sender = sender
}

func (c *Context) GetSender() string {
	return c.sender
}

// Write отправляет ответ клиенту.
func (c *Context) Reply(code Code, data any) error {
	dataToSend, err := json.Marshal(Message{
		ID:     c.message.ID,
		Action: c.message.Action,
		Code:   code,
		Data:   data,
	})
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	return c.SendToIP(c.sender, dataToSend)
}

func (c *Context) ReplyWithError(code Code, err error) error {
	dataToSend, err := json.Marshal(Message{
		ID:     c.message.ID,
		Action: c.message.Action,
		Code:   code,
		Data: map[string]string{
			"error": err.Error(),
		},
	})
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	return c.SendToIP(c.sender, dataToSend)
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

func (c *Context) SendToAll(code Code, message *Message) {
	dataToSend, err := json.Marshal(message)
	if err != nil {
		logger.ErrorContext(
			c.ctx,
			"failed to marshall message",
			"error", err,
		)
		return
	}

	c.server.broadcastCh <- dataToSend
}

func (c *Context) GetRawData() []byte {
	return c.data
}

func (c *Context) SetRawData(data any) error {
	rawData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	c.data = rawData

	return nil
}

func (c *Context) GetContext() context.Context {
	return c.ctx
}
