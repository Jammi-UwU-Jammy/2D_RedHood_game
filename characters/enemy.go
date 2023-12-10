package characters

import (
	"RedHood/util"
	"fmt"
	"github.com/solarlune/paths"
	"math"
	"math/rand"
	"time"
)

const (
	IDLE_STATE = iota
	WALK_STATE //walk state has to be below idle state
	RUN_STATE
	ATTACK_STATE
	EXIT_STATE

	MAGE_IDLE_IMAGES_URI = "enemies/mage/images/Idle.png"
	MAGE_RUN_IMAGES_URI  = "enemies/mage/images/Run.png"
	MAGE_CAST_IMAGES_URI = "enemies/mage/images/Attack1.png"
	MAGE_EXIT_IMAGES_URI = "enemies/mage/images/Death.png"

	SKE_IDLE_IMAGES_URI = "enemies/skeleton/images/idle.png"
	SKE_RUN_IMAGES_URI  = "enemies/skeleton/images/walk.png"
	SKE_CAST_IMAGES_URI = "enemies/skeleton/images/atk.png"
	SKE_EXIT_IMAGES_URI = "enemies/skeleton/images/Skeleton Dead.png"

	SMR_IDLE_IMAGES_URI = "enemies/samurai/IDLE.png"
	SMR_RUN_IMAGES_URI  = "enemies/samurai/RUN.png"
	SMR_ATK_IMAGES_URI  = "enemies/samurai/BASIC ATTACK.png"
	SMR_EXIT_IMAGES_URI = "NON-AVAILABLE"
)

type Mob struct {
	*Character
	state    int
	autoPath *paths.Path
	pathGrid *paths.Grid

	prevState int
}

func NewEnemyMage(path *paths.Grid) *Mob {
	character := &Character{
		LocX: 800 + rand.Float64()*400 - 200, LocY: 450 + rand.Float64()*400 - 200,
		HP:       150,
		Facing:   1,
		lastCast: time.Now(),
		Velocity: util.Vector{X: 0, Y: 0}}
	mob := Mob{Character: character, pathGrid: path}

	mob.idleImages = mob.loadImageAssets(MAGE_IDLE_IMAGES_URI, util.Point{X: 0, Y: 0}, 250, 250)
	mob.runImages = mob.loadImageAssets(MAGE_RUN_IMAGES_URI, util.Point{X: 0, Y: 0}, 250, 250)
	mob.atkImages = mob.loadImageAssets(MAGE_CAST_IMAGES_URI, util.Point{X: 0, Y: 0}, 250, 250)
	mob.ExitImages = mob.loadImageAssets(MAGE_EXIT_IMAGES_URI, util.Point{X: 0, Y: 0}, 250, 250)

	mob.CurrentImg = mob.idleImages[0]
	mob.maxFrame = len(mob.idleImages)
	return &mob
}

func NewEnemySkeleton(path *paths.Grid) *Mob {
	character := &Character{
		LocX: 800 + rand.Float64()*400 - 200, LocY: 450 + rand.Float64()*400 - 200,
		HP:       50,
		Facing:   1,
		lastCast: time.Now(),
		Velocity: util.Vector{X: 0, Y: 0}}
	mob := Mob{Character: character, pathGrid: path}

	mob.idleImages = mob.loadImageAssets(SKE_IDLE_IMAGES_URI, util.Point{X: 0, Y: 0}, 24, 32)
	mob.runImages = mob.loadImageAssets(SKE_RUN_IMAGES_URI, util.Point{X: 0, Y: 0}, 22, 33)
	mob.atkImages = mob.loadImageAssets(SKE_CAST_IMAGES_URI, util.Point{X: 0, Y: 0}, 43, 37)
	mob.ExitImages = mob.loadImageAssets(SKE_EXIT_IMAGES_URI, util.Point{X: 0, Y: 0}, 33, 32)

	mob.CurrentImg = mob.idleImages[0]
	return &mob
}

