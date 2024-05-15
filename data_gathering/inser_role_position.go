package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
)

// Database connection details
var (
	server   = "mednat.ieeta.pt"
	port     = 8101
	user     = "p5g5"
	password = "bo_jack64"
	database = "p5g5"
)

// Function to get role ID based on role name
func getRoleID(db *sql.DB, roleName string) (int, error) {
	var roleID int
	err := db.QueryRow("SELECT role_id FROM Role WHERE name = @name", sql.Named("name", roleName)).Scan(&roleID)
	return roleID, err
}

// Function to get position ID based on position name
func getPositionID(db *sql.DB, positionName string) (int, error) {
	var positionID int
	err := db.QueryRow("SELECT position_id FROM Position WHERE name = @name", sql.Named("name", positionName)).Scan(&positionID)
	return positionID, err
}

// Function to insert role position
func insertRolePosition(db *sql.DB, roleID, positionID int) error {
	_, err := db.Exec("INSERT INTO RolePosition (role_position, position_id) VALUES (@role_id, @position_id)", sql.Named("role_id", roleID), sql.Named("position_id", positionID))
	return err
}

func main() {
	// Build connection string
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", user, password, server, port, database)

	// Open connection
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// Open the CSV file
	file, err := os.Open("keyatts.csv")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("failed to read CSV file: %s", err)
	}

	// Get the headers and indices
	headers := records[0]
	var roleIndex, positionIndex int
	for i, header := range headers {
		switch header {
		case "Role":
			roleIndex = i
		case "Position":
			positionIndex = i
		}
	}

	// Iterate over the records starting from the second row
	for _, record := range records[1:] {
		role := record[roleIndex]
		position := record[positionIndex]

		splited_position := strings.Split(position, ",")

		for _, pos := range splited_position {
			// Trim the role and position
			role = strings.TrimSpace(role)
			pos = strings.TrimSpace(pos)

			// Get the role ID based on the role name
			roleID, err := getRoleID(db, role)
			if err != nil {
				log.Printf("failed to get role ID for role '%s': %s, skipping", role, err)
				continue
			}

			// Get the position ID based on the position name
			positionID, err := getPositionID(db, pos)
			if err != nil {
				log.Printf("failed to get position ID for position '%s': %s, skipping", pos, err)
				continue
			}

			// Insert role-position mapping
			err = insertRolePosition(db, roleID, positionID)
			if err != nil {
				log.Fatalf("failed to insert role position: %s", err)
			}
		}
	}

	fmt.Println("Role positions inserted successfully.")
}
