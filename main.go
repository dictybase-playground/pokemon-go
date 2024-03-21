package main

import (
	"kevintun95/pokebrowser/pokequery"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "Pokefetch",
		Commands: []*cli.Command{
			{
				Name: 	 "list",
				Aliases: []string{"l"},
				Usage: 	 "list pokemon",
				Action: pokequery.ListPokemon,
			},
			{
				Name: 	 "get",
				Aliases: []string{"g"},
				Usage: 	 "get pokemon by id",
				Action: pokequery.PokemonById,
			},
			{
				Name: 	 "type",
				Aliases: []string{"t"},
				Usage: 	 "List pokemon by type",
				Action: pokequery.PokemonByType,
			},
		},
	}
	app.Run(os.Args)
}	