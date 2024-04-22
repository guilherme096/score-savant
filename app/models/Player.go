package models

type Player struct {
	Id                  string
	PlayerBio           *PlayerBio
	Contract            *PlayerContract
	TechnicalAttributes []*Attribute
	PhysicalAttributes  []*Attribute
	MentalAttributes    []*Attribute
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

func NewPlayer(Id string, PlayerBio *PlayerBio, PlayerContract *PlayerContract, TechnicalAttributes []*Attribute, MentalAttributes []*Attribute, PhysicalAttributes []*Attribute) *Player {
	return &Player{
		Id:                  Id,
		PlayerBio:           PlayerBio,
		Contract:            PlayerContract,
		TechnicalAttributes: TechnicalAttributes,
		MentalAttributes:    MentalAttributes,
		PhysicalAttributes:  PhysicalAttributes,
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

func (p *Player) AddTechnicalAttribute(attribute *Attribute) {
	p.TechnicalAttributes = append(p.TechnicalAttributes, attribute)
}

func (p *Player) AddPhysicalAttribute(attribute *Attribute) {
	p.PhysicalAttributes = append(p.PhysicalAttributes, attribute)
}

func (p *Player) AddMentalAttribute(attribute *Attribute) {
	p.MentalAttributes = append(p.MentalAttributes, attribute)
}
