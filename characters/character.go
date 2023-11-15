package characters

import (
	"RedHood/util"
	"fmt"
	"github.com/co0p/tankism/lib/collision"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
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
	//fmt.Println(len(pool))
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

	drawOps.GeoM.Scale(float64(c.facing), 1)
	drawOps.GeoM.Translate(c.LocX-distanceErr/2.2, c.LocY)
	screen.DrawImage(c.CurrentImg, &drawOps)

}

func (c *Character) GetBoundingBox() collision.BoundingBox {
	return collision.BoundingBox{
		X:      c.LocX,
		Y:      c.LocY,
		Width:  10, //float64(c.CurrentImg.Bounds().Dx()),
		Height: 10, //float64(c.CurrentImg.Bounds().Dy()),
	}
}

func (c *Character) collisionVSBG(obstacles []*tiled.Object) bool {
	for _, obs := range obstacles {
		//obsBox := util.AABB{
		//	MinX: obs.X,
		//	MinY: obs.Y,
		//	MaxX: obs.X + obs.Width,
		//	MaxY: obs.Y + obs.Height,
		//}
		//charBox := util.AABB{
		//	MinX: c.LocX,
		//	MinY: c.LocY,
		//	MaxX: c.LocX + float64(c.CurrentImg.Bounds().Dx()),
		//	MaxY: c.LocY + float64(c.CurrentImg.Bounds().Dy()),
		//}
		//if util.IfBoxesCollided(obsBox, charBox) {
		//	return true
		//}
		box := collision.BoundingBox{
			X:      obs.X,
			Y:      obs.Y,
			Width:  obs.Width,
			Height: obs.Height,
		}

		if util.IfCollided(c.GetBoundingBox(), box) {
			fmt.Println(obs.X, " : ", obs.Y)
			fmt.Println("Collided with an obstacle: ", obs.ID)
			fmt.Println(c.LocX, " : ", c.LocY)
			return true
		}
	}
	return false
}
