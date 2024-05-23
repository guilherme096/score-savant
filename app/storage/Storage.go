package storage

type IStorage interface {
	Start()
	Stop()
	LoadPlayerById(id string) (map[string]interface{}, []map[string]interface{}, error)
	GetAttributeList(att_type string) []string
}
