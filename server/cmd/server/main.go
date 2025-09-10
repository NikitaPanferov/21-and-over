package main

import (
	"fmt"
	"os"
	"strconv"

	controllerPkg "github.com/NikitaPanferov/21-and-over/server/internal/controller"
	"github.com/NikitaPanferov/21-and-over/server/internal/domain/services/game"
	"github.com/NikitaPanferov/21-and-over/server/internal/repository/json"
	"github.com/NikitaPanferov/21-and-over/server/pkg/logger"
	tcpserver "github.com/NikitaPanferov/21-and-over/server/pkg/tcp-server"
)

func main() {
	serverAddress := mustGetEnv("SERVER_ADDRESS")
	logLevel := mustGetEnv("LOG_LEVEL")
	initialChips := mustGetEnvAsInt("INITIAL_CHIPS")
	playersFilePath := mustGetEnv("PLAYERS_FILE_PATH")
	maxPlayers := mustGetEnvAsInt("MAX_PLAYERS")

	server := tcpserver.NewServer(serverAddress)
	logger.MustInitGlobal(&logger.Config{Level: logLevel})

	playerRepo := json.New(initialChips, playersFilePath)
	gameService := game.New(maxPlayers, playerRepo)

	controller := controllerPkg.New(gameService)
	controllerPkg.RegisterHandlers(server, controller)

	if err := server.Start(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func mustGetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Sprintf("Environment variable %s is required but not set", key))
	}
	return value
}

func mustGetEnvAsInt(key string) int {
	valueStr := mustGetEnv(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		panic(fmt.Sprintf("Environment variable %s must be an integer, got: %s", key, valueStr))
	}
	return value
}
