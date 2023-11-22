package etc

import (
	"RedHood/util"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

const (
	SWORDS_URI = "items/swords/rpg_icons"
)

type Buff struct {
	HP, MP, ATK, DEF int
}

type Container struct {
	items     []*Item
	MaxWeight int
	MaxSlot   int
}

func NewRandomSword(initX, initY float64) *Item {
	randomItemPath := fmt.Sprintf("%s%d%s", SWORDS_URI, rand.Intn(44)+25, ".png")
	fmt.Println(randomItemPath)
	img := util.LoadEmbeddedImage(randomItemPath, 1000)
	buff := &Buff{ATK: rand.Intn(10), HP: rand.Intn(20)}
	e := Effect{CurrentImage: img, LocX: initX, LocY: initY}
	item := Item{Effect: &e, Weight: 10, Buffs: buff}
	return &item
}

type Item struct {
	*Effect

	Status   bool
	Weight   int
	Category int
	Buffs    *Buff
}

func (i *Item) GetImage() *ebiten.Image {
	return i.CurrentImage
}
