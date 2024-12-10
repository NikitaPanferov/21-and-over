package tcpclient

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (c *Client) Join(name string) error {
	request := JoinRequest{
		Name: name,
	}

	response, err := c.SendMessage(Message{
		ID:     uuid.NewString(),
		Action: "JOIN",
		Data:   request,
	})
	if err != nil {
		return fmt.Errorf("c.SendMessage: %w", err)
	}

	if response.Code != CodeSuccess {
		errorMap := response.Data.(map[string]interface{})
		log.Println(errorMap["error"])

		return fmt.Errorf("unexpected response code: %d", response.Code)
	}

	return nil
}
