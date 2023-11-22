package etc

import (
	"RedHood/util"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

const (
	IMG_PER_SEC = 10
)

type Effect struct {
	CurrentImage *ebiten.Image
	images       []*ebiten.Image
	LocX, LocY   float64
	trackFrame   int
	facing       int
	maxFrame     int
	speed        int
	created      time.Time

	Offset      util.Point
	Destination int

	IsDone bool
}

func NewLastingEffect(images []*ebiten.Image, X, Y float64, facing int) *Effect {
	e := Effect{images: images, CurrentImage: images[0],
		LocX: X, LocY: Y, facing: facing,
		IsDone: false,
	}
	return &e
}

func NewTravellingEffect(images []*ebiten.Image, off util.Point, dest, speed, facing int) *Effect {
	return &Effect{
		CurrentImage: images[0],
		images:       images,
		LocX:         float64(off.X),
		LocY:         float64(off.Y),
		speed:        speed,
		facing:       facing,
		Destination:  dest,
		IsDone:       false,
	}
}

func (e *Effect) Update() {
	if e.trackFrame >= len(e.images)*IMG_PER_SEC {
		if e.speed == 0 {
			e.IsDone = true
		}
	} else {
		e.trackFrame += 1
	}

	if e.speed != 0 {
		e.LocX += float64(e.speed)
		e.LocY += float64(e.speed)
	}
	e.CurrentImage = e.images[(e.trackFrame-1)/IMG_PER_SEC]
}

func (e *Effect) Draw(screen *ebiten.Image) {
	drawOps := ebiten.DrawImageOptions{}
	distanceErr := float64(e.CurrentImage.Bounds().Dx() * e.facing)

	drawOps.GeoM.Scale(float64(e.facing), 1)
	drawOps.GeoM.Translate(e.LocX-distanceErr/2.2, e.LocY)
	screen.DrawImage(e.CurrentImage, &drawOps)
}
