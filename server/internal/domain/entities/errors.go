package entities

import "errors"

var (
	ErrJoinGameIsOn            = errors.New("game is on")
	ErrJoinGameIsFull          = errors.New("game is full")
	ErrJoinPlayerAlreadyInGame = errors.New("player already in game")
)
