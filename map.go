package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Florine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["B"])
	name, ok := elements["Be"]
	fmt.Println(name, ok)
	if name, ok := elements["Ne"]; ok {
		fmt.Println(name, ok)
	}

	myslice := make([]int, 2, 9)
	fmt.Println(len(myslice), myslice)
	x := [6]string{"a", "b", "c", "d", "e", "f"}

	fmt.Println(x[2:5])
}
