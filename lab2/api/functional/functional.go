package functional

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"lab2/internal"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func receiveMessages(c net.Conn) {
	var msg internal.Message
	for {

	}
}

func readMessage(c net.Conn, msg *internal.Message) {

}

func sendingMessages(c net.Conn) {
	reader := bufio.NewReader(c)
	fmt.Println("You can send:")
	fmt.Println("1. Text.")
	fmt.Printf("2. File. \n" +
		"If you want to send a file, the path to your file should be wrapped with '<>'.\n" +
		"Example: '<D:\\Git\\bin>'")
	fmt.Println("Press Enter to exit.")
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			return
		}
		isHaveFile, filePath, err := checkForFile(text)
		if err != nil {
			fmt.Println("Error checking for file:", err)
			continue
		}

		switch isHaveFile {
		case true:
			content, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Ошибка при чтении файла: %v\n", err)
				continue
			}
			msg := &internal.Message{
				Type:    internal.MessageTypeFile,
				Content: content,
				Name:    filepath.Base(filePath),
			}
			err = sendMessage(c, msg)
			if err != nil {
				fmt.Printf("Ошибка при отправке файла: %v\n", err)
				return
			}

			fmt.Printf("Файл %s отправлен\n", msg.Name)
		case false:
			msg := &internal.Message{
				Type:    internal.MessageTypeText,
				Content: []byte(text),
			}
			err := sendMessage(c, msg)
			if err != nil {
				fmt.Println("Error reading message:", err)
			}
		}
	}
}

func checkForFile(text string) (bool, string, error) {
	isHavePath := containsInOrder(text, '<', '>')

	if !isHavePath {
		return false, "", nil
	}

	start := strings.IndexRune(text, '<')
	end := strings.IndexRune(text, '>')

	if start < end {
		return true, text[start+1 : end], nil
	}

	return false, "", nil
}

func containsInOrder(str string, char1, char2 rune) bool {
	index1 := strings.IndexRune(str, char1)
	index2 := strings.IndexRune(str, char2)

	return index1 != -1 && index2 != -1 && index1 < index2
}

func sendMessage(c net.Conn, msg *internal.Message) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, msg.Type)
	if err != nil {
		return fmt.Errorf("sendMessage: %w", err)
	}

	if msg.Type == internal.MessageTypeFile {

		err := binary.Write(buf, binary.BigEndian, msg.Name)
		if err != nil {
			return fmt.Errorf("sendMessage: %w", err)
		}
		_, err = buf.WriteString(msg.Name)
		if err != nil {
			return err
		}

	}
	contentLen := uint32(len(msg.Content))
	err = binary.Write(buf, binary.BigEndian, contentLen)
	if err != nil {
		return err
	}

	_, err = buf.Write(msg.Content)
	if err != nil {
		return err
	}

	_, err = c.Write(buf.Bytes())
	return err
}
