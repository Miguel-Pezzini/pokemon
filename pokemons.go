package main

type Pokemon struct {
	ID            int
	Name          string
	HP            int
	ActualHP      int
	Def           int
	Speed         int
	Level         int
	XPToUp        int
	XP            int
	LevelToEvolve int
	IDEvolution   int
	Attacks       []Attack
	Type          PokemonType
}

func pokemonName(id int) string {
	switch id {
	case 1:
		return "Bulbasaur"
	case 2:
		return "Ivysaur"
	case 3:
		return "Venusaur"
	case 4:
		return "Charmander"
	case 5:
		return "Charmeleon"
	case 6:
		return "Charizard"
	case 7:
		return "Squirtle"
	case 8:
		return "Wartortle"
	case 9:
		return "Blastoise"
	case 10:
		return "Caterpie"
	case 11:
		return "Metapod"
	case 12:
		return "Butterfree"
	case 13:
		return "Weedle"
	case 14:
		return "Kakuna"
	case 15:
		return "Beedrill"
	case 16:
		return "Pidgey"
	case 17:
		return "Pidgeotto"
	case 18:
		return "Pidgeot"
	case 19:
		return "Rattata"
	case 20:
		return "Raticate"
	case 21:
		return "Spearow"
	case 22:
		return "Fearow"
	case 23:
		return "Ekans"
	case 24:
		return "Arbok"
	case 25:
		return "Pikachu"
	case 26:
		return "Raichu"
	case 27:
		return "Sandshrew"
	case 28:
		return "Sandslash"
	case 29:
		return "Nidoran♀"
	case 30:
		return "Nidorina"
	case 31:
		return "Nidoqueen"
	case 32:
		return "Nidoran♂"
	case 33:
		return "Nidorino"
	case 34:
		return "Nidoking"
	case 35:
		return "Clefairy"
	case 36:
		return "Clefable"
	case 37:
		return "Vulpix"
	case 38:
		return "Ninetales"
	case 39:
		return "Jigglypuff"
	case 40:
		return "Wigglytuff"
	case 41:
		return "Zubat"
	case 42:
		return "Golbat"
	case 43:
		return "Oddish"
	case 44:
		return "Gloom"
	case 45:
		return "Vileplume"
	case 46:
		return "Paras"
	case 47:
		return "Parasect"
	case 48:
		return "Venonat"
	case 49:
		return "Venomoth"
	case 50:
		return "Diglett"
	case 51:
		return "Dugtrio"
	case 52:
		return "Meowth"
	case 53:
		return "Persian"
	case 54:
		return "Psyduck"
	case 55:
		return "Golduck"
	case 56:
		return "Poliwag"
	case 57:
		return "Poliwhirl"
	case 58:
		return "Poliwrath"
	case 59:
		return "Abra"
	case 60:
		return "Kadabra"
	case 61:
		return "Alakazam"
	case 62:
		return "Machop"
	case 63:
		return "Machoke"
	case 64:
		return "Machamp"
	case 65:
		return "Bellsprout"
	case 66:
		return "Weepinbell"
	case 67:
		return "Victreebel"
	case 68:
		return "Tentacool"
	case 69:
		return "Tentacruel"
	case 70:
		return "Geodude"
	case 71:
		return "Graveler"
	case 72:
		return "Golem"
	case 73:
		return "Ponyta"
	case 74:
		return "Rapidash"
	case 75:
		return "Slowpoke"
	case 76:
		return "Slowbro"
	case 77:
		return "Magnemite"
	case 78:
		return "Magneton"
	case 79:
		return "Farfetch'd"
	case 80:
		return "Doduo"
	case 81:
		return "Dodrio"
	case 82:
		return "Seel"
	case 83:
		return "Dewgong"
	case 84:
		return "Grimer"
	case 85:
		return "Muk"
	case 86:
		return "Shellder"
	case 87:
		return "Cloyster"
	case 88:
		return "Gastly"
	case 89:
		return "Haunter"
	case 90:
		return "Gengar"
	case 91:
		return "Onix"
	case 92:
		return "Drowzee"
	case 93:
		return "Hypno"
	case 94:
		return "Krabby"
	case 95:
		return "Kingler"
	case 96:
		return "Voltorb"
	case 97:
		return "Electrode"
	case 98:
		return "Exeggcute"
	case 99:
		return "Exeggutor"
	case 100:
		return "Cubone"
	case 101:
		return "Marowak"
	case 102:
		return "Hitmonlee"
	case 103:
		return "Hitmonchan"
	case 104:
		return "Lickitung"
	case 105:
		return "Koffing"
	case 106:
		return "Weezing"
	case 107:
		return "Rhyhorn"
	case 108:
		return "Rhydon"
	case 109:
		return "Chansey"
	case 110:
		return "Tangela"
	case 111:
		return "Kangaskhan"
	case 112:
		return "Horsea"
	case 113:
		return "Seadra"
	case 114:
		return "Goldeen"
	case 115:
		return "Seaking"
	case 116:
		return "Staryu"
	case 117:
		return "Starmie"
	case 118:
		return "Mr. Mime"
	case 119:
		return "Scyther"
	case 120:
		return "Jynx"
	case 121:
		return "Electabuzz"
	case 122:
		return "Magmar"
	case 123:
		return "Pinsir"
	case 124:
		return "Tauros"
	case 125:
		return "Magikarp"
	case 126:
		return "Gyarados"
	case 127:
		return "Lapras"
	case 128:
		return "Ditto"
	case 129:
		return "Eevee"
	case 130:
		return "Vaporeon"
	case 131:
		return "Jolteon"
	case 132:
		return "Flareon"
	case 133:
		return "Porygon"
	case 134:
		return "Omanyte"
	case 135:
		return "Omastar"
	case 136:
		return "Kabuto"
	case 137:
		return "Kabutops"
	case 138:
		return "Aerodactyl"
	case 139:
		return "Snorlax"
	case 140:
		return "Articuno"
	case 141:
		return "Zapdos"
	case 142:
		return "Moltres"
	case 143:
		return "Dratini"
	case 144:
		return "Dragonair"
	case 145:
		return "Dragonite"
	case 146:
		return "Mewtwo"
	case 147:
		return "Mew"
	default:
		return "Unknown Pokémon"
	}
}

