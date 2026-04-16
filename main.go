package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"golang.org/x/term"
)

type GameState int

const (
	INITIAL_HOUSE GameState = iota
	MAP_ONE
	LAB_ONE
)

func getch() byte {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		var buf [1]byte
		_, _ = os.Stdin.Read(buf[:])
		return buf[0]
	}
	defer term.Restore(fd, oldState)

	var buf [1]byte
	_, err = os.Stdin.Read(buf[:])
	if err != nil {
		return 0
	}
	return buf[0]
}

func cls() {
	fmt.Print("\033[2J\033[H")
}

func makeMatrix(rows, cols int) [][]int {
	m := make([][]int, rows)
	for i := range m {
		m[i] = make([]int, cols)
	}
	return m
}

func returnCP(currentState GameState, x, y *int, character *Character) GameState {
	size := len(character.Pokemons)
	for i := 0; i < size; i++ {
		character.Pokemons[i].ActualHP = character.Pokemons[i].HP
	}
	switch currentState {
	case MAP_ONE:
		*x = 5
		*y = 2
		currentState = INITIAL_HOUSE
	}
	return currentState
}

func gameRunning(character *Character) {
	houseMapMat := makeMatrix(7, 10)
	mapMat := makeMatrix(20, 20)
	labMat := makeMatrix(15, 15)

	teamDead := false

	x := 2
	y := 6

	firstDialogWithProfessor := true

	cityOne := true

	currentState := INITIAL_HOUSE

	var key byte

	for cityOne {
		optionPath := 0
		cls()

		switch currentState {
		case INITIAL_HOUSE:
			houseMap(houseMapMat, x, y)
			seeMap(houseMapMat)
			for {
				key = getch()
				optionPath = movement(houseMapMat, key, &x, &y, *character)
				if optionPath == 3 {
					currentState = MAP_ONE
					x = 16
					y = 15
					cls()
					break
				}
				if optionPath == 2 {
					break
				}
			}
		case MAP_ONE:
			mapOne(mapMat, x, y)
			seeMap(mapMat)
			for {
				key = getch()
				optionPath = movement(mapMat, key, &x, &y, *character)
				if optionPath == 3 {
					currentState = INITIAL_HOUSE
					x = 5
					y = 2
					cls()
					break
				}
				if optionPath == 5 {
					currentState = LAB_ONE
					x = 13
					y = 7
					cls()
					break
				}
				if optionPath == 2 {
					break
				}
				if optionPath == 7 {
					teamDead = inCombatInBush(character, 1)
					if teamDead {
						currentState = returnCP(currentState, &x, &y, character)
					}
					break
				}
			}
		case LAB_ONE:
			labOne(labMat, x, y)
			seeMap(labMat)
			for {
				key = getch()
				optionPath = movement(labMat, key, &x, &y, *character)
				if optionPath == 3 {
					currentState = MAP_ONE
					x = 16
					y = 5
					cls()
					break
				}
				if optionPath == 2 {
					break
				}
				if optionPath == 6 {
					dialogLabOne(character, &firstDialogWithProfessor)
					break
				}
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	optionMenu := 0
	isRunning := true
	character := Character{}

	for isRunning {
		showMenu()
		fmt.Scan(&optionMenu)

		switch optionMenu {
		case 1:
			gameRunning(&character)
		case 2:
			menuAbout()
		case 3:
			menuOut()
			isRunning = false
		}
	}
}
