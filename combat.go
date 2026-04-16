package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func weakAgainst(t PokemonType) []PokemonType {
	switch t {
	case TypeGrass:
		return []PokemonType{TypeFire, TypeBug, TypeFlying, TypeIce, TypePoison}
	case TypeFire:
		return []PokemonType{TypeWater, TypeGround, TypeRock}
	case TypeWater:
		return []PokemonType{TypeElectric, TypeGrass}
	case TypeBug:
		return []PokemonType{TypeFire, TypeFlying, TypeRock, TypeGhost, TypeFairy}
	case TypeNormal:
		return []PokemonType{TypeFighting}
	case TypePoison:
		return []PokemonType{TypeGround, TypePsychic}
	case TypeElectric:
		return []PokemonType{TypeGround}
	case TypeGround:
		return []PokemonType{TypeWater, TypeIce, TypeGrass}
	case TypeFighting:
		return []PokemonType{TypePsychic, TypeFlying, TypeFairy}
	case TypePsychic:
		return []PokemonType{TypeBug, TypeGhost, TypeDark}
	case TypeRock:
		return []PokemonType{TypeFighting, TypeGround, TypeSteel, TypeWater, TypeGrass}
	case TypeIce:
		return []PokemonType{TypeFighting, TypeRock, TypeSteel, TypeFire}
	case TypeGhost:
		return []PokemonType{TypeGhost, TypeDark}
	case TypeDragon:
		return []PokemonType{TypeIce, TypeDragon, TypeFairy}
	case TypeFairy:
		return []PokemonType{TypePoison, TypeSteel}
	default:
		return []PokemonType{}
	}
}

func strongAgainst(t PokemonType) []PokemonType {
	switch t {
	case TypeGrass:
		return []PokemonType{TypeWater, TypeRock, TypeGround}
	case TypeFire:
		return []PokemonType{TypeGrass, TypeBug, TypeIce}
	case TypeWater:
		return []PokemonType{TypeFire, TypeRock, TypeGround}
	case TypeBug:
		return []PokemonType{TypeGrass, TypePsychic, TypeDark}
	case TypeNormal:
		return []PokemonType{}
	case TypePoison:
		return []PokemonType{TypeGrass, TypeFairy}
	case TypeElectric:
		return []PokemonType{TypeWater, TypeFlying}
	case TypeGround:
		return []PokemonType{TypeElectric}
	case TypeFighting:
		return []PokemonType{TypeNormal, TypeRock, TypeSteel, TypeIce, TypeDark}
	case TypePsychic:
		return []PokemonType{TypeFighting, TypePoison}
	case TypeRock:
		return []PokemonType{TypeFire, TypeIce, TypeFlying, TypeBug}
	case TypeIce:
		return []PokemonType{TypeGrass, TypeGround, TypeFlying, TypeDragon}
	case TypeGhost:
		return []PokemonType{TypePsychic, TypeGhost}
	case TypeDragon:
		return []PokemonType{TypeDragon}
	case TypeFairy:
		return []PokemonType{TypeFighting, TypeDragon, TypeDark}
	default:
		return []PokemonType{}
	}
}

func menuCombat(p Pokemon) int {
	opt := 0
	for opt < 1 || opt > 4 {
		fmt.Println("What do you want to do: ")
		fmt.Println("1. Fight")
		fmt.Println("2. Bag")
		fmt.Println("3. Run")
		fmt.Println("4. Pokemon")
		fmt.Scan(&opt)

		if opt < 1 || opt > 4 {
			fmt.Println("Invalid option. Please choose again.")
		}
	}
	return opt
}

func checkPokemonDead(character Character) bool {
	sumDead := 0
	size := len(character.Pokemons)
	for i := 0; i < size; i++ {
		if character.Pokemons[i].ActualHP <= 0 {
			sumDead++
		}
	}
	if sumDead == size {
		return false
	}
	return true
}

func attackPokemon(character Character, p *Pokemon, attackOpt int) {
	typeModifier := 1.0

	attackType := character.Pokemons[0].Attacks[attackOpt].Type

	strongTypeAgainst := strongAgainst(attackType)
	weakTypeAgainst := weakAgainst(attackType)

	for i := 0; i < len(strongTypeAgainst); i++ {
		if p.Type == strongTypeAgainst[i] {
			typeModifier = 1.5
		}
	}

	for i := 0; i < len(weakTypeAgainst); i++ {
		if p.Type == weakTypeAgainst[i] {
			typeModifier = 0.5
		}
	}

	dano := (((2*p.Level/5+2)*character.Pokemons[0].Attacks[attackOpt].Power*int(typeModifier*50))/p.Def)/50 + 2
	p.ActualHP -= dano
	if p.ActualHP <= 0 {
		p.ActualHP = 0
	}

	if typeModifier == 1.5 {
		fmt.Printf("%s attack is effective doing %d of damage\n", character.Pokemons[0].Name, dano)
	}
	if typeModifier == 0.5 {
		fmt.Printf("%s attack is not effective doing %d of damage\n", character.Pokemons[0].Name, dano)
	}
	if typeModifier == 1.0 {
		fmt.Printf("%s attack did %d of damage\n", character.Pokemons[0].Name, dano)
	}
}

