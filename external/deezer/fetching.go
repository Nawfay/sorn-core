package deezer


import (
	"encoding/json"
	"errors"
	"net/http"
	"sorn/core/models"

	"github.com/gofiber/fiber/v2"
)


const deezerApi = "https://api.deezer.com/"

func GetAlbum(c *fiber.Ctx) error {
	albumID := c.Params("album_id")
	album, err := FetchAlbum(albumID)
	if err != nil {
		return err
	}
	return c.JSON(album)
}

func GetArtist(c *fiber.Ctx) error {
	artistID := c.Params("artist_id")
	artist, err := FetchArtist(artistID)
	if err != nil {
		return err
	}
	return c.JSON(artist)
}


func FetchAlbum(id string) (models.Album, error) { 

	url := deezerApi + "album/" + id 

	resp, err := http.Get(url)
	if err != nil {
		return models.Album{}, err 
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Album{}, errors.New("failed to fetch albums: " + resp.Status)
	}

	var trackResponse models.AlbumRaw
	err = json.NewDecoder(resp.Body).Decode(&trackResponse)
	if err != nil {
		return models.Album{}, err
	}

	tracks := extractAlbumTracks(trackResponse)


	album := models.Album{
		ID:       trackResponse.ID,
		Title:    trackResponse.Title,
		Artist:   trackResponse.Artist.Name,
		ArtistID: trackResponse.Artist.ID,
		Cover:    trackResponse.Cover,
		CoverBig: trackResponse.CoverBig,
		NbTracks: trackResponse.NumberOfTracks,
		Tracks:   tracks,
	}

	return album, nil
}

func extractAlbumTracks(album models.AlbumRaw) []models.Track {

	url := album.Tracklist

	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	var trackResponse models.AlbumTrackResponse
	err = json.NewDecoder(resp.Body).Decode(&trackResponse)
	if err != nil {
		return nil
	}

	var tracks []models.Track

	for _, track := range trackResponse.Data {

		tracks = append(tracks, models.Track{
			ID:       track.ID,
			Title:    track.Title,
			Duration: track.Duration,
			Artist:   track.Artist.Name,
			Album:    album.Title,
			AlbumID:  album.ID,
			AlbumCover: album.Cover,
			
			})
	}

	return tracks
}

func FetchArtist(id string) (models.Artist, error) {


	url := deezerApi + "artist/" + id 

	resp, err := http.Get(url)
	if err != nil {
		return models.Artist{}, err 
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Artist{}, errors.New("failed to fetch albums: " + resp.Status)
	}

	var trackResponse models.ArtistRaw
	err = json.NewDecoder(resp.Body).Decode(&trackResponse)
	if err != nil {
		return models.Artist{}, err
	}

	artist := models.Artist{
		ID:       trackResponse.ID,
		Name:    trackResponse.Name,
		NbAlbums: trackResponse.NbAlbums,
		PictureBig: trackResponse.PictureBig,
		PictureMedium: trackResponse.PictureMedium,
		PictureSmall: trackResponse.PictureSmall,
		PictureXL: trackResponse.PictureXL,
	}

	return artist, nil

}