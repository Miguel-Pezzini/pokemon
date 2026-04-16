package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	battleAnimationStartDelay = 260 * time.Millisecond
	battleAnimationFrameDelay = 140 * time.Millisecond
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

func renderBattleScreen(player, enemy Pokemon, messages []string) {
	cls()
	fmt.Println()

	enemyLines := []string{
		fmt.Sprintf("%s%-14s%s %s  %sLv.%s%d%s",
			ansiBold+colorWhite, enemy.Name, ansiReset,
			typeBadge(enemy.Type),
			colorGray, colorYellow, enemy.Level, ansiReset),
		hpBar(enemy.ActualHP, enemy.HP, 22, "ENEMY"),
	}
	drawTitledBox("WILD POKEMON", enemyLines, 60, colorRed, colorOrange)

	fmt.Println()

	playerLines := []string{
		fmt.Sprintf("%s%-14s%s %s  %sLv.%s%d%s",
			ansiBold+colorWhite, player.Name, ansiReset,
			typeBadge(player.Type),
			colorGray, colorYellow, player.Level, ansiReset),
		hpBar(player.ActualHP, player.HP, 22, "YOURS"),
	}
	drawTitledBox("YOUR POKEMON", playerLines, 60, colorGreen, colorCyan)

	if len(messages) > 0 {
		fmt.Println()
		drawTitledBox("BATTLE LOG", messages, 60, colorBlue, colorCyan)
	}
}

func hpAnimationStep(diff int) int {
	switch {
	case diff <= 18:
		return 1
	case diff <= 40:
		return 2
	case diff <= 75:
		return 3
	default:
		return diff/30 + 1
	}
}

func animateBattleHPChange(beforePlayer, afterPlayer, beforeEnemy, afterEnemy Pokemon, messages []string) {
	playerFrame := beforePlayer
	enemyFrame := beforeEnemy

	playerDiff := beforePlayer.ActualHP - afterPlayer.ActualHP
	enemyDiff := beforeEnemy.ActualHP - afterEnemy.ActualHP

	if playerDiff <= 0 && enemyDiff <= 0 {
		renderBattleScreen(afterPlayer, afterEnemy, messages)
		return
	}

	playerStep := hpAnimationStep(playerDiff)
	enemyStep := hpAnimationStep(enemyDiff)

	renderBattleScreen(playerFrame, enemyFrame, messages)
	time.Sleep(battleAnimationStartDelay)

	for playerFrame.ActualHP > afterPlayer.ActualHP || enemyFrame.ActualHP > afterEnemy.ActualHP {
		if playerFrame.ActualHP > afterPlayer.ActualHP {
			playerFrame.ActualHP -= playerStep
			if playerFrame.ActualHP < afterPlayer.ActualHP {
				playerFrame.ActualHP = afterPlayer.ActualHP
			}
		}

		if enemyFrame.ActualHP > afterEnemy.ActualHP {
			enemyFrame.ActualHP -= enemyStep
			if enemyFrame.ActualHP < afterEnemy.ActualHP {
				enemyFrame.ActualHP = afterEnemy.ActualHP
			}
		}

		renderBattleScreen(playerFrame, enemyFrame, messages)
		time.Sleep(battleAnimationFrameDelay)
	}
}

func battleMessages(messages []string, extra ...string) []string {
	out := append([]string{}, messages...)
	out = append(out, extra...)
	return out
}

func battleXPGain(enemy Pokemon) int {
	return enemy.Level * 25
}

func showEnemyDefeatedScreen(character *Character, enemy Pokemon, messages []string) {
	xpGained := battleXPGain(enemy)
	currentPokemon := &character.Pokemons[0]
	previousXP := currentPokemon.XP
	currentPokemon.XP += xpGained

	resultMessages := battleMessages(messages,
		colorGreen+ansiBold+enemy.Name+" fainted!"+ansiReset,
		fmt.Sprintf("%sXP gained:%s %s%d XP%s",
			colorGray, ansiReset, colorYellow, xpGained, ansiReset),
		fmt.Sprintf("%s%s%s %sXP:%s %s%d -> %d/%d%s",
			ansiBold+colorWhite, currentPokemon.Name, ansiReset,
			colorGray, ansiReset,
			colorYellow, previousXP, currentPokemon.XP, currentPokemon.XPToUp, ansiReset),
	)

	renderBattleScreen(*currentPokemon, enemy, resultMessages)
	pressEnterToReturnToMap()
}

