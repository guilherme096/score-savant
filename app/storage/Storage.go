package storage

type IStorage interface {
	Start()
	Stop()
	LoadPlayerById(id string) (map[string]interface{}, []map[string]interface{}, error)
	GetAttributeList(att_type string) []string
	GetPlayerPosition(id string) (int, string, error)
	GetRolesByPositionId(PositonId int) []map[string]interface{}
	GetKeyAttributeList(role_id int) []string
	GetRoleByPlayerId(player_id int) (string, error)
	GetPlayerList(page int, amount int, filters map[string]interface{}) ([]map[string]interface{}, error)
	GetClubList(page int, amount int, filters map[string]interface{}) ([]map[string]interface{}, error)
	GetRandomPlayer() (name string, nation string, club string, url string, playerId int, nationId int, clubId int, err error)
	GetClubById(id int) (map[string]interface{}, error)
}