func pokemonHP(id, level int) int {
	var baseHP int
	switch id {
	case 1:
		baseHP = 45
	case 2:
		baseHP = 60
	case 3:
		baseHP = 80
	case 4:
		baseHP = 39
	case 5:
		baseHP = 58
	case 6:
		baseHP = 78
	case 7:
		baseHP = 44
	case 8:
		baseHP = 59
	case 9:
		baseHP = 79
	case 10:
		baseHP = 45
	case 11:
		baseHP = 50
	case 12:
		baseHP = 60
	case 13:
		baseHP = 40
	case 14:
		baseHP = 45
	case 15:
		baseHP = 65
	case 16:
		baseHP = 40
	case 17:
		baseHP = 63
	case 18:
		baseHP = 83
	case 19:
		baseHP = 30
	case 20:
		baseHP = 55
	case 21:
		baseHP = 40
	case 22:
		baseHP = 65
	case 23:
		baseHP = 35
	case 24:
		baseHP = 60
	case 25:
		baseHP = 35
	case 26:
		baseHP = 60
	case 27:
		baseHP = 50
	case 28:
		baseHP = 75
	case 29:
		baseHP = 55
	case 30:
		baseHP = 70
	case 31:
		baseHP = 90
	case 32:
		baseHP = 46
	case 33:
		baseHP = 61
	case 34:
		baseHP = 81
	case 35:
		baseHP = 70
	case 36:
		baseHP = 95
	case 37:
		baseHP = 38
	case 38:
		baseHP = 73
	case 39:
		baseHP = 115
	case 40:
		baseHP = 140
	case 41:
		baseHP = 40
	case 42:
		baseHP = 75
	case 43:
		baseHP = 45
	case 44:
		baseHP = 60
	case 45:
		baseHP = 75
	case 46:
		baseHP = 35
	case 47:
		baseHP = 60
	case 48:
		baseHP = 60
	case 49:
		baseHP = 70
	case 50:
		baseHP = 10
	case 51:
		baseHP = 35
	case 52:
		baseHP = 40
	case 53:
		baseHP = 65
	case 54:
		baseHP = 50
	case 55:
		baseHP = 80
	case 56:
		baseHP = 50
	case 57:
		baseHP = 65
	case 58:
		baseHP = 90
	case 59:
		baseHP = 25
	case 60:
		baseHP = 40
	case 61:
		baseHP = 55
	case 62:
		baseHP = 70
	case 63:
		baseHP = 80
	case 64:
		baseHP = 90
	case 65:
		baseHP = 50
	case 66:
		baseHP = 65
	case 67:
		baseHP = 80
	case 68:
		baseHP = 40
	case 69:
		baseHP = 80
	case 70:
		baseHP = 40
	case 71:
		baseHP = 55
	case 72:
		baseHP = 80
	case 73:
		baseHP = 50
	case 74:
		baseHP = 65
	case 75:
		baseHP = 90
	case 76:
		baseHP = 95
	case 77:
		baseHP = 25
	case 78:
		baseHP = 50
	case 79:
		baseHP = 52
	case 80:
		baseHP = 35
	case 81:
		baseHP = 60
	case 82:
		baseHP = 65
	case 83:
		baseHP = 90
	case 84:
		baseHP = 80
	case 85:
		baseHP = 105
	case 86:
		baseHP = 30
	case 87:
		baseHP = 50
	case 88:
		baseHP = 30
	case 89:
		baseHP = 45
	case 90:
		baseHP = 60
	case 91:
		baseHP = 35
	case 92:
		baseHP = 60
	case 93:
		baseHP = 80
	case 94:
		baseHP = 30
	case 95:
		baseHP = 55
	case 96:
		baseHP = 40
	case 97:
		baseHP = 60
	case 98:
		baseHP = 60
	case 99:
		baseHP = 95
	case 100:
		baseHP = 50
	case 101:
		baseHP = 60
	case 102:
		baseHP = 50
	case 103:
		baseHP = 50
	case 104:
		baseHP = 90
	case 105:
		baseHP = 40
	case 106:
		baseHP = 65
	case 107:
		baseHP = 80
	case 108:
		baseHP = 105
	case 109:
		baseHP = 250
	case 110:
		baseHP = 65
	case 111:
		baseHP = 105
	case 112:
		baseHP = 30
	case 113:
		baseHP = 55
	case 114:
		baseHP = 45
	case 115:
		baseHP = 80
	case 116:
		baseHP = 30
	case 117:
		baseHP = 60
	case 118:
		baseHP = 40
	case 119:
		baseHP = 70
	case 120:
		baseHP = 65
	case 121:
		baseHP = 65
	case 122:
		baseHP = 65
	case 123:
		baseHP = 65
	case 124:
		baseHP = 75
	case 125:
		baseHP = 20
	case 126:
		baseHP = 95
	case 127:
		baseHP = 130
	case 128:
		baseHP = 35
	case 129:
		baseHP = 55
	case 130:
		baseHP = 130
	case 131:
		baseHP = 65
	case 132:
		baseHP = 65
	case 133:
		baseHP = 65
	case 134:
		baseHP = 35
	case 135:
		baseHP = 70
	case 136:
		baseHP = 30
	case 137:
		baseHP = 60
	case 138:
		baseHP = 80
	case 139:
		baseHP = 110
	case 140:
		baseHP = 90
	case 141:
		baseHP = 90
	case 142:
		baseHP = 90
	case 143:
		baseHP = 41
	case 144:
		baseHP = 61
	case 145:
		baseHP = 91
	case 146:
		baseHP = 106
	case 147:
		baseHP = 100
	default:
		baseHP = 50
	}
	return (baseHP*2*level)/100 + level + 10
}

