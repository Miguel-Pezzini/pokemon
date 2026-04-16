package main

type Item struct {
	ID   int
	Name string
	Type ItemType
}

var Pokeball = Item{ID: 11, Name: "Poké Ball", Type: ItemPokeBalls}
