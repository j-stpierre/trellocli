package commands

import (
	"encoding/csv"
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

	saveCSV(cards)
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

func saveCSV(cards []cardresponse) {

	saves := [][]string{
		{"id", "name", "labels"},
	}

	for _, card := range cards {

		labels := ""
		if len(card.Labels) > 0 {

			for _, label := range card.Labels {
				labels = fmt.Sprintf("%s, %s", labels, label.Name)
			}
			labels = labels[2:]

		}

		entry := []string{card.ID, card.Name, labels}
		saves = append(saves, entry)
	}

	csvFile, err := os.Create("cards.csv")

	if err != nil {
		fmt.Println("Error creating CSV file: ", err)
	}

	w := csv.NewWriter(csvFile)
	w.WriteAll(saves)

	if err := w.Error(); err != nil {
		fmt.Println("Error writing to csv file: ", err)
		os.Exit(1)
	}
	csvFile.Close()

}
