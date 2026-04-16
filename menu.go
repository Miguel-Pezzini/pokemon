package main

import (
	"fmt"
	"os"
)

type Character struct {
	Name     string
	Sex      bool
	Pokemons []Pokemon
	Bag      []Item
	Money    int
}

func showMenu() {
	fmt.Println("===============================")
	fmt.Println("              MENU             ")
	fmt.Println("===============================")
	fmt.Println("1. PLAY")
	fmt.Println("2. ABOUT ME")
	fmt.Println("3. LEAVE")
	fmt.Print("SELECT AN OPTION(1/2/3)")
}

func menuOut() {
	cls()
	fmt.Print("LEAVING...")
}

func menuAbout() {
	cls()
	fmt.Println("Project made by Miguel Pezzini with the idea to replicate the game Pokemon from Nintendo using ASCII characters")
}

func menuInGame(character Character) {
	cls()
	opt := 0
	inMenu := true
	sizeVector := 0
	for inMenu {
		fmt.Println("===============================")
		fmt.Println("              MENU             ")
		fmt.Println("===============================")
		fmt.Println("1. POKEMONS")
		fmt.Println("2. BAG")
		fmt.Println("3. CANCEL")
		fmt.Println("4. LEAVE GAME")
		fmt.Scan(&opt)
		switch opt {
		case 1:
			sizeVector = len(character.Pokemons)
			cls()
			if sizeVector >= 1 {
				fmt.Println("LIST OF POKEMONS: ")
				for i := 0; i < sizeVector; i++ {
					fmt.Printf("%s Level: %d\n", character.Pokemons[i].Name, character.Pokemons[i].Level)
				}
			} else {
				fmt.Println("YOU DON'T HAVE ANY POKEMONS!")
			}
		case 2:
			sizeVector = len(character.Bag)
			cls()
			if sizeVector >= 1 {
				fmt.Println("YOUR BAG: ")
				for i := 0; i < sizeVector; i++ {
					fmt.Println(character.Bag[i].Name)
				}
			} else {
				fmt.Println("YOU DON'T HAVE ANY ITEMS!")
			}
		case 3:
			return
		case 4:
			os.Exit(0)
		default:
			inMenu = false
		}
	}
}
