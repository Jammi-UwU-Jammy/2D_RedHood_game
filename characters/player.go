package characters

import (
	"RedHood/etc"
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
		LocX: 32, LocY: 500,
		MaxStat:  &MaxStat{ATK: 1, HP: 100},
		HP:       100,
		facing:   1,
		lastCast: time.Now(),
		Velocity: util.Vector{X: 0, Y: 0}}
	player := Player{Character: character, WeaponE: nil}

	player.idleImages = player.loadImageAssets(IDLE_IMAGES_URI, util.Point{X: 0, Y: 0}, 80, 80)
	player.runImages = player.loadImageAssets(RUN_IMAGES_URI, util.Point{X: 0, Y: 0}, 80, 80)
	player.atkImages = player.loadImageAssets(CAST_IMAGES_URI, util.Point{X: 0, Y: 0}, 80, 80)

	player.CurrentImg = player.idleImages[0]
	player.maxFrame = len(player.idleImages)
	player.LocY -= float64(player.CurrentImg.Bounds().Dy())
	return &player
}

type Player struct {
	*Character
	Bag      []*etc.Item
	WeaponE  *etc.Item
	equipped []*etc.Item
}

func (p *Player) Update(blockedTiles *tiled.Map, obj []*tiled.Object) map[string]interface{} {

	outputs := make(map[string]interface{})
	oldX, oldY := p.LocX, p.LocY
	p.StatUpdate()

	switch {
	case ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
		p.LocX -= SPEED
		p.facing = -1
		p.maxFrame = len(p.runImages)
		p.CurrentImg = p.runImages[(p.trackFrame-1)/IMG_PER_SEC]
	case ebiten.IsKeyPressed(ebiten.KeyArrowRight):
		p.LocX += SPEED
		p.facing = 1
		p.maxFrame = len(p.runImages)
		p.CurrentImg = p.runImages[(p.trackFrame-1)/IMG_PER_SEC]
	case ebiten.IsKeyPressed(ebiten.KeyArrowUp):
		p.LocY -= SPEED
		p.maxFrame = len(p.runImages)
		p.CurrentImg = p.runImages[(p.trackFrame-1)/IMG_PER_SEC]
	case ebiten.IsKeyPressed(ebiten.KeyArrowDown):
		p.LocY += SPEED
		p.maxFrame = len(p.runImages)
		p.CurrentImg = p.runImages[(p.trackFrame-1)/IMG_PER_SEC]
	case ebiten.IsKeyPressed(ebiten.KeySpace):
		outputs["Loot"] = len(p.Bag)
	case ebiten.IsKeyPressed(ebiten.KeyA):
		p.maxFrame = len(p.atkImages)
		if util.IsCDExceeded(0.2, p.lastCast) {
			p.lastCast = time.Now()
			outputs["Damage"] = 10
		}
		p.CurrentImg = p.atkImages[(p.trackFrame-1)/IMG_PER_SEC]
	default:
		p.maxFrame = len(p.idleImages)
		p.CurrentImg = p.idleImages[(p.trackFrame-1)/IMG_PER_SEC]
	}
	p.Velocity.X, p.Velocity.Y = p.LocX-oldX, p.LocY-oldY
	if p.CollisionVSBG(blockedTiles) {
		p.LocX -= p.Velocity.X
		p.LocY -= p.Velocity.Y
		//fmt.Println("Collided")
	}
	is, id := p.CollisionVSObjects(obj)
	if is {
		outputs["Map"] = id
	}
	outputs["LocX"], outputs["LocY"] = p.LocX, p.LocY

	return outputs
}

func (p *Player) EquipItem(item *etc.Item) {
	p.WeaponE = item
}

func (p *Player) StatUpdate() {
	if p.WeaponE != nil {
		p.MaxStat.HP = 100 + p.WeaponE.Buffs.HP
		p.MaxStat.ATK = 1 + p.WeaponE.Buffs.ATK
		p.MaxStat.DEF = 100 + p.WeaponE.Buffs.DEF
	}
}
