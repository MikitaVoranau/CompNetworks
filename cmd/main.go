package cmd

import (
	"CN_lab2/api/client"
	"CN_lab2/api/server"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Choose startup mode: ")
	fmt.Println("1. Server")
	fmt.Println("2. Client")
	var choose int
	_, _ = fmt.Scan(&choose)
	switch choose {
	case 1:
		server.RunServer()
	case 2:
		client.RunClient()
	default:
		fmt.Println("Error choosing option")
		os.Exit(1)
	}

}
