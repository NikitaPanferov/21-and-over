package main

import (
	"fmt"

	controllerPkg "github.com/NikitaPanferov/21-and-over/server/internal/controller"
	"github.com/NikitaPanferov/21-and-over/server/internal/domain/services/game"
	"github.com/NikitaPanferov/21-and-over/server/internal/repository/sqlite"
	"github.com/NikitaPanferov/21-and-over/server/pkg/logger"
	tcpserver "github.com/NikitaPanferov/21-and-over/server/pkg/tcp-server"
)

func main() {
	server := tcpserver.NewServer("localhost:9000")
	logger.MustInitGlobal(&logger.Config{Level: "debug"})

	//TODO: вынести в конфиг
	playerRepo := sqlite.New(1000)

	//TODO: вынести в конфиг
	gameService := game.New(2, playerRepo)

	controller := controllerPkg.New(gameService)

	controllerPkg.RegisterHandlers(server, controller)

	if err := server.Start(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