func pokemonDEF(id, level int) int {
	var baseDefense int
	switch id {
	case 1:
		baseDefense = 49
	case 2:
		baseDefense = 63
	case 3:
		baseDefense = 83
	case 4:
		baseDefense = 43
	case 5:
		baseDefense = 58
	case 6:
		baseDefense = 65
	case 7:
		baseDefense = 48
	case 8:
		baseDefense = 65
	case 9:
		baseDefense = 80
	case 10:
		baseDefense = 30
	case 11:
		baseDefense = 55
	case 12:
		baseDefense = 50
	case 13:
		baseDefense = 35
	case 14:
		baseDefense = 35
	case 15:
		baseDefense = 80
	case 16:
		baseDefense = 40
	case 17:
		baseDefense = 55
	case 18:
		baseDefense = 75
	case 19:
		baseDefense = 30
	case 20:
		baseDefense = 50
	case 21:
		baseDefense = 30
	case 22:
		baseDefense = 65
	case 23:
		baseDefense = 35
	case 24:
		baseDefense = 49
	case 25:
		baseDefense = 40
	case 26:
		baseDefense = 50
	case 27:
		baseDefense = 75
	case 28:
		baseDefense = 100
	case 29:
		baseDefense = 40
	case 30:
		baseDefense = 55
	case 31:
		baseDefense = 70
	case 32:
		baseDefense = 40
	case 33:
		baseDefense = 55
	case 34:
		baseDefense = 80
	case 35:
		baseDefense = 48
	case 36:
		baseDefense = 70
	case 37:
		baseDefense = 40
	case 38:
		baseDefense = 75
	case 39:
		baseDefense = 30
	case 40:
		baseDefense = 70
	case 41:
		baseDefense = 45
	case 42:
		baseDefense = 70
	case 43:
		baseDefense = 55
	case 44:
		baseDefense = 50
	case 45:
		baseDefense = 70
	case 46:
		baseDefense = 55
	case 47:
		baseDefense = 60
	case 48:
		baseDefense = 50
	case 49:
		baseDefense = 60
	case 50:
		baseDefense = 25
	case 51:
		baseDefense = 50
	case 52:
		baseDefense = 45
	case 53:
		baseDefense = 60
	case 54:
		baseDefense = 50
	case 55:
		baseDefense = 80
	case 56:
		baseDefense = 50
	case 57:
		baseDefense = 65
	case 58:
		baseDefense = 90
	case 59:
		baseDefense = 20
	case 60:
		baseDefense = 25
	case 61:
		baseDefense = 40
	case 62:
		baseDefense = 70
	case 63:
		baseDefense = 80
	case 64:
		baseDefense = 90
	case 65:
		baseDefense = 35
	case 66:
		baseDefense = 50
	case 67:
		baseDefense = 65
	case 68:
		baseDefense = 40
	case 69:
		baseDefense = 70
	case 70:
		baseDefense = 30
	case 71:
		baseDefense = 100
	case 72:
		baseDefense = 120
	case 73:
		baseDefense = 55
	case 74:
		baseDefense = 80
	case 75:
		baseDefense = 65
	case 76:
		baseDefense = 80
	case 77:
		baseDefense = 60
	case 78:
		baseDefense = 95
	case 79:
		baseDefense = 50
	case 80:
		baseDefense = 35
	case 81:
		baseDefense = 60
	case 82:
		baseDefense = 85
	case 83:
		baseDefense = 70
	case 84:
		baseDefense = 50
	case 85:
		baseDefense = 75
	case 86:
		baseDefense = 55
	case 87:
		baseDefense = 95
	case 88:
		baseDefense = 30
	case 89:
		baseDefense = 45
	case 90:
		baseDefense = 60
	case 91:
		baseDefense = 160
	case 92:
		baseDefense = 55
	case 93:
		baseDefense = 70
	case 94:
		baseDefense = 50
	case 95:
		baseDefense = 75
	case 96:
		baseDefense = 40
	case 97:
		baseDefense = 60
	case 98:
		baseDefense = 55
	case 99:
		baseDefense = 70
	case 100:
		baseDefense = 50
	case 101:
		baseDefense = 80
	case 102:
		baseDefense = 90
	case 103:
		baseDefense = 110
	case 104:
		baseDefense = 70
	case 105:
		baseDefense = 70
	case 106:
		baseDefense = 120
	case 107:
		baseDefense = 95
	case 108:
		baseDefense = 130
	case 109:
		baseDefense = 30
	case 110:
		baseDefense = 60
	case 111:
		baseDefense = 80
	case 112:
		baseDefense = 50
	case 113:
		baseDefense = 70
	case 114:
		baseDefense = 45
	case 115:
		baseDefense = 70
	case 116:
		baseDefense = 30
	case 117:
		baseDefense = 60
	case 118:
		baseDefense = 70
	case 119:
		baseDefense = 55
	case 120:
		baseDefense = 35
	case 121:
		baseDefense = 57
	case 122:
		baseDefense = 57
	case 123:
		baseDefense = 75
	case 124:
		baseDefense = 75
	case 125:
		baseDefense = 20
	case 126:
		baseDefense = 79
	case 127:
		baseDefense = 85
	case 128:
		baseDefense = 70
	case 129:
		baseDefense = 50
	case 130:
		baseDefense = 60
	case 131:
		baseDefense = 65
	case 132:
		baseDefense = 65
	case 133:
		baseDefense = 65
	case 134:
		baseDefense = 55
	case 135:
		baseDefense = 70
	case 136:
		baseDefense = 55
	case 137:
		baseDefense = 65
	case 138:
		baseDefense = 60
	case 139:
		baseDefense = 110
	case 140:
		baseDefense = 100
	case 141:
		baseDefense = 90
	case 142:
		baseDefense = 90
	case 143:
		baseDefense = 60
	case 144:
		baseDefense = 65
	case 145:
		baseDefense = 95
	case 146:
		baseDefense = 90
	case 147:
		baseDefense = 100
	default:
		baseDefense = 50
	}
	return (baseDefense*2*level)/100 + 5
}

