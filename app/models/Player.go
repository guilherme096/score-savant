package models

type Player struct {
	Id        string
	PlayerBio *PlayerBio
	Contract  *PlayerContract
}

type PlayerBio struct {
	Id     string
	Name   string
	Age    int
	Weight float64
	Height float64
	Nation string
	Foot   string
	Photo  string
}

type PlayerContract struct {
	Id               string
	Wage             float64
	CurrentClub      string
	Value            float64
	ContractDuration int
	ContractEnd      string
	ReleaseClause    float64
}

func NewPlayer(Id string, PlayerBio *PlayerBio, PlayerContract *PlayerContract) *Player {
	return &Player{
		Id:        Id,
		PlayerBio: PlayerBio,
		Contract:  PlayerContract,
	}
}

func NewPlayerContract(id string, wage float64, currentClub string, value float64, contractDuration int, contractEnd string, releaseClause float64) *PlayerContract {
	return &PlayerContract{
		Id:               id,
		Wage:             wage,
		CurrentClub:      currentClub,
		Value:            value,
		ContractDuration: contractDuration,
		ContractEnd:      contractEnd,
		ReleaseClause:    releaseClause,
	}
}

func NewPlayerBio(id string, name string, age int, weight float64, height float64, nation string, foot string, photo string) *PlayerBio {
	return &PlayerBio{
		Id:     id,
		Name:   name,
		Age:    age,
		Weight: weight,
		Height: height,
		Nation: nation,
		Foot:   foot,
		Photo:  photo,
	}
}

func (p *Player) SetContract(contract *PlayerContract) {
	p.Contract = contract
}

func (p *Player) SetBio(bio *PlayerBio) {
	p.PlayerBio = bio
}
