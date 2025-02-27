package commands

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"trellocli/config"
)

type card struct {
	id   string
	name string
}

func Delete(cfg config.DeleteConfig) {
	fmt.Println("Subcommand: delete")
	fmt.Printf("File: %s\n", cfg.File)
	cards := readCSV(cfg.File)

	fmt.Println(cards)

	credentials := GetCredentials()

	deleteCards(credentials.APIKey, credentials.Token, cards)

}

func deleteCards(apikey string, apitoken string, cards []card) {

	for _, card := range cards {

		client := &http.Client{}

		url := fmt.Sprintf("https://api.trello.com/1/cards/%s?key=%s&token=%s", card.id, apikey, apitoken)

		req, err := http.NewRequest("DELETE", url, nil)

		if err != nil {
			fmt.Println("Error building delete request: ", err)
		}

		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("Error sending DELETE request: ", err)
		}

		defer resp.Body.Close()

	}

}

func readCSV(filename string) []card {

	csvFile, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening CSV file: ", err)
		os.Exit(1)
	}

	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()

	if err != nil {
		fmt.Println("Error parsing CSV file: ", err)
		os.Exit(1)
	}

	var cards []card

	for _, record := range records {
		cards = append(cards, card{record[0], record[1]})
	}

	cards = cards[1:]

	return cards
}
