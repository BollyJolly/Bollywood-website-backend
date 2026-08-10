package main

import (
	"log"

	"github.com/amanhasnainy/bingo-backend/internal/config"
	"github.com/amanhasnainy/bingo-backend/internal/database"
	"github.com/amanhasnainy/bingo-backend/internal/song"
	"github.com/amanhasnainy/bingo-backend/routes"
)

func main() {

	config.LoadEnv()

	database.ConnectMongo()

	err := song.LoadSongs(
		"internal/song/songs.json",
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf(
		"✅ Loaded %d songs",
		len(song.Songs),
	)

	router := routes.SetupRoutes()

	router.Run(
		":" + config.GetEnv("PORT"),
	)
}