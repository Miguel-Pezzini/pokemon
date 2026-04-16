package main

import (
	"fmt"
	"os"
	"strings"
)

type Character struct {
	Name     string
	Sex      bool
	Pokemons []Pokemon
	Bag      []Item
	Money    int
}

var pokemonTitle = []string{
	"  ██████╗  ██████╗ ██╗  ██╗███████╗███╗   ███╗ ██████╗ ███╗   ██╗",
	"  ██╔══██╗██╔═══██╗██║ ██╔╝██╔════╝████╗ ████║██╔═══██╗████╗  ██║",
	"  ██████╔╝██║   ██║█████╔╝ █████╗  ██╔████╔██║██║   ██║██╔██╗ ██║",
	"  ██╔═══╝ ██║   ██║██╔═██╗ ██╔══╝  ██║╚██╔╝██║██║   ██║██║╚██╗██║",
	"  ██║     ╚██████╔╝██║  ██╗███████╗██║ ╚═╝ ██║╚██████╔╝██║ ╚████║",
	"  ╚═╝      ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═══╝",
}

func showMenu() {
	cls()
	fmt.Println()
	for _, line := range pokemonTitle {
		fmt.Println(colorGold + line + ansiReset)
	}
	fmt.Println()
	fmt.Println(centerText(colorGray+ansiItalic+"a terminal adventure"+ansiReset, 66))
	fmt.Println()

	lines := []string{
		colorGreen + "  1 " + colorGray + "│ " + colorWhite + "PLAY",
		colorBlue + "  2 " + colorGray + "│ " + colorWhite + "ABOUT",
		colorRed + "  3 " + colorGray + "│ " + colorWhite + "LEAVE",
	}
	drawTitledBox("MAIN MENU", lines, 40, colorGold, colorYellow)
	fmt.Println()
	fmt.Print(colorCyan + "  › " + ansiReset + "Select an option " + colorGray + "(1/2/3): " + ansiReset)
}

func menuOut() {
	cls()
	fmt.Println()
	drawTitledBox("GOODBYE", []string{
		"",
		centerText(colorWhite+"Thanks for playing!"+ansiReset, 32),
		centerText(colorGray+"See you next time, trainer."+ansiReset, 32),
		"",
	}, 40, colorMagenta, colorPink)
	fmt.Println()
}

func menuAbout() {
	cls()
	fmt.Println()
	body := "A terminal-based Pokémon adventure recreated using ASCII art and ANSI colors. " +
		"Made by Miguel Pezzini — inspired by the classic Pokémon games from Nintendo."
	lines := append([]string{""}, wrapText(body, 52)...)
	lines = append(lines, "")
	lines = append(lines, colorGray+"Press Enter to return to the menu."+ansiReset)
	drawTitledBox("ABOUT", lines, 60, colorCyan, colorBlue)
	fmt.Println()

	var dummy string
	fmt.Scanln(&dummy)
	fmt.Scanln(&dummy)
}

func menuInGame(character Character) {
	cls()
	opt := 0
	inMenu := true
	for inMenu {
		cls()
		lines := []string{
			colorGreen + "  1 " + colorGray + "│ " + colorWhite + "POKÉMON",
			colorYellow + "  2 " + colorGray + "│ " + colorWhite + "BAG",
			colorBlue + "  3 " + colorGray + "│ " + colorWhite + "CANCEL",
			colorRed + "  4 " + colorGray + "│ " + colorWhite + "LEAVE GAME",
		}
		fmt.Println()
		drawTitledBox("MENU", lines, 40, colorGold, colorYellow)
		fmt.Println()
		fmt.Print(colorCyan + "  › " + ansiReset + "Choose: " + ansiReset)

		fmt.Scan(&opt)
		switch opt {
		case 1:
			showPokemonList(character)
		case 2:
			showBag(character)
		case 3:
			return
		case 4:
			cls()
			menuOut()
			os.Exit(0)
		default:
			inMenu = false
		}
	}
}

func showPokemonList(character Character) {
	cls()
	fmt.Println()
	size := len(character.Pokemons)
	if size == 0 {
		drawTitledBox("POKÉMON", []string{
			"",
			colorGray + "  You don't have any Pokémon yet." + ansiReset,
			colorGray + "  Visit Professor Silva to get your starter!" + ansiReset,
			"",
		}, 56, colorRed, colorOrange)
	} else {
		var lines []string
		for i, p := range character.Pokemons {
			lead := fmt.Sprintf("%s %d %s", colorGold, i+1, ansiReset)
			name := fmt.Sprintf("%s%-12s%s", colorWhite+ansiBold, p.Name, ansiReset)
			lvl := fmt.Sprintf("%sLv.%s%-3d%s", colorGray, colorYellow, p.Level, ansiReset)
			hpRatio := p.ActualHP
			if p.HP == 0 {
				hpRatio = 0
			}
			hp := fmt.Sprintf("%sHP %s%d%s/%s%d%s", colorGray, colorGreen, hpRatio, colorGray, colorWhite, p.HP, ansiReset)
			lines = append(lines, fmt.Sprintf("%s %s %s  %s  %s  %s", lead, name, typeBadge(p.Type), lvl, hp, ""))
		}
		drawTitledBox("YOUR POKÉMON", lines, 70, colorGreen, colorCyan)
	}
	fmt.Println()
	fmt.Print(colorGray + "  Press Enter to return..." + ansiReset)
	var dummy string
	fmt.Scanln(&dummy)
	fmt.Scanln(&dummy)
}

func showBag(character Character) {
	cls()
	fmt.Println()
	size := len(character.Bag)
	if size == 0 {
		drawTitledBox("BAG", []string{
			"",
			colorGray + "  Your bag is empty." + ansiReset,
			"",
		}, 48, colorYellow, colorOrange)
	} else {
		counts := map[string]int{}
		for _, it := range character.Bag {
			counts[it.Name]++
		}
		var lines []string
		for name, n := range counts {
			dots := strings.Repeat(".", 30-visibleLen(name))
			if dots == "" {
				dots = " "
			}
			lines = append(lines, fmt.Sprintf("  %s%s%s %s%s %sx%d%s",
				colorWhite, name, ansiReset,
				colorDark, dots,
				colorYellow, n, ansiReset))
		}
		drawTitledBox("BAG", lines, 48, colorYellow, colorOrange)
	}
	fmt.Println()
	fmt.Print(colorGray + "  Press Enter to return..." + ansiReset)
	var dummy string
	fmt.Scanln(&dummy)
	fmt.Scanln(&dummy)
}
