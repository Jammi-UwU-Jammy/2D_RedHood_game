package gameImpl

import (
	"RedHood/characters"
	"RedHood/environments"
	"fmt"
	"sync"
	"time"
)

type Manager struct {
	players []*characters.Player
	maps    []*environments.Map

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
		fmt.Println("Updating")
		m.updateEnemies()
	}
	time.Sleep(17 * time.Millisecond)

}

func (m *Manager) updateEnemies() {
	for _, mob := range m.currentMap.Enemies {
		mob.Update(m.currentPlayer)
	}
}
