package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/microsoft/go-mssqldb"
	Player "guilherme096/score-savant/models"
)

type MSqlStorage struct {
	connectionString string
	db               *sql.DB
}

func NewMSqlStorage(username string, password string, host string, port int, databaseName string) *MSqlStorage {

	conString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", host, username, password, port, databaseName)

	return &MSqlStorage{connectionString: conString}
}

func (m *MSqlStorage) Start() {
	fmt.Printf("Connecting to SQL Server: %s\n", m.connectionString)
	db, err := sql.Open("sqlserver", m.connectionString)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to SQL Server")
	m.db = db
	m.LoadPlayerById("1")
}

func (m *MSqlStorage) LoadPlayerById(id string) (*Player.Player, error) {
	rows, err := m.db.Query("SELECT * FROM Player WHERE id=?", 1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var name string
		var age int
		var weight float64
		var height float64
		if err := rows.Scan(&id, &name, &age, &weight, &height); err != nil {
			return nil, err
		}
		fmt.Printf("%s is %d years old and weighs %f kg\n", name, age, weight)
	}

	return nil, nil
}