func pokemonAttackYou(character *Character, p Pokemon, randomAttack int) {
	typeModifier := 1.0

	attackType := p.Attacks[randomAttack].Type

	strongTypeAgainst := strongAgainst(attackType)
	weakTypeAgainst := weakAgainst(attackType)

	for i := 0; i < len(strongTypeAgainst); i++ {
		if character.Pokemons[0].Type == strongTypeAgainst[i] {
			typeModifier = 1.5
		}
	}

	for i := 0; i < len(weakTypeAgainst); i++ {
		if character.Pokemons[0].Type == weakTypeAgainst[i] {
			typeModifier = 0.5
		}
	}

	dano := (((2*character.Pokemons[0].Level/5+2)*p.Attacks[randomAttack].Power*int(typeModifier*50))/character.Pokemons[0].Def)/50 + 2
	character.Pokemons[0].ActualHP -= dano

	if character.Pokemons[0].ActualHP <= 0 {
		character.Pokemons[0].ActualHP = 0
	}

	if typeModifier == 1.5 {
		fmt.Printf("%s attack is effective doing %d of damage\n", p.Name, dano)
	}
	if typeModifier == 0.5 {
		fmt.Printf("%s attack is not effective doing %d of damage\n", p.Name, dano)
	} else {
		fmt.Printf("%s attack did %d of damage\n", p.Name, dano)
	}
}

func changePokemon(character *Character) {
	size := len(character.Pokemons)
	indexPoke := 0
	fmt.Println("The Pokémon fainted! You must choose another Pokémon!: ")
	for i := 0; i < size; i++ {
		if character.Pokemons[i].ActualHP > 0 {
			fmt.Printf("PRESS %d TO SELECT THE POKEMON %s", i, character.Pokemons[i].Name)
		}
	}
	fmt.Scan(&indexPoke)
	temp := character.Pokemons[0]
	character.Pokemons[0] = character.Pokemons[indexPoke]
	character.Pokemons[indexPoke] = temp
}

func pressEnterToContinue() {
	fmt.Print("Press Enter to continue...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func fight(p *Pokemon, character *Character) bool {
	attackOpt := 0
	attacksWildLenght := len(p.Attacks)
	randomAttack := rand.Intn(attacksWildLenght)

	fmt.Print("YOUR ATTACKS: \n")
	for i := 0; i < len(character.Pokemons[0].Attacks); i++ {
		fmt.Printf("%d. %s", i+1, character.Pokemons[0].Attacks[i].Name)
	}
	fmt.Scan(&attackOpt)
	attackOpt--

	if character.Pokemons[0].Speed > p.Speed {
		attackPokemon(*character, p, attackOpt)
		if p.ActualHP <= 0 {
			return false
		}
		pokemonAttackYou(character, *p, randomAttack)
		if character.Pokemons[0].ActualHP <= 0 {
			if !checkPokemonDead(*character) {
				return false
			}
			changePokemon(character)
		}
		pressEnterToContinue()
	}
	if character.Pokemons[0].Speed < p.Speed {
		pokemonAttackYou(character, *p, randomAttack)
		if character.Pokemons[0].ActualHP <= 0 {
			if !checkPokemonDead(*character) {
				return false
			}
			changePokemon(character)
		}
		attackPokemon(*character, p, attackOpt)
		pressEnterToContinue()
		if p.ActualHP <= 0 {
			return false
		}
	}
	return true
}

func seeBag(p Pokemon, character Character) bool {
	return false
}

func run(p Pokemon, character Character) bool {
	if p.Level > character.Pokemons[0].Level {
		fmt.Printf("%s blocked your escape\n", p.Name)
		return true
	}
	fmt.Print("Ran away safely!")
	return false
}

func seePokemon(character Character) {}

func inCombatInBush(character *Character, route int) bool {
	cls()
	var currentRoute Route
	var p Pokemon

	randomLevel := 0
	randomIds := 0
	idsSize := 0
	option := 0
	switch route {
	case 1:
		currentRoute = RouteOne
	}

	idsSize = len(currentRoute.IDs)
	randomIds = rand.Intn(idsSize)
	randomLevel = rand.Intn(currentRoute.Levels[1]-currentRoute.Levels[0]+1) + currentRoute.Levels[0]
	p = createPokemon(currentRoute.IDs[randomIds], randomLevel)

	inCombat := true

	fmt.Printf("Wild %s level: %d appeared!\n\n", p.Name, p.Level)
	for inCombat {
		option = menuCombat(p)
		fmt.Printf("%s HP: %d/%d\n", p.Name, p.ActualHP, p.HP)
		fmt.Printf("%s HP: %d/%d", character.Pokemons[0].Name, character.Pokemons[0].ActualHP, character.Pokemons[0].HP)
		switch option {
		case 1:
			inCombat = fight(&p, character)
			if !checkPokemonDead(*character) {
				return true
			}
		case 2:
			inCombat = seeBag(p, *character)
		case 3:
			inCombat = run(p, *character)
		case 4:
			seePokemon(*character)
		}
	}

	return false
}
