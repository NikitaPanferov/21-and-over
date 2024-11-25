package controller

type (
	GameService interface{}
	Transport   interface{}

	Controller struct {
		gameService GameService
		transport   Transport
	}
)

func New(gameService GameService, transport Transport) *Controller {
	return &Controller{
		gameService: gameService,
		transport:   transport,
	}
}


func (c *Controller) RegisterHandlers() {
	// TODO
}