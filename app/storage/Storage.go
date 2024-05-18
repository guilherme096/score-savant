package storage

import (
	Player "guilherme096/score-savant/models"
)

type IStorage interface {
	Start()
	LoadPlayerById(id string) (*Player.Player, error)
}
