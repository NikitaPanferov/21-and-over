package tcpserver

import (
	"context"
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
func (c *Context) Write(data []byte) error {
	// TODO: обработка ошибок
	// TODO: мб принимать не байты
	return c.server.SendToIP(c.sender, data)
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