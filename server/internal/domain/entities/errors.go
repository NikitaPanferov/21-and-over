package entities

import "errors"

var (
	ErrGameIsFull                = errors.New("game is full")
	ErrPlayerAlreadyInGame       = errors.New("player already in game")
	ErrPlayerNotFound            = errors.New("player not found")
	ErrNotEnoughChips            = errors.New("not enough chips")
	ErrPlayerHasAlreadyDidBet    = errors.New("player has already did bet")
	ErrNotYourTurn               = errors.New("not your turn")
	ErrDeckEmpty                 = errors.New("deck is empty")
	ErrPlayersHandIsDone         = errors.New("player`s hand is done")
	ErrPlayerShouldBeAlreadyDone = errors.New("player should be already done")
)
