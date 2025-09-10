package entities

type State string

const (
	JoinState    State = "joining"
	WaitBetState State = "waiting bet"
	PlayState    State = "playing"
	ResultState  State = "result"
)
