package characters

import (
	//"RedHood/environments"
	"RedHood/util"
	"github.com/co0p/tankism/lib/collision"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"image"
	"time"
)

const (
	STD_TILE_WIDTH = 32
	IMG_PER_SEC    = 30 //assuming 60 frames/sec
)

type Character struct {
	CurrentImg *ebiten.Image
	idleImages []*ebiten.Image
	runImages  []*ebiten.Image
	atkImages  []*ebiten.Image
	exitImages []*ebiten.Image

	HP         int
	trackFrame int
	maxFrame   int

	LocX     float64
	LocY     float64
	Velocity util.Vector
	facing   int
	lastCast time.Time

	*collision.BoundingBox
}

func (c *Character) loadImageAssets(uri string, offset util.Point, width, height int) []*ebiten.Image {
	img := util.LoadEmbeddedImage(uri, 301)
	imgWidth := img.Bounds().Dx()

	var pool []*ebiten.Image
	for i := 0; i < imgWidth/width; i++ {
		subImg := img.SubImage(image.Rect(offset.X, offset.Y, offset.X+width, height))
		subImage := ebiten.NewImageFromImage(subImg)
		pool = append(pool, subImage)
		offset.X += width
	}
	return pool
}

func (c *Character) Draw(screen *ebiten.Image) {
	if c.trackFrame >= len(c.idleImages)*IMG_PER_SEC {
		c.trackFrame = 1
	} else {
		c.trackFrame += 1
	}
	drawOps := ebiten.DrawImageOptions{}
	distanceErr := float64(c.CurrentImg.Bounds().Dx() * c.facing)

	drawOps.GeoM.Scale(float64(c.facing), 1)
	drawOps.GeoM.Translate(c.LocX-distanceErr/2.2, c.LocY)
	screen.DrawImage(c.CurrentImg, &drawOps)

}

func (c *Character) GetBoundingBox() collision.BoundingBox {
	return collision.BoundingBox{
		X:      c.LocX,
		Y:      c.LocY,
		Width:  10,
		Height: 10,
	}
}

func (c *Character) collisionVSBG(obstacles []*tiled.Object) bool {
	for _, obs := range obstacles {
		box := collision.BoundingBox{
			X:      obs.X,
			Y:      obs.Y - STD_TILE_WIDTH,
			Width:  obs.Width,
			Height: obs.Height,
		}
		if util.IfCollided(c.GetBoundingBox(), box) {
			//fmt.Println(obs.X, " : ", obs.Y)
			//fmt.Println("Collided with an obstacle: ", obs.ID)
			//fmt.Println(c.LocX, " : ", c.LocY)
			return true
		}
	}
	return false
}
