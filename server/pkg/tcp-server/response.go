package tcpserver

type ResponseCode uint16

const (
	ResponseCodeSuccess ResponseCode = 2

	ResponseCodeClientError ResponseCode = 400
	ResponseCodeNotFound    ResponseCode = 404

	ResponseCodeServerError ResponseCode = 5
)

type Response struct {
	Data any          `json:"data"`
	Code ResponseCode `json:"code"`
}
