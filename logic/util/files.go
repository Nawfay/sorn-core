package util

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func sterlizeString(str string) string {
	// Replace any characters that could cause issues in Unix filenames
	// str = strings.ToLower(str)
	// str = strings.ReplaceAll(str, " ", "_")
	str = strings.ReplaceAll(str, "/", "_")
	str = strings.ReplaceAll(str, "\\", "_")
	str = strings.ReplaceAll(str, ":", "_")
	str = strings.ReplaceAll(str, "*", "_")
	str = strings.ReplaceAll(str, "?", "_")
	str = strings.ReplaceAll(str, "\"", "_")
	str = strings.ReplaceAll(str, "<", "_")
	str = strings.ReplaceAll(str, ">", "_")
	str = strings.ReplaceAll(str, "|", "_")
	return str
}

func GenerateAlbumFolder(artistPath string, albumName string, imageUrl string) (string, error) {
	
	albumName = sterlizeString(albumName)

	folderName := fmt.Sprintf("%s/%s", artistPath, albumName)

	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		return "", err
	}

	imagePath := fmt.Sprintf("%s/%s.jpg", folderName, "cover")

	err = downloadImage(imageUrl, imagePath)
	if err != nil {
		return "", err
	}

	return folderName, nil
}


func GenerateArtistFolder(artistName string, imageUrl string) (string, error) {
	
	artistName = sterlizeString(artistName)

	folderName := fmt.Sprintf("%s/%s", "/Users/nawaf/Documents/GitHub/sorn-core/media", artistName)

	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		return "", err
	}

	imagePath := fmt.Sprintf("%s/%s.jpg", folderName, "artist")
	
	err = downloadImage(imageUrl, imagePath)
	if err != nil {
		return "", err
	}


	return folderName, nil
}

func downloadImage(imageUrl string, imagePath string) error {

	fmt.Println("Downloading image:", imageUrl, imagePath)
	resp, err := http.Get(imageUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(imagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.ReadFrom(resp.Body)
	return err	
}

