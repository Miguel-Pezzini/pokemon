package main

type PokemonType int

const (
	TypeNormal PokemonType = iota
	TypeFire
	TypeWater
	TypeGrass
	TypeElectric
	TypePsychic
	TypeIce
	TypeDragon
	TypeFairy
	TypeBug
	TypeRock
	TypeGhost
	TypeDark
	TypeFighting
	TypePoison
	TypeGround
	TypeFlying
	TypeSteel
)

type ItemType int

const (
	ItemHealing ItemType = iota
	ItemPokeBalls
	ItemStatBoosting
	ItemEvolution
	ItemEquipable
	ItemUtility
	ItemBattle
)
