package main

import "fmt"

func dialogLabOne(character *Character, firstTry *bool) {
	opt := 0
	cls()
	if *firstTry {
		fmt.Println("Hello, young adventurer! Welcome to my lab! I am Professor Silva. I am here to help you understand the wonderful world of Pokemon.")
		fmt.Println("The Pokemon world is fascinating! There are amazing creatures everywhere, each with its own abilities and characteristics. Some can fly, others can swim, and many have special powers!")
		fmt.Println("To begin your journey, you will need a Pokemon. Here are three options:")
		fmt.Println("OPTION 1 - Bulbasaur - A Grass-type Pokemon. It is gentle and loves nature. With it, you will learn about the power of plants!")
		fmt.Println("OPTION 2 - Charmander - A Fire-type Pokemon. It is fearless and full of energy. With Charmander by your side, you will feel the heat of battle")
		fmt.Println("OPTION 3 - Squirtle - A Water-type Pokemon. It is playful and loves to swim. If you choose Squirtle, you will become a true master of the waters!")
		fmt.Scan(&opt)
		switch opt {
		case 1:
			character.Pokemons = append(character.Pokemons, createPokemon(1, 5))
			*firstTry = false
		case 2:
			character.Pokemons = append(character.Pokemons, createPokemon(4, 5))
			*firstTry = false
		case 3:
			character.Pokemons = append(character.Pokemons, createPokemon(7, 5))
			*firstTry = false
		default:
			fmt.Print("Please pick a pokemon!")
		}
	} else {
		fmt.Println("Go to your jorney through the world of pokemons!")
		fmt.Print("PRESS 1 TO CANCEL")
		fmt.Scan(&opt)
	}
}
