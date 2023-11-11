package characters

import (
	"RedHood/util"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

const (
	SPEED           = 5
	IDLE_IMAGES     = 0
	IDLE_IMAGES_URI = "player/idles/idle-sheet.png"
)

type Character struct {
	CurrentImg *ebiten.Image
	idleImages []*ebiten.Image
	LocX       float64
	LocY       float64
	facing     int
	trackFrame int
}

func NewPlayer() *Player {
	character := &Character{}
	player := Player{Character: character}

	//player.idleImages = make([]*ebiten.Image, 0, 60)
	player.loadImageAssets(IDLE_IMAGES_URI,
		util.Point{X: 0, Y: 0}, 80, 80)
	return &player
}

type Player struct {
	*Character

	runImages  []*ebiten.Image
	castImages []*ebiten.Image
}

func (p *Player) loadImageAssets(uri string, offset util.Point, width, height int) {
	img := util.LoadEmbeddedImage(uri, 301)

	for i := 0; i < 16; i++ {
		subImg := img.SubImage(image.Rect(offset.X, offset.Y, offset.X+width, height))
		subImage := ebiten.NewImageFromImage(subImg)
		//pool = append(pool, subImage)
		p.idleImages = append(p.idleImages, subImage)
		offset.X += width
	}
}

func (p *Player) Update() {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
		p.LocX -= SPEED
	case ebiten.IsKeyPressed(ebiten.KeyArrowRight):
		p.LocX += SPEED
	case ebiten.IsKeyPressed(ebiten.KeyArrowUp):
		p.LocY -= SPEED
	case ebiten.IsKeyPressed(ebiten.KeyArrowDown):
		p.LocY += SPEED
	case ebiten.IsKeyPressed(ebiten.KeyA):

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

	drawOps.GeoM.Reset()
	drawOps.GeoM.Translate(p.LocX, p.LocY)
	screen.DrawImage(p.CurrentImg, &drawOps)

}
