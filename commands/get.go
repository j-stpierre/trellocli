package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"trellocli/config"
)

type cardresponse struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Labels []label `json:"labels"`
}

type label struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func Get(cfg config.GetConfig) {
	fmt.Println("Subcommand: get")
	fmt.Printf("Label: %s", cfg.Label)

	credentials := GetCredentials()

	getCards(credentials.BoardId, credentials.APIKey, credentials.Token)
	//Filter by label
	//Save CSV
}

func getCards(boardid string, apikey string, apitoken string) {
	// https://api.trello.com/1/boards/{BOARDID}/cards?key={APIKEY}&token={APITOKEN}

	url := fmt.Sprintf("https://api.trello.com/1/boards/%s/cards?key=%s&token=%s", boardid, apikey, apitoken)

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error code from trello call : %d", resp.StatusCode)
		os.Exit(1)
	}

	var cards []cardresponse

	if err := json.NewDecoder(resp.Body).Decode(&cards); err != nil {
		fmt.Printf("Error decoding JSON: %s", err)
	}

	for _, card := range cards {
		fmt.Printf("Card: %s", card)

	}
}

func filterCards() {

}

func saveCSV() {

}
