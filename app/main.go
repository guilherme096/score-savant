package main

import (
	"fmt"
	server "guilherme096/score-savant/api"
	storage "guilherme096/score-savant/storage"
)

func main() {
	listen_addr := ":8080"

	fmt.Println("App listening on port: ", listen_addr)

	//db := storage.NewMSqlStorage("p5g5", "bo_jack64", "mednat.ieeta.pt", 8101, "p5g5")
	db := storage.NewMemoryStorage()

	// Create a new server
	server := server.New_server(listen_addr, db)

	server.Start()
}
