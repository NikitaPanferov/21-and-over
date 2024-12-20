package tcpserver

import "errors"

var (
	ErrGettingIPFromCtx  = errors.New("error getting ip from context")
	ErrConnectionRefused = errors.New("connection refused")
	ErrClientNotFound    = errors.New("client not found")
)
