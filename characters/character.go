package characters

import (
	"RedHood/util"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"time"
)

type Character struct {
	CurrentImg *ebiten.Image
	idleImages []*ebiten.Image
	runImages  []*ebiten.Image
	castImages []*ebiten.Image
	exitImages []*ebiten.Image

	HP int

	LocX       float64
	LocY       float64
	Velocity   util.Vector
	facing     int
	trackFrame int
	lastCast   time.Time
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
	fmt.Println(len(pool))
	return pool
}

func (c *Character) Draw(screen *ebiten.Image) {
	if c.trackFrame >= len(c.idleImages)*16 {
		c.trackFrame = 0
	} else {
		c.trackFrame += 1
	}
	drawOps := ebiten.DrawImageOptions{}
	distanceErr := float64(c.CurrentImg.Bounds().Dx() * c.facing)

	drawOps.GeoM.Scale(float64(c.facing*2), 2)
	drawOps.GeoM.Translate(c.LocX-distanceErr/1.3, c.LocY)
	screen.DrawImage(c.CurrentImg, &drawOps)

}