func pokemonXPTOUP(level int) int {
	return level * 100
}

func pokemonLevelToEvolve(id int) int {
	switch id {
	case 1:
		return 16
	case 2:
		return 32
	case 4:
		return 16
	case 5:
		return 36
	case 7:
		return 16
	case 8:
		return 36
	case 10:
		return 7
	case 11:
		return 10
	case 13:
		return 7
	case 14:
		return 10
	case 16:
		return 18
	case 17:
		return 36
	case 19:
		return 20
	case 21:
		return 20
	case 23:
		return 22
	case 25:
		return 16
	case 27:
		return 22
	case 29:
		return 16
	case 30:
		return 31
	case 32:
		return 16
	case 33:
		return 31
	case 35:
		return 35
	case 37:
		return 20
	case 39:
		return 18
	case 41:
		return 22
	case 43:
		return 21
	case 45:
		return 21
	case 46:
		return 24
	case 48:
		return 31
	case 50:
		return 26
	case 52:
		return 28
	case 54:
		return 33
	case 56:
		return 25
	case 57:
		return 36
	case 59:
		return 16
	case 62:
		return 28
	case 63:
		return 35
	case 65:
		return 21
	case 66:
		return 35
	case 68:
		return 30
	case 70:
		return 25
	case 73:
		return 40
	case 75:
		return 30
	case 77:
		return 30
	case 79:
		return 35
	case 82:
		return 34
	case 84:
		return 40
	case 86:
		return 30
	case 88:
		return 25
	case 91:
		return 30
	case 94:
		return 30
	case 96:
		return 30
	case 100:
		return 28
	case 102:
		return 20
	case 103:
		return 20
	case 107:
		return 42
	case 112:
		return 32
	case 114:
		return 33
	case 126:
		return 20
	case 143:
		return 30
	default:
		return 0
	}
}

