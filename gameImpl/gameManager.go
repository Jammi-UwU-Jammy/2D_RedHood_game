package gameImpl

import (
	"RedHood/characters"
	"RedHood/environments"
	"RedHood/etc"
	"RedHood/util"
	"fmt"
	"sync"
	"time"
)

const (
	MAX_BAG_SLOTS = 6
)

type Manager struct {
	players        []*characters.Player
	maps           []*environments.Map
	universalItems []*etc.Item

	//TODO: Below is for running/testing
	currentPlayer *characters.Player
	CurrentGame   *Game
	currentMap    *environments.Map

	gameMu     sync.RWMutex
	stopSignal chan struct{}
}

func (m *Manager) PopulateResources() {
	m.gameMu.Lock()
	defer m.gameMu.Unlock()

	m.spawnPlayer()
	m.CurrentGame = NewGame(m.currentPlayer, m.currentMap)
	//m.CurrentGame.PopulateQuests()
}

func (m *Manager) spawnPlayer() {
	player := characters.NewPlayer()
	m.players = append(m.players, player)
	var mobs1, mobs2 []*characters.Mob
	for i := 0; i < 5; i++ {
		mob1 := characters.NewEnemyMage()
		mob2 := characters.NewEnemySkeleton()
		mobs1 = append(mobs1, mob1, mob2)
	}
	for i := 0; i < 8; i++ {
		mob := characters.NewSamurai()
		mobs2 = append(mobs2, mob)
	}
	m.currentPlayer = player
	m.currentMap = environments.NewDefaultMap(mobs1)

	m.maps = append(m.maps, m.currentMap)
	m.maps = append(m.maps, environments.NewLakeMap(mobs2))
}

func (m *Manager) Start() {
	for {
		select {
		case <-m.stopSignal:
			return
		default:
			m.updateGame()
		}
	}
}

func (m *Manager) Stop() {
	close(m.stopSignal)
}

func (m *Manager) updateGame() {
	//TODO: check collisions and stuff then send back to players to render
	m.gameMu.RLock()
	defer m.gameMu.RUnlock()

	if m.CurrentGame != nil {
		m.getPlayerData()
		m.updateEnemies()
		m.updateEffects()
	}
	time.Sleep(17 * time.Millisecond)

}

func (m *Manager) getPlayerData() (locX, locY float64) {
	locX, _ = m.CurrentGame.PlayerDataToSend["LocX"].(float64)
	locY, _ = m.CurrentGame.PlayerDataToSend["LocY"].(float64)
	m.handlePlayerDamage(locX, locY)
	m.handleLoot(locX, locY)
	m.handlePortal()

	return locX, locY
}

func (m *Manager) updateEnemies() {
	var mobs []*characters.Mob
	for _, mob := range m.currentMap.Enemies {
		if mob.CollisionVSBG(m.currentMap.TiledMap) {
			mob.Velocity.X *= -1
			mob.Velocity.Y *= -1
		}
		if mob.HP > 0 {
			mobs = append(mobs, mob)
			damageFromMob := mob.Update(m.currentPlayer)
			m.currentPlayer.HP -= damageFromMob
		} else {
			drop := etc.NewRandomSword(mob.LocX, mob.LocY)
			m.universalItems = append(m.universalItems, drop)
			mobExit := etc.NewLastingEffect(mob.ExitImages, mob.LocX, mob.LocY, mob.Facing)
			m.CurrentGame.effects = append(m.CurrentGame.effects, mobExit)
			//fmt.Printf("Item: X %.2f | Y %.2f \n", drop.LocX, drop.LocY)
		}
	}
	m.currentMap.Enemies = mobs
	m.CurrentGame.universalItems = m.universalItems
}

func (m *Manager) updateEffects() {
	var tempEffects []*etc.Effect
	for _, e := range m.CurrentGame.effects {
		if !e.IsDone {
			e.Update()
			tempEffects = append(tempEffects, e)
		}
	}
	m.CurrentGame.effects = tempEffects
}

func (m *Manager) UpdateMap(mapID int) {
	m.currentMap = m.maps[mapID]
	m.CurrentGame.background = m.currentMap
	//fmt.Println(len(m.currentMap.Enemies))
	m.CurrentGame.portals = m.CurrentGame.background.TiledMap.ObjectGroups[0].Objects
}

func (m *Manager) handlePortal() {
	_, exists := m.CurrentGame.PlayerDataToSend["Map"]
	if exists {
		id, _ := m.CurrentGame.PlayerDataToSend["Map"].(uint32)
		fmt.Println(id)
		if id == 60 {
			m.UpdateMap(0)
			m.currentPlayer.LocX = m.CurrentGame.portals[0].X - 50
			m.currentPlayer.LocY = m.CurrentGame.portals[0].Y
			util.PlaySound(m.CurrentGame.portalSound)
		} else if id == 135 {
			m.UpdateMap(1)
			m.currentPlayer.LocX = m.CurrentGame.portals[0].X + 100
			m.currentPlayer.LocY = m.CurrentGame.portals[0].Y
			util.PlaySound(m.CurrentGame.portalSound)
		} else if id == 136 {
			if m.currentPlayer.HP < m.currentPlayer.MaxStat.HP {
				m.currentPlayer.HP += 1
				util.PlaySound(m.CurrentGame.healSound)
			}
		}
	}
}

func (m *Manager) handleLoot(locX, locY float64) {
	_, exists := m.CurrentGame.PlayerDataToSend["Loot"]
	if exists {
		slots, _ := m.CurrentGame.PlayerDataToSend["Loot"].(int)
		var tempUniversalItems []*etc.Item
		for _, item := range m.universalItems {
			if util.VectorsDistance(locX, locY, item.LocX, item.LocY) < 60 &&
				slots < MAX_BAG_SLOTS-1 {
				m.CurrentGame.player.Bag = append(m.CurrentGame.player.Bag, item)
				fmt.Println("Added.")
			} else {
				tempUniversalItems = append(tempUniversalItems, item)
			}
		}
		m.universalItems = tempUniversalItems
	}
}

func (m *Manager) handlePlayerDamage(locX, locY float64) {
	_, exists := m.CurrentGame.PlayerDataToSend["Damage"]
	if exists {
		dmg, _ := m.CurrentGame.PlayerDataToSend["Damage"].(int)
		for _, mob := range m.currentMap.Enemies {
			if util.VectorsDistance(mob.LocX, mob.LocY, locX, locY) < 20 {
				mob.HP -= dmg
				fmt.Println(mob.HP)
				break //only one mob
			}
		}
	}
}
