package main

import (
	"fmt"
	server "guilherme096/score-savant/api"
)

func main() {
	listen_addr := ":8080"

	fmt.Println("App listening on port: ", listen_addr)

	// Create a new server
	server := server.New_server(listen_addr)

	server.Start()
}