func pokemonType(id int) PokemonType {
	switch id {
	case 1, 2, 3:
		return TypeGrass
	case 4, 5, 6:
		return TypeFire
	case 7, 8, 9:
		return TypeWater
	case 10, 11, 12, 13, 14, 15:
		return TypeBug
	case 16, 17, 18, 19, 20, 21, 22:
		return TypeNormal
	case 23, 24:
		return TypePoison
	case 25, 26:
		return TypeElectric
	case 27, 28:
		return TypeGround
	case 29, 30:
		return TypePoison
	case 31:
		return TypeGround
	case 32, 33:
		return TypePoison
	case 34:
		return TypeGround
	case 35, 36:
		return TypeFairy
	case 37, 38:
		return TypeFire
	case 39:
		return TypeNormal
	case 40:
		return TypeFairy
	case 41, 42:
		return TypePoison
	case 43, 44, 45:
		return TypeGrass
	case 46, 47, 48, 49:
		return TypeBug
	case 50, 51:
		return TypeGround
	case 52, 53:
		return TypeNormal
	case 54, 55, 56, 57, 58:
		return TypeWater
	case 59, 60, 61:
		return TypePsychic
	case 62, 63, 64:
		return TypeFighting
	case 65, 66, 67:
		return TypeGrass
	case 68, 69:
		return TypeWater
	case 70, 71, 72:
		return TypeRock
	case 73, 74:
		return TypeFire
	case 75, 76:
		return TypeWater
	case 77, 78:
		return TypeElectric
	case 79, 80, 81:
		return TypeNormal
	case 82, 83:
		return TypeWater
	case 84, 85:
		return TypePoison
	case 86, 87:
		return TypeWater
	case 88, 89, 90:
		return TypeGhost
	case 91:
		return TypeRock
	case 92, 93:
		return TypePsychic
	case 94, 95:
		return TypeWater
	case 96, 97:
		return TypeElectric
	case 98, 99:
		return TypeGrass
	case 100, 101:
		return TypeGround
	case 102, 103:
		return TypeFighting
	case 104:
		return TypeNormal
	case 105, 106:
		return TypePoison
	case 107, 108:
		return TypeGround
	case 109:
		return TypeNormal
	case 110:
		return TypeGrass
	case 111:
		return TypeNormal
	case 112, 113, 114, 115, 116, 117:
		return TypeWater
	case 118:
		return TypePsychic
	case 119:
		return TypeBug
	case 120:
		return TypeIce
	case 121:
		return TypeElectric
	case 122:
		return TypeFire
	case 123:
		return TypeBug
	case 124:
		return TypeNormal
	case 125, 126, 127:
		return TypeWater
	case 128, 129:
		return TypeNormal
	case 130:
		return TypeWater
	case 131:
		return TypeElectric
	case 132:
		return TypeFire
	case 133:
		return TypeNormal
	case 134, 135, 136, 137, 138:
		return TypeRock
	case 139:
		return TypeNormal
	case 140:
		return TypeIce
	case 141:
		return TypeElectric
	case 142:
		return TypeFire
	case 143, 144, 145:
		return TypeDragon
	case 146, 147:
		return TypePsychic
	default:
		return TypeNormal
	}
}

