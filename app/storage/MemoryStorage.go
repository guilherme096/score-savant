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
		"1": Player.NewPlayer("1", Player.NewPlayerBio("1", "Lionel Messi", 33, 72.0, 1.7, "Argentina", "Left", "https://cdn.scoresavant.com/players/lionel_messi.jpg"), Player.NewPlayerContract("1", 1000000.0, "Barcelona", 100000000.0, 5, "2025-06-30", 100000000.0)),
		"2": Player.NewPlayer("2", Player.NewPlayerBio("2", "Cristiano Ronaldo", 35, 80.0, 1.85, "Portugal", "Right", "https://cdn.scoresavant.com/players/cristiano_ronaldo.jpg"), Player.NewPlayerContract("2", 1000000.0, "Juventus", 100000000.0, 5, "2022-06-30", 100000000.0)),
	}
}
