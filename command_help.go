package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	commands := getCommands(cfg)
	names := []string{"help", "map", "mapb", "exit", "explore", "catch", "inspect", "pokedex"}

	for _, name := range names {
		if cmd, ok := commands[name]; ok {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		}
	}

	fmt.Println()
	return nil
}



