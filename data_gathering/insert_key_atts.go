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

// Function to insert role and get its ID
func insertRole(db *sql.DB, roleName string) (int, error) {
	var roleID int
	err := db.QueryRow(`
        IF NOT EXISTS (SELECT role_id FROM Role WHERE name = @name)
        BEGIN
            INSERT INTO Role (name) OUTPUT INSERTED.role_id VALUES (@name)
        END
        ELSE
        SELECT role_id FROM Role WHERE name = @name
    `, sql.Named("name", roleName)).Scan(&roleID)
	return roleID, err
}

// Function to insert attribute and get its ID, ignoring duplicates
func insertAttribute(db *sql.DB, attributeName string) (string, error) {
	var attributeID string
	err := db.QueryRow(`
        SELECT name FROM Attribute WHERE name = @name
    `, sql.Named("name", attributeName)).Scan(&attributeID)
	return attributeID, err
}

// Function to insert key attribute
func insertKeyAttribute(db *sql.DB, roleID int, attributeID string) error {
	_, err := db.Exec("INSERT INTO KeyAttributes (role_id, attribute_id) VALUES (@role_id, @attribute_id)", sql.Named("role_id", roleID), sql.Named("attribute_id", attributeID))
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
	var roleIndex, keyAttrIndex int
	for i, header := range headers {
		switch header {
		case "Role":
			roleIndex = i
		case "Key Attr":
			keyAttrIndex = i
		}
	}

	// Insert unique roles
	roleSet := make(map[string]struct{})
	for _, record := range records[1:] {
		role := record[roleIndex]
		if _, exists := roleSet[role]; !exists {
			roleSet[role] = struct{}{}
			_, err := insertRole(db, role)
			if err != nil {
				log.Fatalf("failed to insert role: %s", err)
			}
		}
	}

	// Iterate over the records starting from the second row
	for _, record := range records[1:] {
		role := record[roleIndex]
		keyAttributes := record[keyAttrIndex]

		// Get the role ID based on the role name
		roleID, err := insertRole(db, role)
		if err != nil {
			log.Fatalf("failed to get role ID for role '%s': %s", role, err)
		}

		// Split key attributes
		attributes := strings.Split(keyAttributes, ",")
		for _, attribute := range attributes {
			attribute = strings.TrimSpace(attribute)
			if attribute != "" {
				// Insert attribute and get its ID, ignoring duplicates
				attributeID, err := insertAttribute(db, attribute)
				if err != nil {
					log.Printf("failed to insert attribute: %s, ignoring", err)
					continue
				}
				// Insert key attribute

				err = insertKeyAttribute(db, roleID, attributeID)
				fmt.Println("Role ID: ", roleID)
				fmt.Println("Attribute ID: ", attributeID)
				if err != nil {
					log.Fatalf("failed to insert key attribute: %s", err)
				}
			}
		}
	}

	fmt.Println("Roles and key attributes inserted successfully.")
}
