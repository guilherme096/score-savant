package storage

import (
	Player "guilherme096/score-savant/models"
)

type IStorage interface {
	LoadPlayerById(id string) (*Player.Player, error)
}
