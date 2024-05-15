package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
)

// Database connection details
var (
	insert_server = "mednat.ieeta.pt"
	port          = 8101
	user          = "p5g5"
	password      = "bo_jack64"
	database      = "p5g5"
)

func main() {
	// Build connection string
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", user, password, insert_server, port, database)

	// Open connection
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// Open the file
	file, err := os.Open("attributes.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentType string
	typeTableMapping := map[string]string{
		"Technical":   "Technical_Att",
		"Physical":    "Physical_Att",
		"Mental":      "Mental_Att",
		"Goalkeeping": "Goalkeeping_Att",
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if _, ok := typeTableMapping[line]; ok {
			currentType = line
			// Skip the blank line after the attribute type
			scanner.Scan()
			continue
		}
		if currentType != "" {
			attributeName := line
			// Insert attribute and get its ID
			var attributeID string
			err := db.QueryRow("INSERT INTO Attribute (name) OUTPUT INSERTED.name VALUES (@p1)", attributeName).Scan(&attributeID)
			if err != nil {
				log.Fatalf("failed to insert attribute: %s", err)
			}
			// Insert into the corresponding attribute type table
			tableName := typeTableMapping[currentType]
			_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (att_id) VALUES (@p1)", tableName), attributeID)
			if err != nil {
				log.Fatalf("failed to insert into %s: %s", tableName, err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	fmt.Println("Attributes inserted successfully.")
}
