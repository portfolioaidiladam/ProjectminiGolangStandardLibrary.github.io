package main

import (
	"bufio"
	"io"
	"os"
	// "fmt"
)

func createNewFile(name string, message string) error {
	// membuat file baru
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	// validasi error
	if err != nil {
		return err
	}
	// menutup file
	defer file.Close()
	// menulis data ke file
	file.WriteString(message)
	return nil
}

func addToFile(name string, message string) error {
	// menambahkan data ke file
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	// validasi error
	if err != nil {
		return err
	}
	// menutup file
	defer file.Close()
	// menulis data ke file
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	// membaca file
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	// validasi error
	if err != nil {
		return "", err
	}
	// menutup file
	defer file.Close()

	// membaca file
	reader := bufio.NewReader(file)
	// membuat variable untuk menampung data
	var message string
	// membaca file per baris
	for {
		// membaca file per baris
		line, _, err := reader.ReadLine()
		// validasi error
		if err == io.EOF {
			break
		}
		// menambahkan data ke variable
		message += string(line) + "\n"
	}
	// mengembalikan data
	return message, nil
}

func main() {
	// membuat file baru
	// createNewFile("sample.log", "this is sample log")
	
	// membaca file
	// result, _ := readFile("sample.log")
	// menampilkan data
	// fmt.Println(result)

	// menambahkan data ke file
	addToFile("sample.log", "\nthis is add message")
}
