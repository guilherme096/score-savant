package storage

type IStorage interface {
	Start()
	Stop()
	LoadPlayerById(id string) (map[string]interface{}, error)
}