func pokemonEvolutionId(id int) int {
	switch id {
	case 1:
		return 2
	case 2:
		return 3
	case 4:
		return 5
	case 5:
		return 6
	case 7:
		return 8
	case 8:
		return 9
	case 10:
		return 11
	case 11:
		return 12
	case 13:
		return 14
	case 14:
		return 15
	case 16:
		return 17
	case 17:
		return 18
	case 19:
		return 20
	case 21:
		return 22
	case 23:
		return 24
	case 25:
		return 26
	case 27:
		return 28
	case 29:
		return 30
	case 30:
		return 31
	case 32:
		return 33
	case 33:
		return 34
	case 35:
		return 36
	case 37:
		return 38
	case 39:
		return 40
	case 41:
		return 42
	case 43:
		return 44
	case 44:
		return 45
	case 46:
		return 47
	case 48:
		return 49
	case 50:
		return 51
	case 52:
		return 53
	case 54:
		return 55
	case 56:
		return 57
	case 57:
		return 58
	case 59:
		return 60
	case 60:
		return 61
	case 62:
		return 63
	case 63:
		return 64
	case 65:
		return 66
	case 66:
		return 67
	case 68:
		return 69
	case 70:
		return 71
	case 71:
		return 72
	case 73:
		return 74
	case 75:
		return 76
	case 77:
		return 78
	case 80:
		return 81
	case 82:
		return 83
	case 84:
		return 85
	case 86:
		return 87
	case 88:
		return 89
	case 89:
		return 90
	case 92:
		return 93
	case 94:
		return 95
	case 96:
		return 97
	case 98:
		return 99
	case 100:
		return 101
	case 105:
		return 106
	case 107:
		return 108
	case 112:
		return 113
	case 114:
		return 115
	case 116:
		return 117
	case 125:
		return 126
	case 129:
		return 130
	case 134:
		return 135
	case 136:
		return 137
	case 143:
		return 144
	case 144:
		return 145
	default:
		return 0
	}
}

