package etc

import "github.com/hajimehoshi/ebiten/v2"

type Buff struct {
	HP, MP, ATK, DEF int
}

type Container struct {
	items     []*Item
	MaxWeight int
	MaxSlot   int
}

type Item struct {
	image *ebiten.Image

	Weight   int
	Category int
	Buffs    *Buff
}
