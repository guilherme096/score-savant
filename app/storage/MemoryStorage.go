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
	return TempPlayers
}

var TempPlayers = map[string]*Player.Player{
	"1": Player.NewPlayer(
		"1",
		Player.NewPlayerBio(
			"1",
			"Lionel Messi",
			33,
			72.0,
			1.7,
			"Argentina",
			"Left",
			"https://cdn.scoresavant.com/players/lionel_messi.jpg",
		),
		Player.NewPlayerContract(
			"1",
			100000.0,
			"Barcelona",
			100000.0,
			5,
			"2025-06-30",
			100000.0,
		),
		[]*Player.Attribute{
			Player.NewAttribute("1", "Crossing", 20, Player.Technical),
			Player.NewAttribute("2", "Finishing", 20, Player.Technical),
			Player.NewAttribute("3", "Heading", 15, Player.Technical),
			Player.NewAttribute("4", "Short Passing", 20, Player.Technical),
			Player.NewAttribute("5", "Volleys", 18, Player.Technical),
			Player.NewAttribute("6", "Dribbling", 20, Player.Technical),
			Player.NewAttribute("7", "Curve", 18, Player.Technical),
		},
		[]*Player.Attribute{
			Player.NewAttribute("8", "Aggression", 10, Player.Mental),
			Player.NewAttribute("9", "Interceptions", 10, Player.Mental),
			Player.NewAttribute("10", "Positioning", 15, Player.Mental),
			Player.NewAttribute("11", "Vision", 20, Player.Mental),
		},
		[]*Player.Attribute{
			Player.NewAttribute("12", "Acceleration", 20, Player.Physical),
			Player.NewAttribute("13", "Stamina", 15, Player.Physical),
			Player.NewAttribute("14", "Strength", 10, Player.Physical),
			Player.NewAttribute("15", "Balance", 15, Player.Physical),
			Player.NewAttribute("16", "Sprint Speed", 20, Player.Physical),
		},
	),
}
