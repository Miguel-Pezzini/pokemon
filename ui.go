package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	ansiReset  = "\033[0m"
	ansiBold   = "\033[1m"
	ansiDim    = "\033[2m"
	ansiItalic = "\033[3m"

	colorRed     = "\033[38;5;203m"
	colorYellow  = "\033[38;5;220m"
	colorGold    = "\033[38;5;214m"
	colorGreen   = "\033[38;5;40m"
	colorBlue    = "\033[38;5;75m"
	colorCyan    = "\033[38;5;81m"
	colorMagenta = "\033[38;5;177m"
	colorWhite   = "\033[38;5;231m"
	colorGray    = "\033[38;5;245m"
	colorDark    = "\033[38;5;238m"
	colorOrange  = "\033[38;5;208m"
	colorPink    = "\033[38;5;218m"
)

func visibleLen(s string) int {
	out := 0
	inEsc := false
	for i := 0; i < len(s); {
		if s[i] == '\033' {
			inEsc = true
			i++
			continue
		}
		if inEsc {
			if s[i] == 'm' {
				inEsc = false
			}
			i++
			continue
		}
		_, size := utf8.DecodeRuneInString(s[i:])
		out++
		i += size
	}
	return out
}

func padRight(s string, width int) string {
	diff := width - visibleLen(s)
	if diff <= 0 {
		return s
	}
	return s + strings.Repeat(" ", diff)
}

func centerText(s string, width int) string {
	diff := width - visibleLen(s)
	if diff <= 0 {
		return s
	}
	left := diff / 2
	right := diff - left
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}

func drawBox(lines []string, width int, borderColor string) {
	if width <= 0 {
		for _, l := range lines {
			if n := visibleLen(l); n > width {
				width = n
			}
		}
		width += 4
	}
	inner := width - 2
	top := borderColor + "╔" + strings.Repeat("═", inner) + "╗" + ansiReset
	bot := borderColor + "╚" + strings.Repeat("═", inner) + "╝" + ansiReset
	side := borderColor + "║" + ansiReset

	fmt.Println(top)
	for _, l := range lines {
		fmt.Println(side + " " + padRight(l, inner-2) + " " + side)
	}
	fmt.Println(bot)
}

func drawTitledBox(title string, lines []string, width int, borderColor, titleColor string) {
	if width <= 0 {
		for _, l := range lines {
			if n := visibleLen(l); n > width {
				width = n
			}
		}
		if n := visibleLen(title) + 4; n > width {
			width = n
		}
		width += 4
	}
	inner := width - 2

	titleText := " " + titleColor + ansiBold + title + ansiReset + borderColor + " "
	titleVis := visibleLen(titleText)
	leftPad := 2
	rightPad := inner - leftPad - titleVis
	if rightPad < 0 {
		rightPad = 0
	}
	top := borderColor + "╔" + strings.Repeat("═", leftPad) + ansiReset + titleText + borderColor + strings.Repeat("═", rightPad) + "╗" + ansiReset
	bot := borderColor + "╚" + strings.Repeat("═", inner) + "╝" + ansiReset
	side := borderColor + "║" + ansiReset

	fmt.Println(top)
	for _, l := range lines {
		fmt.Println(side + " " + padRight(l, inner-2) + " " + side)
	}
	fmt.Println(bot)
}

func wrapText(s string, width int) []string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return []string{""}
	}
	var lines []string
	cur := ""
	for _, w := range words {
		if cur == "" {
			cur = w
			continue
		}
		if visibleLen(cur)+1+visibleLen(w) > width {
			lines = append(lines, cur)
			cur = w
		} else {
			cur += " " + w
		}
	}
	if cur != "" {
		lines = append(lines, cur)
	}
	return lines
}

func hpBar(current, max, width int, label string) string {
	if max <= 0 {
		max = 1
	}
	if current < 0 {
		current = 0
	}
	ratio := float64(current) / float64(max)
	filled := int(ratio*float64(width) + 0.5)
	if filled > width {
		filled = width
	}

	color := colorGreen
	switch {
	case ratio <= 0.25:
		color = colorRed
	case ratio <= 0.5:
		color = colorYellow
	}

	bar := color + strings.Repeat("█", filled) + colorDark + strings.Repeat("░", width-filled) + ansiReset
	hpText := fmt.Sprintf("%s%3d%s/%s%d%s", colorWhite, current, colorGray, colorWhite, max, ansiReset)
	return fmt.Sprintf("%s%-12s%s %s%s HP %s %s", colorBold(), label, ansiReset, colorGray+"["+ansiReset, bar, colorGray+"]"+ansiReset, hpText)
}

func colorBold() string { return ansiBold + colorWhite }

func typeColor(t PokemonType) string {
	switch t {
	case TypeFire:
		return "\033[48;5;202m\033[38;5;231m"
	case TypeWater:
		return "\033[48;5;33m\033[38;5;231m"
	case TypeGrass:
		return "\033[48;5;34m\033[38;5;231m"
	case TypeElectric:
		return "\033[48;5;220m\033[38;5;232m"
	case TypePsychic:
		return "\033[48;5;177m\033[38;5;232m"
	case TypeIce:
		return "\033[48;5;123m\033[38;5;232m"
	case TypeDragon:
		return "\033[48;5;55m\033[38;5;231m"
	case TypeFairy:
		return "\033[48;5;218m\033[38;5;232m"
	case TypeBug:
		return "\033[48;5;107m\033[38;5;232m"
	case TypeRock:
		return "\033[48;5;138m\033[38;5;231m"
	case TypeGhost:
		return "\033[48;5;92m\033[38;5;231m"
	case TypeDark:
		return "\033[48;5;238m\033[38;5;231m"
	case TypeFighting:
		return "\033[48;5;124m\033[38;5;231m"
	case TypePoison:
		return "\033[48;5;127m\033[38;5;231m"
	case TypeGround:
		return "\033[48;5;178m\033[38;5;232m"
	case TypeFlying:
		return "\033[48;5;111m\033[38;5;232m"
	case TypeSteel:
		return "\033[48;5;145m\033[38;5;232m"
	default:
		return "\033[48;5;249m\033[38;5;232m"
	}
}

func typeName(t PokemonType) string {
	switch t {
	case TypeNormal:
		return "NORMAL"
	case TypeFire:
		return "FIRE"
	case TypeWater:
		return "WATER"
	case TypeGrass:
		return "GRASS"
	case TypeElectric:
		return "ELECTRIC"
	case TypePsychic:
		return "PSYCHIC"
	case TypeIce:
		return "ICE"
	case TypeDragon:
		return "DRAGON"
	case TypeFairy:
		return "FAIRY"
	case TypeBug:
		return "BUG"
	case TypeRock:
		return "ROCK"
	case TypeGhost:
		return "GHOST"
	case TypeDark:
		return "DARK"
	case TypeFighting:
		return "FIGHTING"
	case TypePoison:
		return "POISON"
	case TypeGround:
		return "GROUND"
	case TypeFlying:
		return "FLYING"
	case TypeSteel:
		return "STEEL"
	}
	return "?"
}

func typeBadge(t PokemonType) string {
	return typeColor(t) + " " + typeName(t) + " " + ansiReset
}
