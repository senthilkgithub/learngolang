package main

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}
func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	kids := []Person{
		{"Jill", 9}, {"Jack", 10}, {"mack", 10}, {"rack", 10}, {"buck", 10}, {"disk", 10}, {"Jursk", 10}, {"wirsk", 10},
	}
	sort.Sort(ByName(kids))

	fmt.Println(kids)

	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)
}
