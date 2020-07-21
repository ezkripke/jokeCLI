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
	t := flag.String(
		"t",
		string(client.DefaultType),
		"Joke Type (Single, Twopart)",
	)
	flag.Parse()

	jokeAPIclient := client.NewJokeAPIClient()

	joke, err := jokeAPIclient.Fetch(client.JokeCategory(*c), client.JokeType(*t))
	if err != nil {
		fmt.Printf("An error occured fetching your comic:\n%s\n", err)
	}
	fmt.Println(joke.Fulljoke)
}
