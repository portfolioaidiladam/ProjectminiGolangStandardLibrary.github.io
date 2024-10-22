package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	// input string				
	input := strings.NewReader("this is long string\nwith new line\n")
	// create reader
	reader := bufio.NewReader(input)

	for {
		// read line	
		line, _, err := reader.ReadLine()
		// if error is end of file then break
		if err == io.EOF {
			break
		}

		// print line
		fmt.Println(string(line))
	}
}
