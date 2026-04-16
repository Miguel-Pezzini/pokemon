package main

import (
	"fmt"
)

func speakerBox(speaker string, paragraphs []string, color string) {
	var lines []string
	for i, para := range paragraphs {
		for _, l := range wrapText(para, 62) {
			lines = append(lines, colorWhite+l+ansiReset)
		}
		if i < len(paragraphs)-1 {
			lines = append(lines, "")
		}
	}
	drawTitledBox(speaker, lines, 70, color, colorYellow)
}

func starterCard(num int, name, typeLabel, desc string, typeColorBg string) {
	header := fmt.Sprintf("%s ▸ %d %s │ %s%s%s  %s %s %s",
		colorGold, num, colorGray,
		ansiBold+colorWhite, name, ansiReset,
		typeColorBg, typeLabel, ansiReset)

	lines := []string{header}
	for _, l := range wrapText(desc, 62) {
		lines = append(lines, "    "+colorGray+l+ansiReset)
	}
	drawBox(lines, 70, colorDark)
}

func dialogLabOne(character *Character, firstTry *bool) {
	cls()
	fmt.Println()

	if *firstTry {
		speakerBox("PROFESSOR SILVA", []string{
			"Hello, young adventurer! Welcome to my lab. I am Professor Silva, and I'm here to guide you through the wonderful world of Pokémon.",
			"The Pokémon world is fascinating! Amazing creatures live everywhere — some can fly, others can swim, and many have special powers waiting to be discovered.",
			"To begin your journey, you must choose your very first companion. Pick wisely!",
		}, colorGreen)

		fmt.Println()
		starterCard(1, "Bulbasaur", "GRASS", "A Grass-type Pokémon. Gentle and nature-loving — with it at your side, you will master the power of plants.", typeColor(TypeGrass))
		starterCard(2, "Charmander", "FIRE", "A Fire-type Pokémon. Fearless and full of energy — feel the heat of battle with Charmander by your side.", typeColor(TypeFire))
		starterCard(3, "Squirtle", "WATER", "A Water-type Pokémon. Playful and spirited — choose Squirtle and become a true master of the waters.", typeColor(TypeWater))

		fmt.Println()
		fmt.Print(colorCyan + "  › " + ansiReset + "Which Pokémon will you choose? " + colorGray + "(1/2/3): " + ansiReset)

		opt := 0
		fmt.Scan(&opt)
		switch opt {
		case 1:
			character.Pokemons = append(character.Pokemons, createPokemon(1, 5))
			*firstTry = false
			confirmStarter("Bulbasaur")
		case 2:
			character.Pokemons = append(character.Pokemons, createPokemon(4, 5))
			*firstTry = false
			confirmStarter("Charmander")
		case 3:
			character.Pokemons = append(character.Pokemons, createPokemon(7, 5))
			*firstTry = false
			confirmStarter("Squirtle")
		default:
			fmt.Println()
			fmt.Println(colorRed + "  ! You must choose a Pokémon!" + ansiReset)
			pressEnterToContinue()
		}
	} else {
		speakerBox("PROFESSOR SILVA", []string{
			"Head out into the world and become a great Pokémon trainer! I believe in you.",
		}, colorGreen)

		fmt.Println()
		fmt.Print(colorCyan + "  › " + ansiReset + "Press " + colorYellow + "1" + ansiReset + " to leave: ")
		opt := 0
		fmt.Scan(&opt)
	}
}

func confirmStarter(name string) {
	fmt.Println()
	drawTitledBox("✦ NEW PARTNER ✦", []string{
		"",
		centerText(fmt.Sprintf("%sYou received %s%s%s!%s",
			colorWhite, ansiBold+colorYellow, name, ansiReset+colorWhite, ansiReset), 52),
		centerText(colorGray+"Take good care of it."+ansiReset, 52),
		"",
	}, 60, colorGold, colorYellow)
	pressEnterToContinue()
}
