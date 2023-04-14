package main

import (
	"os"
	"prompting/internal/prompting"
)

func main() {
	command := prompting.NewPromptingCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
