package storage

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

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

	// execute stored procedure
	rows, err := m.db.Query("SELECT * FROM GetPlayerById(@player_id)", sql.Named("player_id", id))
	if err != nil {
		return nil, err
	}
	// close when the function ends
	defer rows.Close()

	// get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// create a slice of interfaces to store the values from the database
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	result := make(map[string]interface{})

	// get the values from each row
	if rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, err
		}

		for i, col := range columns {
			val := values[i]

			// If the value is nil, set it to a zero value
			if val == nil {
				result[col] = nil
			} else {
				switch v := val.(type) {
				case int64:
					result[col] = int(v)
				case int:
					result[col] = int(v)
				case []uint8:
					// Convert []uint8 to string then to float64
					strVal := string(v)
					floatVal, err := strconv.ParseFloat(strVal, 64)
					if err != nil {
						return nil, fmt.Errorf("error converting %s to float64: %v", col, err)
					}
					result[col] = floatVal
				case time.Time:
					result[col] = strings.Split(v.String(), " ")[0]
				default:
					result[col] = val
				}
			}
		}
	} else {
		return nil, fmt.Errorf("player with id %s not found", id)
	}

	// convert int64 to int if needed
	for key, value := range result {
		switch v := value.(type) {
		case int64:
			result[key] = int(v)

		case float64:
			result[key] = v
		}
	}

	return result, nil

}
