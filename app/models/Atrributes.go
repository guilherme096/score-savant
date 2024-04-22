package models

type AttType string

const (
	Physical    AttType = "Physical"
	Mental      AttType = "Mental"
	Technical   AttType = "Technical"
	Goalkeeping AttType = "Goalkeeping"
)

type Attribute struct {
	Id    string
	Name  string
	Value int
	Type  AttType
}

func NewAttribute(id string, name string, value int, attType AttType) *Attribute {
	if value < 0 || value > 20 {
		return nil
	}
	return &Attribute{
		Id:    id,
		Name:  name,
		Value: value,
		Type:  attType,
	}
}
