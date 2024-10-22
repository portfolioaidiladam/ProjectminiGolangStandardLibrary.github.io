package main

import (
	"container/list"
	"fmt"
)

func main() {
	// struktur data double linked list
	var data *list.List = list.New()

	data.PushBack("Aidil")
	data.PushBack("Adam")
	data.PushBack("Baik")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // aidil

	next := head.Next() // adam
	fmt.Println(next.Value)

	next = next.Next() // baik
	fmt.Println(next.Value)

	// dengan for
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
