package ripper

import (
	"encoding/json"
	"log"
	"os"
)

type DeezerLogin struct {
	ARLCookie      string
	LicenseToken   string
	BlowfishSecret string
	BlowfishIV     string
}


func setCredentials(credentials DeezerLogin) {
	data, err := json.Marshal(credentials)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	err = os.WriteFile("cred.json", data, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}



func getCredentials() DeezerLogin {
	data, err := os.ReadFile("cred.json")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var credentials DeezerLogin
	err = json.Unmarshal(data, &credentials)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	return credentials
}



