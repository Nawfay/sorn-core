package deezer

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sorn/core/models"

	"github.com/gofiber/fiber/v2"
)


func DeezerSearch(c *fiber.Ctx) error {
	name := c.Params("name")

	tracks, err := DeezerSearchTracks(name, 4)
	if err != nil {
		return err
	}

	albums, err := DeezerSearchAlbums(name, 4)
	if err != nil {
		return err
	}

	if len(tracks) == 0 {
		tracks = []models.Track{}
	}

	if len(albums) == 0 {
		albums = []models.Album{}
	}

	searchResult := map[string]interface{}{
		"tracks": tracks,
		"albums": albums,
	}

	return c.JSON(searchResult)	
}

	

func DeezerSearchTracks(name string, limit int) ([]models.Track, error) {
	url := deezerApi + "search?limit=" + fmt.Sprint(limit) + "&q=" + name

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch tracks: " + resp.Status)
	}

	// Decode the response body into the TrackResponse struct
	var trackResponse models.TrackResponse
	err = json.NewDecoder(resp.Body).Decode(&trackResponse)
	if err != nil {
		return nil, err
	}

	var tracks []models.Track
	for _, track := range trackResponse.Data {
		tracks = append(tracks, models.Track{
			ID:       track.ID,
			Title:    track.Title,
			Duration: track.Duration,
			Artist:   track.Artist.Name,
			Album:    track.Album.Title,
			AlbumID:  track.Album.ID,
			AlbumCover: track.Album.Cover,
		})
	}

	return tracks, nil
	

}


func DeezerSearchAlbums(name string, limit int) ([]models.Album, error) {
	// Construct the API URL
	url := deezerApi + "search/album?limit=" + fmt.Sprint(limit) + "&q=" + name

	// Perform the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for non-200 HTTP responses
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch albums: " + resp.Status)
	}

	// Decode the response body into the AlbumResponse struct
	var albumResponse models.AlbumResponse
	err = json.NewDecoder(resp.Body).Decode(&albumResponse)
	if err != nil {
		return nil, err
	}

	var albums []models.Album

	for _, album := range albumResponse.Data {
		
		tracks := extractAlbumTracks(album)

		albums = append(albums, models.Album{
			ID:       album.ID,
			Title:    album.Title,
			Artist:   album.Artist.Name,
			ArtistID: album.Artist.ID,
			Cover:    album.Cover,
			CoverBig: album.CoverBig,
			NbTracks: album.NumberOfTracks,
			Tracks:   tracks,
		})
	}

	return albums, nil
}	
