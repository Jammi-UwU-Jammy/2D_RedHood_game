package characters

import (
	"RedHood/util"
	"github.com/co0p/tankism/lib/collision"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/lafriks/go-tiled"
	"image"
	"time"
)

const (
	STD_TILE_WIDTH = 32
	IMG_PER_SEC    = 20 //assuming 60 frames/sec
)

type MaxStat struct {
	HP, ATK, DEF int
}

type Character struct {
	CurrentImg *ebiten.Image
	idleImages []*ebiten.Image
	runImages  []*ebiten.Image
	atkImages  []*ebiten.Image
	ExitImages []*ebiten.Image

	MaxStat *MaxStat
	HP      int
	ATK     int
	DEF     int

	trackFrame int
	maxFrame   int

	LocX     float64
	LocY     float64
	Velocity util.Vector
	Facing   int
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

func (c *Character) loadSoundAssets(uri string) *audio.Player {
	return util.LoadEmbededSound(uri, 302)
}

func (c *Character) Draw(screen *ebiten.Image) {
	if c.trackFrame >= len(c.idleImages)*IMG_PER_SEC {
		c.trackFrame = 1
	} else {
		c.trackFrame += 1
	}

	drawOps := ebiten.DrawImageOptions{}
	distanceErr := float64(c.CurrentImg.Bounds().Dx() * c.Facing)

	drawOps.GeoM.Scale(float64(c.Facing), 1)
	drawOps.GeoM.Translate(c.LocX-distanceErr/2.2, c.LocY)
	screen.DrawImage(c.CurrentImg, &drawOps)

}

func (c *Character) GetBoundingBox() collision.BoundingBox {
	return collision.BoundingBox{
		X:      c.LocX,
		Y:      c.LocY,
		Width:  20,
		Height: 20,
	}
}

func (c *Character) CollisionVSObjects(obstacles []*tiled.Object) (bool, uint32) {
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
			return true, obs.ID
		}
	}
	return false, 1
}

func (c *Character) CollisionVSBG(tiledMap *tiled.Map) bool {
	tileX := int(c.LocX) / tiledMap.TileWidth
	tileY := int(c.LocY)/tiledMap.TileHeight + 1

	tileToDraw := tiledMap.Layers[0].Tiles[tileY*tiledMap.Width+tileX]
	if tileToDraw.ID != 0 {
		return true
	}
	return false
}
