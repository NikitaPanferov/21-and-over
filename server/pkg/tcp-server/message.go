package tcpserver

type Code uint16

const (
	CodeSuccess Code = 20

	CodeClientError Code = 40
	CodeNotFound    Code = 44

	CodeServerError Code = 50
)

type Message struct {
	ID     string
	Action string
	Code   Code
	Data   any
}
