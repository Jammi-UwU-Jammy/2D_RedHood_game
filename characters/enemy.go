package characters

import (
	"RedHood/util"
	"time"
)

const (
	E_IDLE_IMAGES_URI = "enemies/mage/images/Idle.png"
	E_RUN_IMAGES_URI  = "enemies/mage/images/Run.png"
	E_CAST_IMAGES_URI = "enemies/mage/images/Attack1.png"
)

type Mob struct {
	*Character
}

func NewEnemy() *Mob {
	character := &Character{facing: 1, lastCast: time.Now(), Velocity: util.Vector{X: 0, Y: 0}}
	mob := Mob{Character: character}

	mob.idleImages = mob.loadImageAssets(E_IDLE_IMAGES_URI, util.Point{X: 0, Y: 0}, 250, 250)
	mob.runImages = mob.loadImageAssets(E_RUN_IMAGES_URI, util.Point{X: 0, Y: 0}, 250, 250)
	mob.castImages = mob.loadImageAssets(E_CAST_IMAGES_URI, util.Point{X: 0, Y: 0}, 250, 250)

	mob.CurrentImg = mob.idleImages[0]
	return &mob
}

func (m *Mob) Update() {
	m.CurrentImg = m.idleImages[(m.trackFrame-1)/60]
}
