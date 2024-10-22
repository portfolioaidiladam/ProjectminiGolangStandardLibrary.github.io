package main

import (
	"bufio"
	"os"
)

func main() {
	// create writer	
	writer := bufio.NewWriter(os.Stdout)
	// write string
	_, _ = writer.WriteString("hello world\n")
	_, _ = writer.WriteString("Selamat belajar\n")
	// flush
	writer.Flush()
}
