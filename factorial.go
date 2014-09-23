package main

import (
	"fmt"
)

func main() {

	var (
		keytoGetFactorial int64
		factvalue         int64
		i                 int64
	)
	keytoGetFactorial = 54
	var pressKey string
	//fmt.Println("Please enter number to find factorial value")
	//fmt.Scanf("%d", &KeytoGetFactorial)

	factvalue = 1
	for i = 1; i <= keytoGetFactorial; i++ {
		factvalue = factvalue * i
	}
	fmt.Println(factvalue)
	fmt.Scanf("%s", &pressKey)
	fmt.Println("Pressed Key %s", pressKey)
}
