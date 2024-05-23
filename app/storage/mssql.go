package storage

import (
	"database/sql"
	"fmt"

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

func (m *MSqlStorage) LoadPlayerById(id string) (map[string]interface{}, error) {
	rows, err := m.db.Query("SELECT * FROM GetPlayerById(@player_id)", sql.Named("player_id", id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Prepare a slice of interfaces to hold column values
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	result := make(map[string]interface{})

	if rows.Next() {
		// Scan the result into the slice of interfaces
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, err
		}

		// Populate the map with the column names and values
		for i, col := range columns {
			val := values[i]

			// If the value is nil, set it to a zero value
			if val == nil {
				result[col] = nil
			} else {
				result[col] = val
			}
		}
	} else {
		return nil, fmt.Errorf("player with id %s not found", id)
	}
	// Example of accessing and handling int64 values
	for key, value := range result {
		switch v := value.(type) {
		case int64:
			// Convert int64 to int if needed
			result[key] = int(v)
		}
	}

	return result, nil

}