func NewSamurai(path *paths.Grid) *Mob {
	character := &Character{
		LocX: 1000 + rand.Float64()*400 - 200, LocY: 200 + rand.Float64()*400 - 200,
		HP:       100,
		Facing:   1,
		lastCast: time.Now(),
		Velocity: util.Vector{X: 0, Y: 0}}
	mob := Mob{Character: character, pathGrid: path}

	mob.idleImages = mob.loadImageAssets(SMR_IDLE_IMAGES_URI, util.Point{X: 0, Y: 0}, 158, 125)
	mob.runImages = mob.loadImageAssets(SMR_RUN_IMAGES_URI, util.Point{X: 0, Y: 0}, 158, 125)
	mob.atkImages = mob.loadImageAssets(SMR_ATK_IMAGES_URI, util.Point{X: 0, Y: 0}, 158, 125)
	mob.ExitImages = mob.idleImages
	mob.CurrentImg = mob.idleImages[0]
	return &mob
}

func (m *Mob) Update(player *Player) int {
	m.CurrentImg = m.idleImages[(m.trackFrame-1)/IMG_PER_SEC]
	m.stateUpdate(player)

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
		m.Facing = int(m.Velocity.X / math.Abs(m.Velocity.X))
		m.maxFrame = len(m.runImages)
		m.CurrentImg = m.runImages[(m.trackFrame-1)/IMG_PER_SEC]
	case RUN_STATE:
		x, _ := util.UnitVectorFromTwoPoints(m.LocX, m.LocY, player.LocX, player.LocY)
		//m.LocX += x
		//m.LocY += y
		m.Facing = int(x / math.Abs(x))
		m.updatePath(player)
		m.maxFrame = len(m.runImages)
		m.CurrentImg = m.runImages[(m.trackFrame-1)/IMG_PER_SEC]
	case EXIT_STATE:
		m.maxFrame = len(m.runImages)
		m.CurrentImg = m.ExitImages[(m.trackFrame-1)/IMG_PER_SEC]
	default:
		if util.VectorLength(m.Velocity.X, m.Velocity.Y) < 0.01 && rand.Intn(1000) < 3 {
			m.Velocity.X, m.Velocity.Y = rand.Float64()*5-2, rand.Float64()*5-2
		}
		m.maxFrame = len(m.idleImages)
		m.CurrentImg = m.idleImages[(m.trackFrame-1)/IMG_PER_SEC]
	}
	return outputDamage
}

func (m *Mob) stateUpdate(player *Player) {
	distance := util.VectorsDistance(m.LocX, m.LocY, player.LocX, player.LocY)
	m.prevState = m.state

	if m.HP <= 0 {
		m.state = EXIT_STATE
		return
	}
	if distance < 300 && distance > 20 {
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

func (m *Mob) updatePath(player *Player) {
	if m.prevState != m.state && m.state != EXIT_STATE {
		m.makePath(player)
	}

	if m.autoPath != nil {
		//if len(m.autoPath.Cells) == 0 {
		//	return
		//}
		pathCell := m.autoPath.Current()
		if math.Abs(float64(pathCell.X*util.TILEWIDTH)-(m.LocX)) <= 1 &&
			math.Abs(float64(pathCell.Y*util.TILEHEIGHT)-(m.LocY)) <= 1 { //if we are now on the tile we need to be on
			m.autoPath.Advance()
			fmt.Println("Advanced")
		}
		direction := 0.0
		if pathCell.X*util.TILEWIDTH > int(m.LocX) {
			direction = 1
		} else if pathCell.X*util.TILEWIDTH < int(m.LocX) {
			direction = -1
		}
		fmt.Println("Cell ", pathCell.X*32, " : ", pathCell.Y*32)
		fmt.Println("Mob ", int(m.LocX), ":", int(m.LocY))

		directionY := 0.0
		if pathCell.Y*util.TILEHEIGHT > int(m.LocY) {
			directionY = 1
		} else if pathCell.Y*util.TILEHEIGHT < int(m.LocY) {
			directionY = -1
		}
		m.LocX += direction * 2
		m.LocY += directionY * 2

	} else {
		fmt.Println("NIL")
	}
}

func (m *Mob) makePath(player *Player) {
	startRow := int(m.LocY) / util.TILEHEIGHT
	startCol := int(m.LocX) / util.TILEWIDTH
	startCell := m.pathGrid.Get(startCol, startRow)
	endCell := m.pathGrid.Get(int(player.LocX)/util.TILEHEIGHT, int(player.LocY)/util.TILEWIDTH)
	//fmt.Println("Mob: ", startCell, " -- ", endCell)
	m.autoPath = m.pathGrid.GetPathFromCells(startCell, endCell, false, false)
}
