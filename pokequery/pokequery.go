package pokequery

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/urfave/cli/v2"
)

type NamedApiResource struct {
	Name string `json:"name"`
}

type ListPokemonResponseBodyData struct {
	Count int `json:"count"`
	Results []NamedApiResource `json:"results"`
}

type PokemonOfType struct {
	Type string `json:"name"`
	Pokemon []TypePokemon `json:"pokemon"`
}

type TypePokemon struct {
	Pokemon NamedApiResource `json:"pokemon"`
}
type Pokemon struct {
	Id int `json:"id"`
	Name string `json:"name"`
}



func mapPokemonNames(pokemonResults []NamedApiResource) []string {
	names := make([]string, len(pokemonResults))
	for i, e := range pokemonResults {
		names[i] = e.Name
	}
	return names
}

func bodyCloser (resp *http.Response) {
	err := resp.Body.Close()
	if err != nil { 
		log.Fatalln(err)
	}
}

func ListPokemon(cCtx *cli.Context) error { 	
	params := url.Values{}
	params.Add("limit", cCtx.Args().First())
	endPoint := "https://pokeapi.co/api/v2/pokemon/?" + params.Encode()

	resp, err := http.Get(endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	defer bodyCloser(resp)
	
	var data ListPokemonResponseBodyData
	
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(mapPokemonNames(data.Results))

	return nil
}	

func PokemonById(cCtx *cli.Context) error { 
	endPoint := "https://pokeapi.co/api/v2/pokemon/" + cCtx.Args().First()
	
	resp, err := http.Get(endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	defer bodyCloser(resp)
	
	var data Pokemon
	
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)

	return nil
}

func PokemonByType(cCtx *cli.Context) error {
	endPoint := "https://pokeapi.co/api/v2/type/" + cCtx.Args().First()
	
	resp, err := http.Get(endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	defer bodyCloser(resp)
	
	var data PokemonOfType
	
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)

	return nil
}