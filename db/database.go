package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"sorn/core/models"
)



var db *gorm.DB

func DBInit() (error) {
	var err error
	db, err = gorm.Open(sqlite.Open("sorn.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&models.TrackDB{}, &models.ArtistDB{}, &models.AlbumDB{})	
	
	initQueue()

	
	fmt.Println("DB initialized")
	return nil
}

func DBClose() {
	SQLdb, err := db.DB()
	if err != nil {
		panic(err)
	}

	fmt.Println("DB closed")
	SQLdb.Close()
}



func GetArtist(id uint) (models.ArtistDB, error) {
	fmt.Println("Fetching artist from DB:", id)

	// DBInit()

	if db == nil {
		fmt.Println("DB not initialized")
	}

	var artist models.ArtistDB
	result := db.First(&artist, id)
	if result.Error != nil {
		fmt.Println("Error fetching artist from DB:", result.Error)
		return models.ArtistDB{}, result.Error
	}
	fmt.Println("Fetched artist from DB:", artist)
	return artist, nil
}

func CreateArtist(artist models.ArtistDB) error{
	
	fmt.Println("adding artist to db")
	res := db.Create(&artist)

	if res.Error != nil {
		return res.Error
	}

	return nil

}

func GetAlbums(id uint) models.AlbumDB {
	var albums models.AlbumDB
	db.First(&albums, id)
	return albums
}

func CreateAlbum(album models.AlbumDB) error{
	
	res := db.Create(&album)	

	if res.Error != nil {	
		return res.Error
	}

	return nil

}

