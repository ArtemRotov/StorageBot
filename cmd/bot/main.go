package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sadbard/StorageBot/internal/app/client"
	"github.com/sadbard/StorageBot/internal/app/commands"
	"github.com/sadbard/StorageBot/internal/service/keyboard"
	"github.com/sadbard/StorageBot/internal/storage/postgres"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")

	tgClient := client.NewClient(token)

	dataBase := postgres.NewDataBase()
	defer dataBase.Close()

	keyboardService := keyboard.NewService()

	commander := commands.NewCommander(tgClient, keyboardService, dataBase)

	updates := tgClient.GetUpdatesChan()

	for update := range updates {
		commander.HandleUpdate(&update)
	}
}
