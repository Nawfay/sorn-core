package logic

import (
	"sorn/core/external/deezer"
	"sorn/core/logic/downloader"

	"github.com/gofiber/fiber/v2"
)

func DownloadAlbum(c *fiber.Ctx) error {
	albumID := c.Params("album_id")

	album, err := deezer.FetchAlbum(albumID)

	if err != nil {
		return err
	}

	running, err := downloader.DownloadPrep(album.ID)

	if err != nil || !running {
		return err
	}

	downloader.UploadAlbum(album)

	c.SendStatus(fiber.StatusOK)

	return nil
}
