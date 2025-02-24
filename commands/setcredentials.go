package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"trellocli/config"
)

func SetCredentials(cfg config.CredentialConfig) {

	if cfg.BoardId == "" || cfg.APIKey == "" || cfg.Token == "" {
		fmt.Printf("Credentials missing: Board Id, API Key and Token must not be empty")
		os.Exit(1)
	}

	credentials := map[string]string{
		"boardid": cfg.BoardId,
		"apikey":  cfg.APIKey,
		"token":   cfg.Token,
	}

	jsonData, err := json.MarshalIndent(credentials, "", "	")
	if err != nil {
		fmt.Printf("Could not marshal credentials: %s", err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory", err)
	}

	fileName := "trellocredentials.json"
	dirName := filepath.Join(homeDir, ".trellocli")
	filePath := filepath.Join(dirName, fileName)

	err = os.MkdirAll(dirName, 0755)
	if err != nil {
		fmt.Println("Error creating directory: ", err)
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error Writing credential file: ", err)
		os.Exit(1)
	}

}

func GetCredentials() config.CredentialConfig {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory", err)
	}

	fileName := filepath.Join(homeDir, ".trellocli/trellocredentials.json")

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error retrieving credentials: ", err)
	}

	var credential config.CredentialConfig
	err = json.Unmarshal(data, &credential)
	if err != nil {
		fmt.Println("Error reading credential file: ", err)
	}

	return credential
}
