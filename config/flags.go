package config

import (
	"fmt"
	"os"
)

type Config struct {
	validate bool
	stdin    bool
}

func ParseFlags() {
	return nil
}

func ParseCommand(cfg *Config) error {
	if len(os.Args) < 2 {
		return fmt.Errorf("Error: Not enough arguments")
	}
	command := os.Args[1]
	switch command {
	case "validate":

	case "generate":
	case "information":
	case "issue":
	}
	return nil
}
