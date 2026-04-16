package main

type PokemonID int

type PokemonBaseStats struct {
	HP      int
	Defense int
	Speed   int
}

type PokemonEvolution struct {
	Level  int
	NextID PokemonID
}

type PokemonSpecies struct {
	ID        PokemonID
	Name      string
	BaseStats PokemonBaseStats
	Evolution PokemonEvolution
	Type      PokemonType
}

type Pokemon struct {
	ID            PokemonID
	Name          string
	HP            int
	ActualHP      int
	Def           int
	Speed         int
	Level         int
	XPToUp        int
	XP            int
	LevelToEvolve int
	IDEvolution   PokemonID
	Attacks       []Attack
	Type          PokemonType
}

const (
	defaultPokemonBaseHP        = 50
	defaultPokemonBaseDefense   = 50
	defaultPokemonBaseSpeed     = 50
	pokemonStatBaseMultiplier   = 2
	pokemonStatLevelDivisor     = 100
	pokemonHPFlatBonus          = 10
	pokemonCombatStatFlatBonus  = 5
	pokemonXPPerLevelMultiplier = 100
	pokemonAttackSlotLimit      = 4
	unknownPokemonDisplayName   = "Unknown Pokemon"
)

var defaultPokemonSpecies = PokemonSpecies{
	Name: unknownPokemonDisplayName,
	BaseStats: PokemonBaseStats{
		HP:      defaultPokemonBaseHP,
		Defense: defaultPokemonBaseDefense,
		Speed:   defaultPokemonBaseSpeed,
	},
	Type: TypeNormal,
}

