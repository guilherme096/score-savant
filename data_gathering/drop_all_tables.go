package main

import (
	"database/sql"
	"fmt"
	"log"

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

func main() {
	// Build connection string
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", user, password, server, port, database)

	// Open connection
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// Drop all foreign key constraints
	err = dropForeignKeyConstraints(db)
	if err != nil {
		log.Fatalf("Failed to drop foreign key constraints: %s", err)
	}

	// Drop all tables
	err = dropAllTables(db)
	if err != nil {
		log.Fatalf("Failed to drop tables: %s", err)
	}

	fmt.Println("All tables and foreign key constraints dropped successfully.")
}

func dropForeignKeyConstraints(db *sql.DB) error {
	query := `
        SELECT
            fk.name AS FK_Name,
            tp.name AS TableName
        FROM 
            sys.foreign_keys AS fk
        INNER JOIN 
            sys.tables AS tp ON fk.parent_object_id = tp.object_id
    `
	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("failed to query foreign key constraints: %w", err)
	}
	defer rows.Close()

	var fkName, tableName string
	for rows.Next() {
		err := rows.Scan(&fkName, &tableName)
		if err != nil {
			return fmt.Errorf("failed to scan foreign key constraint: %w", err)
		}

		dropFKStmt := fmt.Sprintf("ALTER TABLE %s DROP CONSTRAINT %s", tableName, fkName)
		_, err = db.Exec(dropFKStmt)
		if err != nil {
			return fmt.Errorf("failed to drop foreign key constraint %s on table %s: %w", fkName, tableName, err)
		}
	}

	return nil
}

func dropAllTables(db *sql.DB) error {
	query := `
        SELECT TABLE_NAME
        FROM INFORMATION_SCHEMA.TABLES
        WHERE TABLE_TYPE = 'BASE TABLE'
    `
	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("failed to query table names: %w", err)
	}
	defer rows.Close()

	var tableName string
	for rows.Next() {
		err := rows.Scan(&tableName)
		if err != nil {
			return fmt.Errorf("failed to scan table name: %w", err)
		}

		dropTableStmt := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
		_, err = db.Exec(dropTableStmt)
		if err != nil {
			return fmt.Errorf("failed to drop table %s: %w", tableName, err)
		}
	}

	return nil
}
