package models

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
type Player struct {
	Id     string
	Name   string
	Age    int
	Weight float64
	Height float64
	Nation string
	Club   string
	Foot   string
	Value  float64
	Photo  string
}

func NewPlayer(id string, name string, age int, weight float64, height float64, nation string, club string, foot string, value float64, photo string) *Player {
	return &Player{
		Id:     id,
		Name:   name,
		Age:    age,
		Weight: weight,
		Height: height,
		Nation: nation,
		Club:   club,
		Foot:   foot,
		Value:  value,
		Photo:  photo,
	}
}

func (p *Player) Bio() *PlayerBio {
	return &PlayerBio{
		Id:     p.Id,
		Name:   p.Name,
		Age:    p.Age,
		Weight: p.Weight,
		Height: p.Height,
		Nation: p.Nation,
		Foot:   p.Foot,
		Photo:  p.Photo,
	}
}
