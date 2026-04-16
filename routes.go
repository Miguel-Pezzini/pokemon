package main

type Route struct {
	Levels []int
	IDs    []PokemonID
}

var RouteOne = Route{
	Levels: []int{2, 5},
	IDs:    []PokemonID{10, 11, 13, 14, 16, 19},
}
