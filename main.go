package main

import (
	"cli-blackjack/cmd"
	"cli-blackjack/database"
	"log"
	"math/rand"
	"time"
)

func main() {
	database.Init()
	defer database.Shutdown()

	database.FetchStats()
	defer func() {
		err := database.SaveStats()
		if err != nil {
			log.Fatal("Error saving stats.")
		}
	}()

	rand.Seed(time.Now().UnixNano())

	cmd.Execute()
}
