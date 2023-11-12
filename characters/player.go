package characters

import (
	"RedHood/util"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"time"
)

const (
	SPEED           = 4
	IDLE_IMAGES     = 0
	IDLE_IMAGES_URI = "player/images/idle-sheet.png"
	RUN_IMAGES_URI  = "player/images/run-sheet.png"
	CAST_IMAGES_URI = "player/images/atk-sheet.png"
)

type Character struct {
	CurrentImg *ebiten.Image
	idleImages []*ebiten.Image
	LocX       float64
	LocY       float64
	facing     int
	trackFrame int
	lastCast   time.Time
}

func NewPlayer() *Player {
	character := &Character{facing: 1, lastCast: time.Now()}
	player := Player{Character: character}

	player.idleImages = player.loadImageAssets(IDLE_IMAGES_URI, util.Point{X: 0, Y: 0}, 80, 80)
	player.runImages = player.loadImageAssets(RUN_IMAGES_URI, util.Point{X: 0, Y: 0}, 80, 80)
	player.castImages = player.loadImageAssets(CAST_IMAGES_URI, util.Point{X: 0, Y: 0}, 80, 80)
	return &player
}

type Player struct {
	*Character

	runImages  []*ebiten.Image
	castImages []*ebiten.Image
}

func (p *Player) loadImageAssets(uri string, offset util.Point, width, height int) []*ebiten.Image {
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

func (p *Player) Update() {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
		p.LocX -= SPEED
		p.facing = -1
		p.CurrentImg = p.runImages[(p.trackFrame-1)/16]
	case ebiten.IsKeyPressed(ebiten.KeyArrowRight):
		p.LocX += SPEED
		p.facing = 1
		p.CurrentImg = p.runImages[(p.trackFrame-1)/16]
	case ebiten.IsKeyPressed(ebiten.KeyArrowUp):
		p.LocY -= SPEED
		p.CurrentImg = p.runImages[(p.trackFrame-1)/16]
	case ebiten.IsKeyPressed(ebiten.KeyArrowDown):
		p.LocY += SPEED
		p.CurrentImg = p.runImages[(p.trackFrame-1)/16]
	case ebiten.IsKeyPressed(ebiten.KeyA):
		p.CurrentImg = p.castImages[(p.trackFrame-1)/16]
		if util.IsCDExceeded(2, p.lastCast) {

		} else {

		}
	default:
		p.CurrentImg = p.idleImages[(p.trackFrame-1)/16]
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	if p.trackFrame >= len(p.idleImages)*16 {
		p.trackFrame = 0
	} else {
		p.trackFrame += 1
	}
	drawOps := ebiten.DrawImageOptions{}
	distanceErr := float64(p.CurrentImg.Bounds().Dx() * p.facing)

	drawOps.GeoM.Scale(float64(p.facing*2), 2)
	drawOps.GeoM.Translate(p.LocX-distanceErr/1.3, p.LocY)
	screen.DrawImage(p.CurrentImg, &drawOps)

}
