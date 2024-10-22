package main

import (
	"fmt"
	"path"
)

func main() {
	// dir
	fmt.Println(path.Dir("hello/world.go"))
	// base
	fmt.Println(path.Base("hello/world.go"))
	// ext
	fmt.Println(path.Ext("hello/world.go"))
	// join
	fmt.Println(path.Join("hello", "world", "main.go"))
}