func pokemonSPEED(id, level int) int {
	var baseSpeed int
	switch id {
	case 1:
		baseSpeed = 45
	case 2:
		baseSpeed = 60
	case 3:
		baseSpeed = 80
	case 4:
		baseSpeed = 65
	case 5:
		baseSpeed = 80
	case 6:
		baseSpeed = 100
	case 7:
		baseSpeed = 43
	case 8:
		baseSpeed = 58
	case 9:
		baseSpeed = 78
	case 10:
		baseSpeed = 45
	case 11:
		baseSpeed = 30
	case 12:
		baseSpeed = 70
	case 13:
		baseSpeed = 40
	case 14:
		baseSpeed = 35
	case 15:
		baseSpeed = 75
	case 16:
		baseSpeed = 56
	case 17:
		baseSpeed = 71
	case 18:
		baseSpeed = 91
	case 19:
		baseSpeed = 72
	case 20:
		baseSpeed = 95
	case 21:
		baseSpeed = 60
	case 22:
		baseSpeed = 100
	case 23:
		baseSpeed = 55
	case 24:
		baseSpeed = 80
	case 25:
		baseSpeed = 90
	case 26:
		baseSpeed = 110
	case 27:
		baseSpeed = 40
	case 28:
		baseSpeed = 65
	case 29:
		baseSpeed = 45
	case 30:
		baseSpeed = 55
	case 31:
		baseSpeed = 76
	case 32:
		baseSpeed = 50
	case 33:
		baseSpeed = 65
	case 34:
		baseSpeed = 85
	case 35:
		baseSpeed = 35
	case 36:
		baseSpeed = 60
	case 37:
		baseSpeed = 65
	case 38:
		baseSpeed = 100
	case 39:
		baseSpeed = 45
	case 40:
		baseSpeed = 70
	case 41:
		baseSpeed = 55
	case 42:
		baseSpeed = 90
	case 43:
		baseSpeed = 55
	case 44:
		baseSpeed = 60
	case 45:
		baseSpeed = 75
	case 46:
		baseSpeed = 55
	case 47:
		baseSpeed = 60
	case 48:
		baseSpeed = 45
	case 49:
		baseSpeed = 70
	case 50:
		baseSpeed = 50
	case 51:
		baseSpeed = 110
	case 52:
		baseSpeed = 90
	case 53:
		baseSpeed = 115
	case 54:
		baseSpeed = 50
	case 55:
		baseSpeed = 80
	case 56:
		baseSpeed = 90
	case 57:
		baseSpeed = 65
	case 58:
		baseSpeed = 70
	case 59:
		baseSpeed = 90
	case 60:
		baseSpeed = 105
	case 61:
		baseSpeed = 120
	case 62:
		baseSpeed = 70
	case 63:
		baseSpeed = 80
	case 64:
		baseSpeed = 90
	case 65:
		baseSpeed = 50
	case 66:
		baseSpeed = 55
	case 67:
		baseSpeed = 70
	case 68:
		baseSpeed = 55
	case 69:
		baseSpeed = 70
	case 70:
		baseSpeed = 80
	case 71:
		baseSpeed = 45
	case 72:
		baseSpeed = 55
	case 73:
		baseSpeed = 60
	case 74:
		baseSpeed = 90
	case 75:
		baseSpeed = 15
	case 76:
		baseSpeed = 30
	case 77:
		baseSpeed = 50
	case 78:
		baseSpeed = 70
	case 79:
		baseSpeed = 60
	case 80:
		baseSpeed = 65
	case 81:
		baseSpeed = 110
	case 82:
		baseSpeed = 85
	case 83:
		baseSpeed = 70
	case 84:
		baseSpeed = 40
	case 85:
		baseSpeed = 75
	case 86:
		baseSpeed = 40
	case 87:
		baseSpeed = 70
	case 88:
		baseSpeed = 80
	case 89:
		baseSpeed = 95
	case 90:
		baseSpeed = 110
	case 91:
		baseSpeed = 70
	case 92:
		baseSpeed = 40
	case 93:
		baseSpeed = 70
	case 94:
		baseSpeed = 25
	case 95:
		baseSpeed = 50
	case 96:
		baseSpeed = 100
	case 97:
		baseSpeed = 140
	case 98:
		baseSpeed = 55
	case 99:
		baseSpeed = 60
	case 100:
		baseSpeed = 50
	case 101:
		baseSpeed = 80
	case 102:
		baseSpeed = 87
	case 103:
		baseSpeed = 100
	case 104:
		baseSpeed = 30
	case 105:
		baseSpeed = 60
	case 106:
		baseSpeed = 60
	case 107:
		baseSpeed = 45
	case 108:
		baseSpeed = 40
	case 109:
		baseSpeed = 70
	case 110:
		baseSpeed = 55
	case 111:
		baseSpeed = 90
	case 112:
		baseSpeed = 60
	case 113:
		baseSpeed = 65
	case 114:
		baseSpeed = 63
	case 115:
		baseSpeed = 68
	case 116:
		baseSpeed = 50
	case 117:
		baseSpeed = 115
	case 118:
		baseSpeed = 60
	case 119:
		baseSpeed = 105
	case 120:
		baseSpeed = 95
	case 121:
		baseSpeed = 87
	case 122:
		baseSpeed = 93
	case 123:
		baseSpeed = 85
	case 124:
		baseSpeed = 100
	case 125:
		baseSpeed = 80
	case 126:
		baseSpeed = 81
	case 127:
		baseSpeed = 60
	case 128:
		baseSpeed = 48
	case 129:
		baseSpeed = 55
	case 130:
		baseSpeed = 65
	case 131:
		baseSpeed = 130
	case 132:
		baseSpeed = 65
	case 133:
		baseSpeed = 60
	case 134:
		baseSpeed = 55
	case 135:
		baseSpeed = 60
	case 136:
		baseSpeed = 55
	case 137:
		baseSpeed = 80
	case 138:
		baseSpeed = 130
	case 139:
		baseSpeed = 30
	case 140:
		baseSpeed = 100
	case 141:
		baseSpeed = 100
	case 142:
		baseSpeed = 90
	case 143:
		baseSpeed = 60
	case 144:
		baseSpeed = 80
	case 145:
		baseSpeed = 100
	case 146:
		baseSpeed = 130
	case 147:
		baseSpeed = 100
	default:
		baseSpeed = 50
	}
	return (baseSpeed*2*level)/100 + 5
}

