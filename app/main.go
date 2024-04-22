package main

import (
	"fmt"
	server "guilherme096/score-savant/api"
	storage "guilherme096/score-savant/storage"
)

func main() {
	listen_addr := ":8080"

	fmt.Println("App listening on port: ", listen_addr)

	db := storage.NewMemoryStorage()

	// Create a new server
	server := server.New_server(listen_addr, db)

	server.Start()
}
