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
	GetLeagueList(page int, amount int, filters map[string]interface{}) ([]map[string]interface{}, error)
	GetLeagueById(id int) (map[string]interface{}, error)
	AddPlayer(name string, age int, weight int, height int, nation string, nation_league_id int, league string, club string, foot string, value int, position string, role string, wage float64, contract_end string, release_clause int, atts []string, url string)
	DeletePlayer(id int)
	GetNationList(page int, amount int, filters map[string]interface{}) ([]map[string]interface{}, error)
	GetNationById(id int) (map[string]interface{}, error)
	StarPlayer(id int)
	GetStaredPlayers(pageNumber int) ([]map[string]interface{}, error)
}
