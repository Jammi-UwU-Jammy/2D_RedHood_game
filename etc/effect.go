package etc

import (
	"RedHood/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Effect struct {
	CurrentImage *ebiten.Image
	images       []*ebiten.Image
	LocX, LocY   float64
	trackFrame   int
	maxFrame     int

	Offset      util.Point
	Destination util.Point
}

func NewEffect() {

}

func (e *Effect) Update() {

}

func (e *Effect) Draw() {

}
