package commands

import (
	"fmt"
	"trellocli/config"
)

func Delete(cfg config.DeleteConfig) {
	fmt.Println("Subcommand: delete")
	fmt.Printf("File: %s", cfg.File)
}
