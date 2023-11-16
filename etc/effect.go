package etc

import (
	"RedHood/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Effect struct {
	image       *ebiten.Image
	LocX, LocY  float64
	Offset      util.Point
	Destination util.Point
}
