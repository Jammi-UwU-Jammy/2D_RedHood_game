package characters

import (
	"RedHood/util"
	"math"
	"math/rand"
	"time"
)

const (
	IDLE_STATE = iota
	WALK_STATE //walk state has to be below idle state
	RUN_STATE
	ATTACK_STATE

	MAGE_IDLE_IMAGES_URI = "enemies/mage/images/Idle.png"
	MAGE_RUN_IMAGES_URI  = "enemies/mage/images/Run.png"
	MAGE_CAST_IMAGES_URI = "enemies/mage/images/Attack1.png"

	SKE_IDLE_IMAGES_URI = "enemies/skeleton/images/idle.png"
	SKE_RUN_IMAGES_URI  = "enemies/skeleton/images/walk.png"
	SKE_CAST_IMAGES_URI = "enemies/skeleton/images/atk.png"
)

type Mob struct {
	*Character
	state int
}

func NewEnemyMage() *Mob {
	character := &Character{
		LocX: 176, LocY: 527,
		HP:       150,
		facing:   1,
		lastCast: time.Now(),
		Velocity: util.Vector{X: 0, Y: 0}}
	mob := Mob{Character: character}

	mob.idleImages = mob.loadImageAssets(MAGE_IDLE_IMAGES_URI, util.Point{X: 0, Y: 0}, 250, 250)
	mob.runImages = mob.loadImageAssets(MAGE_RUN_IMAGES_URI, util.Point{X: 0, Y: 0}, 250, 250)
	mob.atkImages = mob.loadImageAssets(MAGE_CAST_IMAGES_URI, util.Point{X: 0, Y: 0}, 250, 250)

	mob.CurrentImg = mob.idleImages[0]
	mob.maxFrame = len(mob.idleImages)
	return &mob
}

func NewEnemySkeleton() *Mob {
	character := &Character{
		LocX: 176, LocY: 527,
		HP:       50,
		facing:   1,
		lastCast: time.Now(),
		Velocity: util.Vector{X: 0, Y: 0}}
	mob := Mob{Character: character}

	mob.idleImages = mob.loadImageAssets(SKE_IDLE_IMAGES_URI, util.Point{X: 0, Y: 0}, 24, 32)
	mob.runImages = mob.loadImageAssets(SKE_RUN_IMAGES_URI, util.Point{X: 0, Y: 0}, 22, 33)
	mob.atkImages = mob.loadImageAssets(SKE_CAST_IMAGES_URI, util.Point{X: 0, Y: 0}, 43, 37)

	mob.CurrentImg = mob.idleImages[0]
	return &mob
}

func (m *Mob) Update(player *Player) int {
	m.CurrentImg = m.idleImages[(m.trackFrame-1)/IMG_PER_SEC]
	m.stateUpdate(player.LocX, player.LocY)

	outputDamage := 0

	switch m.state {
	case ATTACK_STATE:
		m.maxFrame = len(m.atkImages)
		if util.IsCDExceeded(1, m.lastCast) {
			m.lastCast = time.Now()
			outputDamage = 1
		}
		m.CurrentImg = m.atkImages[(m.trackFrame-1)/IMG_PER_SEC]
	case WALK_STATE:
		x, y := util.UnitVector(m.Velocity.X, m.Velocity.Y)
		m.LocX += x
		m.LocY += y
		m.Velocity.X *= 0.95
		m.Velocity.Y *= 0.95
		m.facing = int(m.Velocity.X / math.Abs(m.Velocity.X))
		m.maxFrame = len(m.runImages)
		m.CurrentImg = m.runImages[(m.trackFrame-1)/IMG_PER_SEC]
	case RUN_STATE:
		x, y := util.UnitVectorFromTwoPoints(m.LocX, m.LocY, player.LocX, player.LocY)
		m.LocX += x
		m.LocY += y
		m.facing = int(x / math.Abs(x))
		m.maxFrame = len(m.runImages)
		m.CurrentImg = m.runImages[(m.trackFrame-1)/IMG_PER_SEC]
	default:
		if util.VectorLength(m.Velocity.X, m.Velocity.Y) < 0.01 && rand.Intn(1000) < 3 {
			m.Velocity.X, m.Velocity.Y = rand.Float64()*5-2, rand.Float64()*5-2
		}
		m.maxFrame = len(m.idleImages)
		m.CurrentImg = m.idleImages[(m.trackFrame-1)/IMG_PER_SEC]
	}
	return outputDamage
}

func (m *Mob) stateUpdate(charLocX, charLocY float64) {
	distance := util.VectorsDistance(m.LocX, m.LocY, charLocX, charLocY)
	if distance < 100 && distance > 20 {
		m.state = RUN_STATE
	} else if distance <= 20 {
		m.state = ATTACK_STATE
	} else { // too far away from player
		if util.VectorLength(m.Velocity.X, m.Velocity.Y) < 0.01 {
			m.state = IDLE_STATE
		} else {
			m.state = WALK_STATE
		}
	}
}
