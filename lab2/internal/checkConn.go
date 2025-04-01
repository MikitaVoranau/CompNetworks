package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CheckConnection() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	var ip, port string
	fmt.Println("Enter your IP(by default 127.0.0.1): ")
	ip, _ = reader.ReadString('\n')
	ip = strings.TrimSpace(ip)
	if ip == "" {
		ip = "127.0.0.1"
	}
	fmt.Println("Enter your port(by default 8080): ")
	port, _ = reader.ReadString('\n')
	port = strings.TrimSpace(port)
	if port == "" {
		port = "8080"
	}
	return ip, port
}
