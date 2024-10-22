package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Aidi Adam Baik"
	// encode base64
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	// decode base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}
}
