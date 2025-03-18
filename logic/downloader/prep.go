package downloader

import (
	"fmt"
	"sorn/core/db"
	"sorn/core/logic/util"
	"sorn/core/models"
	"sorn/core/external/deezer"
)

func DownloadPrep(albumID int) (bool, error) {

	album, err := deezer.FetchAlbum(fmt.Sprintf("%d", albumID))

	if err != nil {
		return false, err
	}

	artistExists(album.ArtistID)
	albumExists(album.ID)

	return true, nil

}


func artistExists(artistID int) bool {

	artistIDUint := uint(artistID)

	_, err := db.GetArtist(artistIDUint)
	if err != nil {
		makeArtist(artistID)
	}
	return true
}


func makeArtist(id int) error{

	artist, err := deezer.FetchArtist(fmt.Sprintf("%d", id))
	if err != nil {
		return err
	}

	folderName, err := util.GenerateArtistFolder(artist.Name, artist.PictureMedium)
	if err != nil {
		return err
	}

	artistDb := models.ArtistDB{
		ID:          uint(artist.ID),
		Tracked:     false,
		Youtube:     false,
		Name:        artist.Name,
		ImagePath:    fmt.Sprintf("%s/%s.jpg", folderName, "artist"),
		ImageLink:   artist.PictureMedium,
		SmallImgLink: artist.PictureSmall,
		NbAlbums:    artist.NbAlbums,
		Path:         folderName,
		
	}

	db.CreateArtist(artistDb)

	return nil

}	




func albumExists(albumID int) bool {

	albumIDUint := uint(albumID)

	album := db.GetAlbums(albumIDUint)
	if album.ID == 0 {
		MakeAlbum(albumID)
	}
	return true
}


func MakeAlbum(id int) error{

	
	album, err := deezer.FetchAlbum(fmt.Sprintf("%d", id))
	if err != nil {
		return err
	}

	artist, err := db.GetArtist(uint(album.ArtistID))

	if err != nil {
		return err
	}

	folderName, err := util.GenerateAlbumFolder(artist.Path, album.Title, album.CoverBig)
	if err != nil {
		return err
	}


	albumDb := models.AlbumDB{
		ID:          uint(album.ID),
		Tracked:     false,
		Youtube:     false,
		Name:        album.Title,
		ImgPath:      fmt.Sprintf("%s/%s.jpg", folderName, "cover"),
		ImgLink:     album.Cover,
		ImgSmallLink: album.Cover,
		ArtistID:    artist.ID,
		ArtistPath:   artist.Path,
		NbTracks:    album.NbTracks,
		Path:         folderName,
	}
	
	db.CreateAlbum(albumDb)
	return nil
}	



