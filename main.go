package main

import (
	"flag"
	"fmt"
	"github.com/ezkripke/jokeCLI/client"
)

func main() {
	// flags
	c := flag.String(
		"c",
		string(client.DefaultCategory),
		"Joke category (Any, Miscellaneous, Programming, Dark)",
	)

	flag.Parse()

	jokeAPIclient := client.NewJokeAPIClient()

	joke, err := jokeAPIclient.Fetch(client.JokeCategory(*c))
	if err != nil {
		fmt.Printf("An error occured fetching your comic:\n%s\n", err)
	}

	fmt.Println(joke.Fulljoke)
}
