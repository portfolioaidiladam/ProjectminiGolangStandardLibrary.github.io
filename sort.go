package main
// sort adalah package untuk mengurutkan data
import (
	"fmt"
	"sort"
)
// struct data
type User struct {
	Name string
	Age  int
}

// slice data
type UserSlice []User

// len data
func (s UserSlice) Len() int {
	return len(s)
}

// less data
func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

// swap data
func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Aidil", 30},
		{"Budi", 35},
		{"Joko", 25},
		{"Adit", 20},
	}
	
 // mengurutkan data
	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
