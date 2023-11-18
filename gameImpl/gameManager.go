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

type Manager struct {
	players        []*characters.Player
	maps           []*environments.Map
	universalItems []*etc.Item

	//TODO: Below is for local running/testing
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
}

func (m *Manager) spawnPlayer() {
	player := characters.NewPlayer()
	m.players = append(m.players, player)
	var mobs []*characters.Mob
	for i := 0; i < 5; i++ {
		mob1 := characters.NewEnemyMage()
		mob2 := characters.NewEnemySkeleton()
		mobs = append(mobs, mob1, mob2)
	}
	m.currentPlayer = player
	m.currentMap = environments.NewDefaultBackground(mobs)
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
	}
	time.Sleep(17 * time.Millisecond)

}

func (m *Manager) getPlayerData() (locX, locY float64) {
	locX, _ = m.CurrentGame.PlayerData["LocX"].(float64)
	locY, _ = m.CurrentGame.PlayerData["LocY"].(float64)
	_, exists := m.CurrentGame.PlayerData["Damage"]
	if exists {
		for _, mob := range m.currentMap.Enemies {
			dmg, _ := m.CurrentGame.PlayerData["Damage"].(int)
			if util.VectorsDistance(mob.LocX, mob.LocY, locX, locY) < 20 {
				mob.HP -= dmg
				fmt.Println(mob.HP)
				break //only one mob
			}
		}
	}
	return locX, locY
}

func (m *Manager) updateEnemies() {
	var mobs []*characters.Mob
	for _, mob := range m.currentMap.Enemies {
		if mob.HP > 0 {
			mobs = append(mobs, mob)
			damageFromMob := mob.Update(m.currentPlayer)
			m.currentPlayer.HP -= damageFromMob
		} else {
			drop := etc.NewRandomSword(mob.LocX, mob.LocY)
			m.universalItems = append(m.universalItems, drop)
			//fmt.Printf("Item: X %.2f | Y %.2f \n", drop.LocX, drop.LocY)
		}
	}
	m.currentMap.Enemies = mobs
	m.CurrentGame.universalItems = m.universalItems
}
