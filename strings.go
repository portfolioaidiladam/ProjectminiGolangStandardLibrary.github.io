package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Aidil Adam", "Aidil"))
	fmt.Println(strings.Split("Aidil Adam", " "))
	fmt.Println(strings.ToLower("Aidil Adam"))
	fmt.Println(strings.ToUpper("Aidil Adam"))
	fmt.Println(strings.Trim("      Aidil Adam      ", " "))
	fmt.Println(strings.ReplaceAll("Aidil Adam Aidil Adam", "Aidil", "Budi"))
}