func showPlayerFaintedScreen(character *Character, enemy Pokemon, messages []string) bool {
	resultMessages := battleMessages(messages, colorRed+ansiBold+character.Pokemons[0].Name+" fainted!"+ansiReset)

	if !checkPokemonDead(*character) {
		resultMessages = append(resultMessages,
			colorRed+ansiBold+"You were defeated!"+ansiReset,
			colorGray+"Your team needs to recover before battling again."+ansiReset,
		)
		renderBattleScreen(character.Pokemons[0], enemy, resultMessages)
		pressEnterToReturnToMap()
		return false
	}

	renderBattleScreen(character.Pokemons[0], enemy, resultMessages)
	pressEnterToContinue()
	changePokemon(character)
	return true
}

func menuCombat(p Pokemon) int {
	opt := 0
	for opt < 1 || opt > 4 {
		fmt.Println()
		lines := []string{
			colorRed + "  1 " + colorGray + "| " + colorWhite + "FIGHT  " + colorDark + "- attack the enemy",
			colorYellow + "  2 " + colorGray + "| " + colorWhite + "BAG    " + colorDark + "- use an item",
			colorBlue + "  3 " + colorGray + "| " + colorWhite + "RUN    " + colorDark + "- flee the battle",
			colorGreen + "  4 " + colorGray + "| " + colorWhite + "POKEMON" + colorDark + " - switch out",
		}
		drawTitledBox("ACTIONS", lines, 60, colorGold, colorYellow)
		fmt.Print(colorCyan + "  > " + ansiReset + "What will you do? " + colorGray + "(1-4): " + ansiReset)
		scanIntInput(&opt)
		if opt < 1 || opt > 4 {
			fmt.Println(colorRed + "  ! Invalid option. Try again." + ansiReset)
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

func formatDamage(attacker string, dmg int, modifier float64) string {
	switch {
	case modifier == 1.5:
		return fmt.Sprintf("%s%s%s %sused an %seffective%s attack - %s-%d HP%s",
			ansiBold+colorGreen, attacker, ansiReset,
			colorGray, colorYellow, colorGray,
			colorRed, dmg, ansiReset)
	case modifier == 0.5:
		return fmt.Sprintf("%s%s%s %sattack %swasn't very effective%s - %s-%d HP%s",
			ansiBold+colorGray, attacker, ansiReset,
			colorGray, colorGray, colorGray,
			colorOrange, dmg, ansiReset)
	default:
		return fmt.Sprintf("%s%s%s %sattacked for %s-%d HP%s",
			ansiBold+colorWhite, attacker, ansiReset,
			colorGray, colorRed, dmg, ansiReset)
	}
}

func attackPokemon(character Character, p *Pokemon, attackOpt int) string {
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
	return formatDamage(character.Pokemons[0].Name, dano, typeModifier)
}

func pokemonAttackYou(character *Character, p Pokemon, randomAttack int) string {
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
	return formatDamage(p.Name, dano, typeModifier)
}

func changePokemon(character *Character) {
	size := len(character.Pokemons)
	indexPoke := 0
	fmt.Println()
	var lines []string
	lines = append(lines, colorRed+"  Your Pokemon fainted!"+ansiReset)
	lines = append(lines, "")
	for i := 0; i < size; i++ {
		if character.Pokemons[i].ActualHP > 0 {
			lines = append(lines, fmt.Sprintf("%s  %d %s| %s%-12s%s  %s  %sLv.%s%d%s",
				colorYellow, i, colorGray,
				colorWhite+ansiBold, character.Pokemons[i].Name, ansiReset,
				typeBadge(character.Pokemons[i].Type),
				colorGray, colorYellow, character.Pokemons[i].Level, ansiReset))
		}
	}
	drawTitledBox("CHOOSE YOUR NEXT POKEMON", lines, 60, colorRed, colorOrange)
	fmt.Print(colorCyan + "  > " + ansiReset + "Choose: " + ansiReset)
	scanIntInput(&indexPoke)
	temp := character.Pokemons[0]
	character.Pokemons[0] = character.Pokemons[indexPoke]
	character.Pokemons[indexPoke] = temp
}

func scanIntInput(value *int) {
	fmt.Scan(value)
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}

func pressEnterToContinue() {
	fmt.Println()
	fmt.Print(colorGray + "  Press Enter to continue..." + ansiReset)
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}

func pressEnterToReturnToMap() {
	fmt.Println()
	fmt.Print(colorGray + "  Press Enter to return to the map..." + ansiReset)
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}

func chooseAttack(character *Character) int {
	attackOpt := 0
	for {
		var lines []string
		for i, atk := range character.Pokemons[0].Attacks {
			pp := fmt.Sprintf("%sPWR %s%-3d%s", colorGray, colorYellow, atk.Power, ansiReset)
			lines = append(lines, fmt.Sprintf("%s  %d %s| %s%-14s%s  %s  %s",
				colorGold, i+1, colorGray,
				colorWhite+ansiBold, atk.Name, ansiReset,
				typeBadge(atk.Type),
				pp))
		}
		drawTitledBox(strings.ToUpper(character.Pokemons[0].Name)+"'S ATTACKS", lines, 60, colorMagenta, colorPink)
		fmt.Print(colorCyan + "  > " + ansiReset + "Pick an attack: " + ansiReset)
		scanIntInput(&attackOpt)
		if attackOpt >= 1 && attackOpt <= len(character.Pokemons[0].Attacks) {
			return attackOpt - 1
		}
		fmt.Println(colorRed + "  ! Invalid attack." + ansiReset)
	}
}

func fight(p *Pokemon, character *Character) bool {
	attacksWildLenght := len(p.Attacks)
	randomAttack := rand.Intn(attacksWildLenght)

	attackOpt := chooseAttack(character)

	var msgs []string
	if character.Pokemons[0].Speed > p.Speed {
		playerBefore := character.Pokemons[0]
		enemyBefore := *p
		msgs = append(msgs, attackPokemon(*character, p, attackOpt))
		animateBattleHPChange(playerBefore, character.Pokemons[0], enemyBefore, *p, msgs)
		if p.ActualHP <= 0 {
			showEnemyDefeatedScreen(character, *p, msgs)
			return false
		}

		playerBefore = character.Pokemons[0]
		enemyBefore = *p
		msgs = append(msgs, pokemonAttackYou(character, *p, randomAttack))
		animateBattleHPChange(playerBefore, character.Pokemons[0], enemyBefore, *p, msgs)
		if character.Pokemons[0].ActualHP <= 0 {
			return showPlayerFaintedScreen(character, *p, msgs)
		}

		pressEnterToContinue()
		return true
	}

	playerBefore := character.Pokemons[0]
	enemyBefore := *p
	msgs = append(msgs, pokemonAttackYou(character, *p, randomAttack))
	animateBattleHPChange(playerBefore, character.Pokemons[0], enemyBefore, *p, msgs)
	if character.Pokemons[0].ActualHP <= 0 {
		return showPlayerFaintedScreen(character, *p, msgs)
	}

	playerBefore = character.Pokemons[0]
	enemyBefore = *p
	msgs = append(msgs, attackPokemon(*character, p, attackOpt))
	animateBattleHPChange(playerBefore, character.Pokemons[0], enemyBefore, *p, msgs)
	if p.ActualHP <= 0 {
		showEnemyDefeatedScreen(character, *p, msgs)
		return false
	}

	pressEnterToContinue()
	return true
}

func seeBag(p Pokemon, character Character) bool {
	return false
}

func run(p Pokemon, character Character) bool {
	fmt.Println()
	if p.Level > character.Pokemons[0].Level {
		drawBox([]string{
			colorRed + ansiBold + "  ! " + p.Name + " blocked your escape!" + ansiReset,
		}, 50, colorRed)
		pressEnterToContinue()
		return true
	}
	drawBox([]string{
		colorGreen + ansiBold + "  Safe escape!" + ansiReset,
	}, 50, colorGreen)
	pressEnterToContinue()
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

	cls()
	fmt.Println()
	drawTitledBox("WILD ENCOUNTER", []string{
		"",
		centerText(fmt.Sprintf("%sA wild %s%s%s appeared!%s",
			colorWhite, ansiBold+colorYellow, strings.ToUpper(p.Name), ansiReset+colorWhite, ansiReset), 52),
		centerText(fmt.Sprintf("%sLevel %s%d%s  %s",
			colorGray, colorYellow, p.Level, ansiReset, typeBadge(p.Type)), 52),
		"",
	}, 60, colorRed, colorOrange)
	pressEnterToContinue()

	for inCombat {
		renderBattleScreen(character.Pokemons[0], p, nil)
		option = menuCombat(p)
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
