package characters

import (
	"RedHood/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"time"
)

const (
	SPEED           = 4
	IDLE_IMAGES_URI = "player/images/idle-sheet.png"
	RUN_IMAGES_URI  = "player/images/run-sheet.png"
	CAST_IMAGES_URI = "player/images/atk-sheet.png"
)

func NewPlayer() *Player {
	character := &Character{
		LocX: 800, LocY: 500,
		facing:   1,
		lastCast: time.Now(),
		Velocity: util.Vector{X: 0, Y: 0}}
	player := Player{Character: character}

	player.idleImages = player.loadImageAssets(IDLE_IMAGES_URI, util.Point{X: 0, Y: 0}, 80, 80)
	player.runImages = player.loadImageAssets(RUN_IMAGES_URI, util.Point{X: 0, Y: 0}, 80, 80)
	player.castImages = player.loadImageAssets(CAST_IMAGES_URI, util.Point{X: 0, Y: 0}, 80, 80)

	player.CurrentImg = player.idleImages[0]
	player.LocY -= float64(player.CurrentImg.Bounds().Dy())
	return &player
}

type Player struct {
	*Character

	castImages []*ebiten.Image
}

func (p *Player) Update(obstacles []*tiled.Object) {
	oldX, oldY := p.LocX, p.LocY

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
	p.Velocity.X, p.Velocity.Y = p.LocX-oldX, p.LocY-oldY
	if p.collisionVSBG(obstacles) {
		p.LocX -= p.Velocity.X
		p.LocY -= p.Velocity.Y
		//fmt.Println("Collided")
	}
}
