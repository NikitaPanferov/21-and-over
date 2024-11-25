package tcpserver

// Handler представляет функцию-обработчик для сообщений.
type Handler func(ctx *Context) error
