package main

import (
	"bufio"
	"fmt"
	"lab2/api/client"
	"lab2/api/server"
	"os"
	"strings"
)

func main() {
	var choose string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Choose startup mode: ")
		fmt.Println("1. Server")
		fmt.Println("2. Client")
		fmt.Println("Press enter to quit.")
		choose, _ = reader.ReadString('\n')
		choose = strings.TrimSpace(choose)
		switch choose {
		case "1":
			server.RunServer()
		case "2":
			client.RunClient()
		case "":
			os.Exit(1)
		default:
			fmt.Println("Error choosing option")
			fmt.Println()
		}
	}
}
