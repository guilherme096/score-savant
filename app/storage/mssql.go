package storage

import (
	"database/sql"
	"fmt"
	Player "guilherme096/score-savant/models"

	_ "github.com/microsoft/go-mssqldb"
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
}

func (m *MSqlStorage) Stop() {
	m.db.Close()
	fmt.Println("Disconnected from SQL Server")
}

func (m *MSqlStorage) LoadPlayerById(id string) (*Player.Player, error) {
	query := "SELECT * FROM Player WHERE player_id=@id"
	prep, err := m.db.Prepare(query)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}
	defer prep.Close()

	rows, err := prep.Query(sql.Named("id", id))

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	defer rows.Close()
	fmt.Println("Player found")
	var player *Player.Player = new(Player.Player)
	var PlayerBio *Player.PlayerBio = new(Player.PlayerBio)
	player.PlayerBio = PlayerBio
	rows.Next()
	fmt.Println(rows.Scan(&player.Id, &player.PlayerBio.Name, &player.PlayerBio.Age, &player.PlayerBio.Weight, &player.PlayerBio.Height, &player.PlayerBio.Nation, &player.Contract, &player.PlayerBio.Foot, &player.TechnicalAttributes))
	fmt.Println(player.PlayerBio.Name)
	return nil, nil
}
