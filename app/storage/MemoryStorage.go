package storage

import (
	Player "guilherme096/score-savant/models"
)

type MemoryStorage struct {
	Players map[string]*Player.Player
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		Players: temp_players(),
	}
}

func (m *MemoryStorage) LoadPlayerById(id string) (*Player.Player, error) {
	player, ok := m.Players[id]
	if !ok {
		return nil, nil
	}
	return player, nil
}

func temp_players() map[string]*Player.Player {
	return map[string]*Player.Player{
		"1": Player.NewPlayer("1", "Cristiano Ronaldo", 35, 80.0, 1.87, "Portugal", "Juventus", "Right", 100.0, "https://cdn.cnn.com/cnnnext/dam/assets/200805115818-ronaldo-juventus-file-super-tease.jpg"),
		"2": Player.NewPlayer("2", "Lionel Messi", 33, 72.0, 1.70, "Argentina", "Barcelona", "Left", 100.0, "https://www.fcbarcelona.com/photo-resources/2020/08/30/6c9b5e9b-9b2a-4b5e-8c4d-1e5e1c7e3f2d/mini_Messi-1-.jpg?width=1200&height=750"),
		"3": Player.NewPlayer("3", "Neymar Jr", 28, 68.0, 1.75, "Brazil", "Paris Saint-Germain", "Right", 100.0, "https://www.fcbarcelona.com/photo-resources/2020/08/30/6c9b5e9b-9b2a-4b5e-8c4d-1e5e1c7e3f2d/mini_Messi-1-.jpg?width=1200&height=750"),
	}
}
