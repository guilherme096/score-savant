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

func (m *MSqlStorage) LoadPlayerById(id string) (map[string]interface{}, []map[string]interface{}, error) {

	// execute stored procedure
	rows, err := m.db.Query("SELECT * FROM GetPlayerById(@player_id)", sql.Named("player_id", id))
	// close when the function ends
	defer rows.Close()

	// get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, nil, err
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
			return nil, nil, err
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
						return nil, nil, fmt.Errorf("error converting %s to float64: %v", col, err)
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
		return nil, nil, fmt.Errorf("player with id %s not found", id)
	}

	// Execute query to fetch player attributes
	attributesRows, err := m.db.Query("SELECT * FROM GetPlayerAttributes(@player_id)", sql.Named("player_id", id))
	if err != nil {
		return nil, nil, err
	}
	defer attributesRows.Close()

	// Initialize a slice of maps to hold the attribute data
	var attributes []map[string]interface{}

	// Iterate over each row of attributes
	for attributesRows.Next() {
		// Get column names for attributes
		attributeColumns, err := attributesRows.Columns()
		if err != nil {
			return nil, nil, err
		}

		// Create a slice to hold the values of each attribute row
		var attributeValues []interface{}

		// Get the values from the current row
		attributeValues, err = scanValues(attributesRows, attributeColumns)
		if err != nil {
			return nil, nil, err
		}

		// Create a map for the current attribute row
		attributeRow := make(map[string]interface{})

		// Populate attributeRow map with column names and values
		for i, col := range attributeColumns {
			var convertedvalue interface{}
			switch v := attributeValues[i].(type) {
			case int64:
				convertedvalue = int(v)
			default:
				convertedvalue = v
			}
			attributeRow[col] = convertedvalue
		}

		// Add attributeRow to the slice of attributes
		attributes = append(attributes, attributeRow)
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

	return result, attributes, nil
}

func (m *MSqlStorage) GetAttributeList(att_type string) []string {
	query := ""
	switch att_type {
	case "Physical":
		query = "SELECT * FROM Physical_Att"
	case "Mental":
		query = "SELECT * FROM Mental_Att"
	case "Technical":
		query = "SELECT * FROM Technical_Att"
	case "Goalkeeping":
		query = "SELECT * FROM Goalkeeping_Att"
	}

	rows, err := m.db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var atts []string

	for rows.Next() {
		var att string
		err := rows.Scan(&att)
		if err != nil {
			panic(err)
		}
		atts = append(atts, att)
	}

	return atts
}

func (m *MSqlStorage) GetPlayerPosition(id string) (int, string, error) {
	rows, err := m.db.Query("SELECT * FROM GetPositionByPlayerID(@player_id)", sql.Named("player_id", id))
	if err != nil {
		fmt.Println(err)
		return -1, "", err
	}
	defer rows.Close()

	var position_id int
	var position_name string

	if rows.Next() {
		err := rows.Scan(&position_id, &position_name)
		if err != nil {
			return -1, "", err
		}
	} else {
		return -1, "", fmt.Errorf("player with id %s not found", id)
	}

	return position_id, position_name, nil
}

// Function to scan values from a row into a slice of interfaces
func scanValues(rows *sql.Rows, columns []string) ([]interface{}, error) {
	// Create a slice to hold the values of each row
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	// Scan the values from the row into the slice of interfaces
	err := rows.Scan(valuePtrs...)
	if err != nil {
		return nil, err
	}

	return values, nil
}
