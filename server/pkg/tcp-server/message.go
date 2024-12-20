package tcpserver

type Code uint16

const (
	CodeSuccess Code = 20

	CodeClientError Code = 40
	CodeNotFound    Code = 44

	CodeServerError Code = 50
)

type Action string

const (
	ActionJoin  Action = "JOIN"
	ActionBet   Action = "BET"
	ActionHit   Action = "HIT"
	ActionStand Action = "STAND"
)

type EventType string

const (
	EventTypePlayerJoined EventType = "PLAYER_JOINED"
	EventTypeWaitingBet   EventType = "WAITING_BET"
	EventTypeGameStarted  EventType = "GAME_STARTED"
	EventTypeGameEnded    EventType = "GAME_ENDED"
)

type Message struct {
	ID        string
	Action    Action
	EventType EventType
	Code      Code
	Data      any
}
