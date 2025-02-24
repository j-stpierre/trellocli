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
	fmt.Printf("Label: %s\n", cfg.Label)

	credentials := GetCredentials()

	cards := getCards(credentials.BoardId, credentials.APIKey, credentials.Token)

	if cfg.Label != "" {
		cards = filterCards(cards, cfg.Label)
	}

	for _, card := range cards {
		fmt.Printf("Card: %s", card)
	}

	//Save CSV
}

func getCards(boardid string, apikey string, apitoken string) []cardresponse {

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

	return cards
}

func filterCards(cards []cardresponse, filter string) (filtered []cardresponse) {

	for _, card := range cards {
		for _, label := range card.Labels {
			if label.Name == filter {
				filtered = append(filtered, card)
				break
			}
		}
	}

	return filtered
}

func saveCSV() {

}
