package tcpserver

import (
	"context"

	"github.com/NikitaPanferov/21-and-over/server/pkg/tcp-server/types"
)

func GetUserIP(ctx context.Context) (string, error) {
	value, ok := ctx.Value(types.ContextKeyUserIP).(string)
	if !ok {
		return "", types.ErrGettingIPFromCtx
	}
	return value, nil
}