// pokemonSpeciesCatalog centralizes all static species data used to build runtime Pokemon instances.
var pokemonSpeciesCatalog = map[PokemonID]PokemonSpecies{
	1:   {ID: 1, Name: "Bulbasaur", BaseStats: PokemonBaseStats{HP: 45, Defense: 49, Speed: 45}, Evolution: PokemonEvolution{Level: 16, NextID: 2}, Type: TypeGrass},
	2:   {ID: 2, Name: "Ivysaur", BaseStats: PokemonBaseStats{HP: 60, Defense: 63, Speed: 60}, Evolution: PokemonEvolution{Level: 32, NextID: 3}, Type: TypeGrass},
	3:   {ID: 3, Name: "Venusaur", BaseStats: PokemonBaseStats{HP: 80, Defense: 83, Speed: 80}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeGrass},
	4:   {ID: 4, Name: "Charmander", BaseStats: PokemonBaseStats{HP: 39, Defense: 43, Speed: 65}, Evolution: PokemonEvolution{Level: 16, NextID: 5}, Type: TypeFire},
	5:   {ID: 5, Name: "Charmeleon", BaseStats: PokemonBaseStats{HP: 58, Defense: 58, Speed: 80}, Evolution: PokemonEvolution{Level: 36, NextID: 6}, Type: TypeFire},
	6:   {ID: 6, Name: "Charizard", BaseStats: PokemonBaseStats{HP: 78, Defense: 65, Speed: 100}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeFire},
	7:   {ID: 7, Name: "Squirtle", BaseStats: PokemonBaseStats{HP: 44, Defense: 48, Speed: 43}, Evolution: PokemonEvolution{Level: 16, NextID: 8}, Type: TypeWater},
	8:   {ID: 8, Name: "Wartortle", BaseStats: PokemonBaseStats{HP: 59, Defense: 65, Speed: 58}, Evolution: PokemonEvolution{Level: 36, NextID: 9}, Type: TypeWater},
	9:   {ID: 9, Name: "Blastoise", BaseStats: PokemonBaseStats{HP: 79, Defense: 80, Speed: 78}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeWater},
	10:  {ID: 10, Name: "Caterpie", BaseStats: PokemonBaseStats{HP: 45, Defense: 30, Speed: 45}, Evolution: PokemonEvolution{Level: 7, NextID: 11}, Type: TypeBug},
	11:  {ID: 11, Name: "Metapod", BaseStats: PokemonBaseStats{HP: 50, Defense: 55, Speed: 30}, Evolution: PokemonEvolution{Level: 10, NextID: 12}, Type: TypeBug},
	12:  {ID: 12, Name: "Butterfree", BaseStats: PokemonBaseStats{HP: 60, Defense: 50, Speed: 70}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeBug},
	13:  {ID: 13, Name: "Weedle", BaseStats: PokemonBaseStats{HP: 40, Defense: 35, Speed: 40}, Evolution: PokemonEvolution{Level: 7, NextID: 14}, Type: TypeBug},
	14:  {ID: 14, Name: "Kakuna", BaseStats: PokemonBaseStats{HP: 45, Defense: 35, Speed: 35}, Evolution: PokemonEvolution{Level: 10, NextID: 15}, Type: TypeBug},
	15:  {ID: 15, Name: "Beedrill", BaseStats: PokemonBaseStats{HP: 65, Defense: 80, Speed: 75}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeBug},
	16:  {ID: 16, Name: "Pidgey", BaseStats: PokemonBaseStats{HP: 40, Defense: 40, Speed: 56}, Evolution: PokemonEvolution{Level: 18, NextID: 17}, Type: TypeNormal},
	17:  {ID: 17, Name: "Pidgeotto", BaseStats: PokemonBaseStats{HP: 63, Defense: 55, Speed: 71}, Evolution: PokemonEvolution{Level: 36, NextID: 18}, Type: TypeNormal},
	18:  {ID: 18, Name: "Pidgeot", BaseStats: PokemonBaseStats{HP: 83, Defense: 75, Speed: 91}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeNormal},
	19:  {ID: 19, Name: "Rattata", BaseStats: PokemonBaseStats{HP: 30, Defense: 30, Speed: 72}, Evolution: PokemonEvolution{Level: 20, NextID: 20}, Type: TypeNormal},
	20:  {ID: 20, Name: "Raticate", BaseStats: PokemonBaseStats{HP: 55, Defense: 50, Speed: 95}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeNormal},
	21:  {ID: 21, Name: "Spearow", BaseStats: PokemonBaseStats{HP: 40, Defense: 30, Speed: 60}, Evolution: PokemonEvolution{Level: 20, NextID: 22}, Type: TypeNormal},
	22:  {ID: 22, Name: "Fearow", BaseStats: PokemonBaseStats{HP: 65, Defense: 65, Speed: 100}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeNormal},
	23:  {ID: 23, Name: "Ekans", BaseStats: PokemonBaseStats{HP: 35, Defense: 35, Speed: 55}, Evolution: PokemonEvolution{Level: 22, NextID: 24}, Type: TypePoison},
	24:  {ID: 24, Name: "Arbok", BaseStats: PokemonBaseStats{HP: 60, Defense: 49, Speed: 80}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypePoison},
	25:  {ID: 25, Name: "Pikachu", BaseStats: PokemonBaseStats{HP: 35, Defense: 40, Speed: 90}, Evolution: PokemonEvolution{Level: 16, NextID: 26}, Type: TypeElectric},
	26:  {ID: 26, Name: "Raichu", BaseStats: PokemonBaseStats{HP: 60, Defense: 50, Speed: 110}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeElectric},
	27:  {ID: 27, Name: "Sandshrew", BaseStats: PokemonBaseStats{HP: 50, Defense: 75, Speed: 40}, Evolution: PokemonEvolution{Level: 22, NextID: 28}, Type: TypeGround},
	28:  {ID: 28, Name: "Sandslash", BaseStats: PokemonBaseStats{HP: 75, Defense: 100, Speed: 65}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeGround},
	29:  {ID: 29, Name: "Nidoran F", BaseStats: PokemonBaseStats{HP: 55, Defense: 40, Speed: 45}, Evolution: PokemonEvolution{Level: 16, NextID: 30}, Type: TypePoison},
	30:  {ID: 30, Name: "Nidorina", BaseStats: PokemonBaseStats{HP: 70, Defense: 55, Speed: 55}, Evolution: PokemonEvolution{Level: 31, NextID: 31}, Type: TypePoison},
	31:  {ID: 31, Name: "Nidoqueen", BaseStats: PokemonBaseStats{HP: 90, Defense: 70, Speed: 76}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeGround},
	32:  {ID: 32, Name: "Nidoran M", BaseStats: PokemonBaseStats{HP: 46, Defense: 40, Speed: 50}, Evolution: PokemonEvolution{Level: 16, NextID: 33}, Type: TypePoison},
	33:  {ID: 33, Name: "Nidorino", BaseStats: PokemonBaseStats{HP: 61, Defense: 55, Speed: 65}, Evolution: PokemonEvolution{Level: 31, NextID: 34}, Type: TypePoison},
	34:  {ID: 34, Name: "Nidoking", BaseStats: PokemonBaseStats{HP: 81, Defense: 80, Speed: 85}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeGround},
	35:  {ID: 35, Name: "Clefairy", BaseStats: PokemonBaseStats{HP: 70, Defense: 48, Speed: 35}, Evolution: PokemonEvolution{Level: 35, NextID: 36}, Type: TypeFairy},
	36:  {ID: 36, Name: "Clefable", BaseStats: PokemonBaseStats{HP: 95, Defense: 70, Speed: 60}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeFairy},
	37:  {ID: 37, Name: "Vulpix", BaseStats: PokemonBaseStats{HP: 38, Defense: 40, Speed: 65}, Evolution: PokemonEvolution{Level: 20, NextID: 38}, Type: TypeFire},
	38:  {ID: 38, Name: "Ninetales", BaseStats: PokemonBaseStats{HP: 73, Defense: 75, Speed: 100}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeFire},
	39:  {ID: 39, Name: "Jigglypuff", BaseStats: PokemonBaseStats{HP: 115, Defense: 30, Speed: 45}, Evolution: PokemonEvolution{Level: 18, NextID: 40}, Type: TypeNormal},
	40:  {ID: 40, Name: "Wigglytuff", BaseStats: PokemonBaseStats{HP: 140, Defense: 70, Speed: 70}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeFairy},
	41:  {ID: 41, Name: "Zubat", BaseStats: PokemonBaseStats{HP: 40, Defense: 45, Speed: 55}, Evolution: PokemonEvolution{Level: 22, NextID: 42}, Type: TypePoison},
	42:  {ID: 42, Name: "Golbat", BaseStats: PokemonBaseStats{HP: 75, Defense: 70, Speed: 90}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypePoison},
	43:  {ID: 43, Name: "Oddish", BaseStats: PokemonBaseStats{HP: 45, Defense: 55, Speed: 55}, Evolution: PokemonEvolution{Level: 21, NextID: 44}, Type: TypeGrass},
	44:  {ID: 44, Name: "Gloom", BaseStats: PokemonBaseStats{HP: 60, Defense: 50, Speed: 60}, Evolution: PokemonEvolution{Level: 0, NextID: 45}, Type: TypeGrass},
	45:  {ID: 45, Name: "Vileplume", BaseStats: PokemonBaseStats{HP: 75, Defense: 70, Speed: 75}, Evolution: PokemonEvolution{Level: 21, NextID: 0}, Type: TypeGrass},
	46:  {ID: 46, Name: "Paras", BaseStats: PokemonBaseStats{HP: 35, Defense: 55, Speed: 55}, Evolution: PokemonEvolution{Level: 24, NextID: 47}, Type: TypeBug},
	47:  {ID: 47, Name: "Parasect", BaseStats: PokemonBaseStats{HP: 60, Defense: 60, Speed: 60}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeBug},
	48:  {ID: 48, Name: "Venonat", BaseStats: PokemonBaseStats{HP: 60, Defense: 50, Speed: 45}, Evolution: PokemonEvolution{Level: 31, NextID: 49}, Type: TypeBug},
	49:  {ID: 49, Name: "Venomoth", BaseStats: PokemonBaseStats{HP: 70, Defense: 60, Speed: 70}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeBug},
	50:  {ID: 50, Name: "Diglett", BaseStats: PokemonBaseStats{HP: 10, Defense: 25, Speed: 50}, Evolution: PokemonEvolution{Level: 26, NextID: 51}, Type: TypeGround},
	51:  {ID: 51, Name: "Dugtrio", BaseStats: PokemonBaseStats{HP: 35, Defense: 50, Speed: 110}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeGround},
	52:  {ID: 52, Name: "Meowth", BaseStats: PokemonBaseStats{HP: 40, Defense: 45, Speed: 90}, Evolution: PokemonEvolution{Level: 28, NextID: 53}, Type: TypeNormal},
	53:  {ID: 53, Name: "Persian", BaseStats: PokemonBaseStats{HP: 65, Defense: 60, Speed: 115}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeNormal},
	54:  {ID: 54, Name: "Psyduck", BaseStats: PokemonBaseStats{HP: 50, Defense: 50, Speed: 50}, Evolution: PokemonEvolution{Level: 33, NextID: 55}, Type: TypeWater},
	55:  {ID: 55, Name: "Golduck", BaseStats: PokemonBaseStats{HP: 80, Defense: 80, Speed: 80}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeWater},
	56:  {ID: 56, Name: "Poliwag", BaseStats: PokemonBaseStats{HP: 50, Defense: 50, Speed: 90}, Evolution: PokemonEvolution{Level: 25, NextID: 57}, Type: TypeWater},
	57:  {ID: 57, Name: "Poliwhirl", BaseStats: PokemonBaseStats{HP: 65, Defense: 65, Speed: 65}, Evolution: PokemonEvolution{Level: 36, NextID: 58}, Type: TypeWater},
	58:  {ID: 58, Name: "Poliwrath", BaseStats: PokemonBaseStats{HP: 90, Defense: 90, Speed: 70}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeWater},
	59:  {ID: 59, Name: "Abra", BaseStats: PokemonBaseStats{HP: 25, Defense: 20, Speed: 90}, Evolution: PokemonEvolution{Level: 16, NextID: 60}, Type: TypePsychic},
	60:  {ID: 60, Name: "Kadabra", BaseStats: PokemonBaseStats{HP: 40, Defense: 25, Speed: 105}, Evolution: PokemonEvolution{Level: 0, NextID: 61}, Type: TypePsychic},
	61:  {ID: 61, Name: "Alakazam", BaseStats: PokemonBaseStats{HP: 55, Defense: 40, Speed: 120}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypePsychic},
	62:  {ID: 62, Name: "Machop", BaseStats: PokemonBaseStats{HP: 70, Defense: 70, Speed: 70}, Evolution: PokemonEvolution{Level: 28, NextID: 63}, Type: TypeFighting},
	63:  {ID: 63, Name: "Machoke", BaseStats: PokemonBaseStats{HP: 80, Defense: 80, Speed: 80}, Evolution: PokemonEvolution{Level: 35, NextID: 64}, Type: TypeFighting},
	64:  {ID: 64, Name: "Machamp", BaseStats: PokemonBaseStats{HP: 90, Defense: 90, Speed: 90}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeFighting},
	65:  {ID: 65, Name: "Bellsprout", BaseStats: PokemonBaseStats{HP: 50, Defense: 35, Speed: 50}, Evolution: PokemonEvolution{Level: 21, NextID: 66}, Type: TypeGrass},
	66:  {ID: 66, Name: "Weepinbell", BaseStats: PokemonBaseStats{HP: 65, Defense: 50, Speed: 55}, Evolution: PokemonEvolution{Level: 35, NextID: 67}, Type: TypeGrass},
	67:  {ID: 67, Name: "Victreebel", BaseStats: PokemonBaseStats{HP: 80, Defense: 65, Speed: 70}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeGrass},
	68:  {ID: 68, Name: "Tentacool", BaseStats: PokemonBaseStats{HP: 40, Defense: 40, Speed: 55}, Evolution: PokemonEvolution{Level: 30, NextID: 69}, Type: TypeWater},
	69:  {ID: 69, Name: "Tentacruel", BaseStats: PokemonBaseStats{HP: 80, Defense: 70, Speed: 70}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeWater},
	70:  {ID: 70, Name: "Geodude", BaseStats: PokemonBaseStats{HP: 40, Defense: 30, Speed: 80}, Evolution: PokemonEvolution{Level: 25, NextID: 71}, Type: TypeRock},
	71:  {ID: 71, Name: "Graveler", BaseStats: PokemonBaseStats{HP: 55, Defense: 100, Speed: 45}, Evolution: PokemonEvolution{Level: 0, NextID: 72}, Type: TypeRock},
	72:  {ID: 72, Name: "Golem", BaseStats: PokemonBaseStats{HP: 80, Defense: 120, Speed: 55}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeRock},
	73:  {ID: 73, Name: "Ponyta", BaseStats: PokemonBaseStats{HP: 50, Defense: 55, Speed: 60}, Evolution: PokemonEvolution{Level: 40, NextID: 74}, Type: TypeFire},
	74:  {ID: 74, Name: "Rapidash", BaseStats: PokemonBaseStats{HP: 65, Defense: 80, Speed: 90}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeFire},
	75:  {ID: 75, Name: "Slowpoke", BaseStats: PokemonBaseStats{HP: 90, Defense: 65, Speed: 15}, Evolution: PokemonEvolution{Level: 30, NextID: 76}, Type: TypeWater},
	76:  {ID: 76, Name: "Slowbro", BaseStats: PokemonBaseStats{HP: 95, Defense: 80, Speed: 30}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeWater},
	77:  {ID: 77, Name: "Magnemite", BaseStats: PokemonBaseStats{HP: 25, Defense: 60, Speed: 50}, Evolution: PokemonEvolution{Level: 30, NextID: 78}, Type: TypeElectric},
	78:  {ID: 78, Name: "Magneton", BaseStats: PokemonBaseStats{HP: 50, Defense: 95, Speed: 70}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeElectric},
	79:  {ID: 79, Name: "Farfetch'd", BaseStats: PokemonBaseStats{HP: 52, Defense: 50, Speed: 60}, Evolution: PokemonEvolution{Level: 35, NextID: 0}, Type: TypeNormal},
	80:  {ID: 80, Name: "Doduo", BaseStats: PokemonBaseStats{HP: 35, Defense: 35, Speed: 65}, Evolution: PokemonEvolution{Level: 0, NextID: 81}, Type: TypeNormal},
	81:  {ID: 81, Name: "Dodrio", BaseStats: PokemonBaseStats{HP: 60, Defense: 60, Speed: 110}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeNormal},
	82:  {ID: 82, Name: "Seel", BaseStats: PokemonBaseStats{HP: 65, Defense: 85, Speed: 85}, Evolution: PokemonEvolution{Level: 34, NextID: 83}, Type: TypeWater},
	83:  {ID: 83, Name: "Dewgong", BaseStats: PokemonBaseStats{HP: 90, Defense: 70, Speed: 70}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeWater},
	84:  {ID: 84, Name: "Grimer", BaseStats: PokemonBaseStats{HP: 80, Defense: 50, Speed: 40}, Evolution: PokemonEvolution{Level: 40, NextID: 85}, Type: TypePoison},
	85:  {ID: 85, Name: "Muk", BaseStats: PokemonBaseStats{HP: 105, Defense: 75, Speed: 75}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypePoison},
	86:  {ID: 86, Name: "Shellder", BaseStats: PokemonBaseStats{HP: 30, Defense: 55, Speed: 40}, Evolution: PokemonEvolution{Level: 30, NextID: 87}, Type: TypeWater},
	87:  {ID: 87, Name: "Cloyster", BaseStats: PokemonBaseStats{HP: 50, Defense: 95, Speed: 70}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeWater},
	88:  {ID: 88, Name: "Gastly", BaseStats: PokemonBaseStats{HP: 30, Defense: 30, Speed: 80}, Evolution: PokemonEvolution{Level: 25, NextID: 89}, Type: TypeGhost},
	89:  {ID: 89, Name: "Haunter", BaseStats: PokemonBaseStats{HP: 45, Defense: 45, Speed: 95}, Evolution: PokemonEvolution{Level: 0, NextID: 90}, Type: TypeGhost},
	90:  {ID: 90, Name: "Gengar", BaseStats: PokemonBaseStats{HP: 60, Defense: 60, Speed: 110}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeGhost},
	91:  {ID: 91, Name: "Onix", BaseStats: PokemonBaseStats{HP: 35, Defense: 160, Speed: 70}, Evolution: PokemonEvolution{Level: 30, NextID: 0}, Type: TypeRock},
	92:  {ID: 92, Name: "Drowzee", BaseStats: PokemonBaseStats{HP: 60, Defense: 55, Speed: 40}, Evolution: PokemonEvolution{Level: 0, NextID: 93}, Type: TypePsychic},
	93:  {ID: 93, Name: "Hypno", BaseStats: PokemonBaseStats{HP: 80, Defense: 70, Speed: 70}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypePsychic},
	94:  {ID: 94, Name: "Krabby", BaseStats: PokemonBaseStats{HP: 30, Defense: 50, Speed: 25}, Evolution: PokemonEvolution{Level: 30, NextID: 95}, Type: TypeWater},
	95:  {ID: 95, Name: "Kingler", BaseStats: PokemonBaseStats{HP: 55, Defense: 75, Speed: 50}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeWater},
	96:  {ID: 96, Name: "Voltorb", BaseStats: PokemonBaseStats{HP: 40, Defense: 40, Speed: 100}, Evolution: PokemonEvolution{Level: 30, NextID: 97}, Type: TypeElectric},
	97:  {ID: 97, Name: "Electrode", BaseStats: PokemonBaseStats{HP: 60, Defense: 60, Speed: 140}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeElectric},
	98:  {ID: 98, Name: "Exeggcute", BaseStats: PokemonBaseStats{HP: 60, Defense: 55, Speed: 55}, Evolution: PokemonEvolution{Level: 0, NextID: 99}, Type: TypeGrass},
	99:  {ID: 99, Name: "Exeggutor", BaseStats: PokemonBaseStats{HP: 95, Defense: 70, Speed: 60}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeGrass},
	100: {ID: 100, Name: "Cubone", BaseStats: PokemonBaseStats{HP: 50, Defense: 50, Speed: 50}, Evolution: PokemonEvolution{Level: 28, NextID: 101}, Type: TypeGround},
	101: {ID: 101, Name: "Marowak", BaseStats: PokemonBaseStats{HP: 60, Defense: 80, Speed: 80}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeGround},
	102: {ID: 102, Name: "Hitmonlee", BaseStats: PokemonBaseStats{HP: 50, Defense: 90, Speed: 87}, Evolution: PokemonEvolution{Level: 20, NextID: 0}, Type: TypeFighting},
	103: {ID: 103, Name: "Hitmonchan", BaseStats: PokemonBaseStats{HP: 50, Defense: 110, Speed: 100}, Evolution: PokemonEvolution{Level: 20, NextID: 0}, Type: TypeFighting},
	104: {ID: 104, Name: "Lickitung", BaseStats: PokemonBaseStats{HP: 90, Defense: 70, Speed: 30}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeNormal},
	105: {ID: 105, Name: "Koffing", BaseStats: PokemonBaseStats{HP: 40, Defense: 70, Speed: 60}, Evolution: PokemonEvolution{Level: 0, NextID: 106}, Type: TypePoison},
	106: {ID: 106, Name: "Weezing", BaseStats: PokemonBaseStats{HP: 65, Defense: 120, Speed: 60}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypePoison},
	107: {ID: 107, Name: "Rhyhorn", BaseStats: PokemonBaseStats{HP: 80, Defense: 95, Speed: 45}, Evolution: PokemonEvolution{Level: 42, NextID: 108}, Type: TypeGround},
	108: {ID: 108, Name: "Rhydon", BaseStats: PokemonBaseStats{HP: 105, Defense: 130, Speed: 40}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeGround},
	109: {ID: 109, Name: "Chansey", BaseStats: PokemonBaseStats{HP: 250, Defense: 30, Speed: 70}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeNormal},
	110: {ID: 110, Name: "Tangela", BaseStats: PokemonBaseStats{HP: 65, Defense: 60, Speed: 55}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeGrass},
	111: {ID: 111, Name: "Kangaskhan", BaseStats: PokemonBaseStats{HP: 105, Defense: 80, Speed: 90}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeNormal},
	112: {ID: 112, Name: "Horsea", BaseStats: PokemonBaseStats{HP: 30, Defense: 50, Speed: 60}, Evolution: PokemonEvolution{Level: 32, NextID: 113}, Type: TypeWater},
	113: {ID: 113, Name: "Seadra", BaseStats: PokemonBaseStats{HP: 55, Defense: 70, Speed: 65}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeWater},
	114: {ID: 114, Name: "Goldeen", BaseStats: PokemonBaseStats{HP: 45, Defense: 45, Speed: 63}, Evolution: PokemonEvolution{Level: 33, NextID: 115}, Type: TypeWater},
	115: {ID: 115, Name: "Seaking", BaseStats: PokemonBaseStats{HP: 80, Defense: 70, Speed: 68}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeWater},
	116: {ID: 116, Name: "Staryu", BaseStats: PokemonBaseStats{HP: 30, Defense: 30, Speed: 50}, Evolution: PokemonEvolution{Level: 0, NextID: 117}, Type: TypeWater},
	117: {ID: 117, Name: "Starmie", BaseStats: PokemonBaseStats{HP: 60, Defense: 60, Speed: 115}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeWater},
	118: {ID: 118, Name: "Mr. Mime", BaseStats: PokemonBaseStats{HP: 40, Defense: 70, Speed: 60}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypePsychic},
	119: {ID: 119, Name: "Scyther", BaseStats: PokemonBaseStats{HP: 70, Defense: 55, Speed: 105}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeBug},
	120: {ID: 120, Name: "Jynx", BaseStats: PokemonBaseStats{HP: 65, Defense: 35, Speed: 95}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeIce},
	121: {ID: 121, Name: "Electabuzz", BaseStats: PokemonBaseStats{HP: 65, Defense: 57, Speed: 87}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeElectric},
	122: {ID: 122, Name: "Magmar", BaseStats: PokemonBaseStats{HP: 65, Defense: 57, Speed: 93}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeFire},
	123: {ID: 123, Name: "Pinsir", BaseStats: PokemonBaseStats{HP: 65, Defense: 75, Speed: 85}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeBug},
	124: {ID: 124, Name: "Tauros", BaseStats: PokemonBaseStats{HP: 75, Defense: 75, Speed: 100}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeNormal},
	125: {ID: 125, Name: "Magikarp", BaseStats: PokemonBaseStats{HP: 20, Defense: 20, Speed: 80}, Evolution: PokemonEvolution{Level: 0, NextID: 126}, Type: TypeWater},
	126: {ID: 126, Name: "Gyarados", BaseStats: PokemonBaseStats{HP: 95, Defense: 79, Speed: 81}, Evolution: PokemonEvolution{Level: 20, NextID: 0}, Type: TypeWater},
	127: {ID: 127, Name: "Lapras", BaseStats: PokemonBaseStats{HP: 130, Defense: 85, Speed: 60}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeWater},
	128: {ID: 128, Name: "Ditto", BaseStats: PokemonBaseStats{HP: 35, Defense: 70, Speed: 48}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeNormal},
	129: {ID: 129, Name: "Eevee", BaseStats: PokemonBaseStats{HP: 55, Defense: 50, Speed: 55}, Evolution: PokemonEvolution{Level: 0, NextID: 130}, Type: TypeNormal},
	130: {ID: 130, Name: "Vaporeon", BaseStats: PokemonBaseStats{HP: 130, Defense: 60, Speed: 65}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeWater},
	131: {ID: 131, Name: "Jolteon", BaseStats: PokemonBaseStats{HP: 65, Defense: 65, Speed: 130}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeElectric},
	132: {ID: 132, Name: "Flareon", BaseStats: PokemonBaseStats{HP: 65, Defense: 65, Speed: 65}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeFire},
	133: {ID: 133, Name: "Porygon", BaseStats: PokemonBaseStats{HP: 65, Defense: 65, Speed: 60}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeNormal},
	134: {ID: 134, Name: "Omanyte", BaseStats: PokemonBaseStats{HP: 35, Defense: 55, Speed: 55}, Evolution: PokemonEvolution{Level: 0, NextID: 135}, Type: TypeRock},
	135: {ID: 135, Name: "Omastar", BaseStats: PokemonBaseStats{HP: 70, Defense: 70, Speed: 60}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeRock},
	136: {ID: 136, Name: "Kabuto", BaseStats: PokemonBaseStats{HP: 30, Defense: 55, Speed: 55}, Evolution: PokemonEvolution{Level: 0, NextID: 137}, Type: TypeRock},
	137: {ID: 137, Name: "Kabutops", BaseStats: PokemonBaseStats{HP: 60, Defense: 65, Speed: 80}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeRock},
	138: {ID: 138, Name: "Aerodactyl", BaseStats: PokemonBaseStats{HP: 80, Defense: 60, Speed: 130}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeRock},
	139: {ID: 139, Name: "Snorlax", BaseStats: PokemonBaseStats{HP: 110, Defense: 110, Speed: 30}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeNormal},
	140: {ID: 140, Name: "Articuno", BaseStats: PokemonBaseStats{HP: 90, Defense: 100, Speed: 100}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeIce},
	141: {ID: 141, Name: "Zapdos", BaseStats: PokemonBaseStats{HP: 90, Defense: 90, Speed: 100}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeElectric},
	142: {ID: 142, Name: "Moltres", BaseStats: PokemonBaseStats{HP: 90, Defense: 90, Speed: 90}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeFire},
	143: {ID: 143, Name: "Dratini", BaseStats: PokemonBaseStats{HP: 41, Defense: 60, Speed: 60}, Evolution: PokemonEvolution{Level: 30, NextID: 144}, Type: TypeDragon},
	144: {ID: 144, Name: "Dragonair", BaseStats: PokemonBaseStats{HP: 61, Defense: 65, Speed: 80}, Evolution: PokemonEvolution{Level: 0, NextID: 145}, Type: TypeDragon},
	145: {ID: 145, Name: "Dragonite", BaseStats: PokemonBaseStats{HP: 91, Defense: 95, Speed: 100}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypeDragon},
	146: {ID: 146, Name: "Mewtwo", BaseStats: PokemonBaseStats{HP: 106, Defense: 90, Speed: 130}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypePsychic},
	147: {ID: 147, Name: "Mew", BaseStats: PokemonBaseStats{HP: 100, Defense: 100, Speed: 100}, Evolution: PokemonEvolution{Level: 0, NextID: 0}, Type: TypePsychic},
}

