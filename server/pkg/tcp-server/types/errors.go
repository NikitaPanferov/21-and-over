package types

import "errors"

var (
	ErrGettingIPFromCtx  = errors.New("error getting ip from context")
	ErrConnectionRefused = errors.New("connection refused")
)
