package main

import (
	"github.com/neet-007/pokeapi"
	"os"
)

func exitCommand(config *pokeapi.Config, args ...string) error {
	os.Exit(0)
	return nil
}