func pokemonSpecies(id PokemonID) PokemonSpecies {
	species, found := pokemonSpeciesCatalog[id]
	if found {
		return species
	}

	unknownSpecies := defaultPokemonSpecies
	unknownSpecies.ID = id
	return unknownSpecies
}

func calculatePokemonHP(baseHP, level int) int {
	return (baseHP*pokemonStatBaseMultiplier*level)/pokemonStatLevelDivisor + level + pokemonHPFlatBonus
}

func calculatePokemonCombatStat(baseStat, level int) int {
	return (baseStat*pokemonStatBaseMultiplier*level)/pokemonStatLevelDivisor + pokemonCombatStatFlatBonus
}

func pokemonXPToUp(level int) int {
	return level * pokemonXPPerLevelMultiplier
}

func (species PokemonSpecies) attacksForLevel(level int) []Attack {
	result := make([]Attack, 0, pokemonAttackSlotLimit)

	for _, attack := range attacks {
		if len(result) == pokemonAttackSlotLimit {
			break
		}

		if attack.Type == species.Type && attack.LevelToLearn < level {
			result = append(result, attack)
		}
	}

	return result
}

func (species PokemonSpecies) create(level int) Pokemon {
	hp := calculatePokemonHP(species.BaseStats.HP, level)

	return Pokemon{
		ID:            species.ID,
		Name:          species.Name,
		HP:            hp,
		ActualHP:      hp,
		Def:           calculatePokemonCombatStat(species.BaseStats.Defense, level),
		Speed:         calculatePokemonCombatStat(species.BaseStats.Speed, level),
		Level:         level,
		XPToUp:        pokemonXPToUp(level),
		XP:            0,
		LevelToEvolve: species.Evolution.Level,
		IDEvolution:   species.Evolution.NextID,
		Attacks:       species.attacksForLevel(level),
		Type:          species.Type,
	}
}

func createPokemon(id PokemonID, level int) Pokemon {
	return pokemonSpecies(id).create(level)
}

func applyPokemonXPGain(p *Pokemon, xpGained int) int {
	if xpGained <= 0 {
		return 0
	}

	p.XP += xpGained
	levelsGained := 0
	species := pokemonSpecies(p.ID)

	for p.XP >= p.XPToUp {
		oldHP := p.HP
		p.XP -= p.XPToUp
		p.Level++
		p.XPToUp = pokemonXPToUp(p.Level)
		p.HP = calculatePokemonHP(species.BaseStats.HP, p.Level)
		p.ActualHP += p.HP - oldHP
		if p.ActualHP > p.HP {
			p.ActualHP = p.HP
		}
		p.Def = calculatePokemonCombatStat(species.BaseStats.Defense, p.Level)
		p.Speed = calculatePokemonCombatStat(species.BaseStats.Speed, p.Level)
		p.Attacks = species.attacksForLevel(p.Level)
		levelsGained++
	}

	return levelsGained
}
