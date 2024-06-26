package pokequery

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"text/template"

	"github.com/urfave/cli/v2"
)

type NamedApiResource struct {
	Name string `json:"name"`
}

type NamedAPIResourceList struct {
	Count int `json:"count"`
	Results []NamedApiResource `json:"results"`
}

type Type struct {
	Name string `json:"name"`
	Pokemon []TypePokemon `json:"pokemon"`
}

type TypePokemon struct {
	Pokemon NamedApiResource `json:"pokemon"`
}
type Pokemon struct {
	Id int `json:"id"`
	Name string `json:"name"`
}


func bodyCloser (resp *http.Response) {
	err := resp.Body.Close()
	if err != nil { 
		log.Fatalln(err)
	}
}

func Increment(i int) int {
	return i + 1
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
	
	var data NamedAPIResourceList
	
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalln(err)
	}
	
	funcMap := template.FuncMap{
		"inc": Increment,
	}
	tmplFile := "listPokemon.tmpl"
	tmpl := template.Must(template.New(tmplFile).Funcs(funcMap).ParseFiles(tmplFile))
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}	

func PokemonById(cCtx *cli.Context) error { 
	id := cCtx.Args().First()
	endPoint := "https://pokeapi.co/api/v2/pokemon/" + id
	
	resp, err := http.Get(endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == 404 {
		log.Fatalln("No Pokemon with Id: " + id)
	}
	
	defer bodyCloser(resp)
	
	var data Pokemon
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	tmplFile := "pokemonById.tmpl"
	tmpl := template.Must(template.New(tmplFile).ParseFiles(tmplFile))
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func PokemonByType(cCtx *cli.Context) error {
	pokemonType := cCtx.Args().First()
	endPoint := "https://pokeapi.co/api/v2/type/" + pokemonType
	
	resp, err := http.Get(endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == 404 {
		log.Fatalln("No Pokemon with type: " + pokemonType)
	}
	
	defer bodyCloser(resp)
	
	var data Type
	
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalln(err)
	}
	
		
	funcMap := template.FuncMap{
		"inc": Increment,
	}
	tmplFile := "pokemonByType.tmpl"
	tmpl := template.Must(template.New(tmplFile).Funcs(funcMap).ParseFiles(tmplFile))
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}