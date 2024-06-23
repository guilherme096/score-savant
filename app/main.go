package main

import (
	"fmt"
	server "guilherme096/score-savant/api"
	mssqlStorage "guilherme096/score-savant/storage/Mssql"
)

func main() {
	listen_addr := ":8080"

	fmt.Println("App listening on port: ", listen_addr)

	db := mssqlStorage.NewMSqlStorage(".", ".", "mednat.ieeta.pt", 8101, "p5g5")

	// Create a new server
	server := server.New_server(listen_addr, db)

	server.Start()
}
