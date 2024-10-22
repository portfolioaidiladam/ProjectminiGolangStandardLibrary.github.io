package main

import (
	"fmt"
	"regexp"
)

func main() {
	// compile regex
	var regex *regexp.Regexp = regexp.MustCompile(`a[a-z]+l`)
	//var regex *regexp.Regexp = regexp.MustCompile(`a[a-zA-Z]+l`)
	// match string
	fmt.Println(regex.MatchString("aidil"))
	fmt.Println(regex.MatchString("adil"))
	fmt.Println(regex.MatchString("aiDil"))
	// find all string
	fmt.Println(regex.FindAllString("aidil adil adal adul adam adem aiDil", 10))
}