func pokemonAttacks(id, level int) []Attack {
	var result []Attack
	t := pokemonType(id)
	for i := 0; i < len(attacks); i++ {
		if len(result) == 4 {
			return result
		}
		if attacks[i].Type == t {
			if attacks[i].LevelToLearn < level {
				result = append(result, attacks[i])
			}
		}
	}
	return result
}

func createPokemon(id, level int) Pokemon {
	p := Pokemon{}
	p.ID = id
	p.Name = pokemonName(id)
	p.HP = pokemonHP(id, level)
	p.ActualHP = p.HP
	p.Def = pokemonDEF(id, level)
	p.Speed = pokemonSPEED(id, level)
	p.Level = level
	p.XP = 0
	p.XPToUp = pokemonXPTOUP(level)
	p.LevelToEvolve = pokemonLevelToEvolve(id)
	p.IDEvolution = pokemonEvolutionId(id)
	p.Type = pokemonType(id)
	p.Attacks = pokemonAttacks(id, level)
	return p
}

func applyPokemonXPGain(p *Pokemon, xpGained int) int {
	if xpGained <= 0 {
		return 0
	}

	p.XP += xpGained
	levelsGained := 0

	for p.XP >= p.XPToUp {
		oldHP := p.HP
		p.XP -= p.XPToUp
		p.Level++
		p.XPToUp = pokemonXPTOUP(p.Level)
		p.HP = pokemonHP(p.ID, p.Level)
		p.ActualHP += p.HP - oldHP
		if p.ActualHP > p.HP {
			p.ActualHP = p.HP
		}
		p.Def = pokemonDEF(p.ID, p.Level)
		p.Speed = pokemonSPEED(p.ID, p.Level)
		p.Attacks = pokemonAttacks(p.ID, p.Level)
		levelsGained++
	}

	return levelsGained
}
