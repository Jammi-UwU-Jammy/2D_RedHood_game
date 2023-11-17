package gameImpl

import (
	"RedHood/characters"
	"RedHood/environments"
)

type Manager struct {
	players []*characters.Player

	maps []*environments.Map
}

func (m *Manager) spawnPlayer() {
	player := characters.NewPlayer()
	m.players = append(m.players, player)
}

func (m *Manager) updateGame() {
	//TODO: check collisions and stuff then send back to players to render
}
