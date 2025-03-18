package main

import (
	"sorn/core/external/deezer"
	"sorn/core/logic"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/logger"
	// "log"
)

func SearchRoutes(app *fiber.App) {
	api := app.Group("/api")

	// api.Get("/search/:name", deezer.SearchForNewTracks)
	api.Get("/album/:album_id", deezer.GetAlbum)
	api.Get("/artist/:artist_id", deezer.GetArtist)	
	api.Get("/search/:name", deezer.DeezerSearch)
	api.Get("/download/:album_id", logic.DownloadAlbum)
}
