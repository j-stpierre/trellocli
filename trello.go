package main

import (
	"flag"
	"fmt"
	"os"
	"trellocli/commands"
	"trellocli/config"
)

func main() {

	credentialsCmd := flag.NewFlagSet("set-credentials", flag.ExitOnError)
	boardPtr := credentialsCmd.String("boardid", "", "The board ID for your calls (Required)")
	apikeyPtr := credentialsCmd.String("apikey", "", "The api key for your calls (Required)")
	tokenPtr := credentialsCmd.String("token", "", "The api token for your calls (Required)")

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	labelPtr := getCmd.String("label", "", "Labels to filter by")
	//outputPtr := getCmd.String("output", "csv", "Format to output results as")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	filePtr := deleteCmd.String("file", "cards.csv", "Path to file with elements to delete")

	if len(os.Args) < 2 {
		fmt.Println("Expected sub-command")

		os.Exit(1)
	}

	switch os.Args[1] {

	case "set-credentials":
		credentialsCmd.Parse(os.Args[2:])
		cfg := config.CredentialConfig{BoardId: *boardPtr, APIKey: *apikeyPtr, Token: *tokenPtr}
		commands.SetCredentials(cfg)
	case "get":
		getCmd.Parse(os.Args[2:])
		cfg := config.GetConfig{Label: *labelPtr}
		commands.Get(cfg)
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		cfg := config.DeleteConfig{File: *filePtr}
		commands.Delete(cfg)
	case "-h", "--help":
		fmt.Printf("\nset-credentials: \n")
		credentialsCmd.PrintDefaults()
		fmt.Printf("\nget: \n")
		getCmd.PrintDefaults()
		fmt.Printf("\ndelete: \n")
		deleteCmd.PrintDefaults()
	default:
		os.Exit(1)
	}

}
