package main

import (
	_ "chezmoi-tui/pkg/commands"
	"log"

	"chezmoi-tui/pkg/root"
)

func main() {
	if err := root.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
