package downloader

import (
	"sorn/core/db"
	"sorn/core/models"
	"time"
)


func UploadAlbum(album models.Album) error {

	for _, tracks := range album.Tracks {

		trackQueue := models.QueueItem{
			TrackID: uint(tracks.ID),
			TrackName: tracks.Title,
			CreatedAt: time.Now(),
		}
		db.AddToQueue(trackQueue)
		
	}

	return nil
	
}