package main

import "fmt"

func main() {

	var input uint
	var inputKey string
	fmt.Println("PLease Enter value to find out factorial")
	fmt.Scanf("%d", &input)
	fmt.Println("The factorial value of %d", input, " is %d", factorial(input))
	fmt.Println("Press any key to exit")
	count := 0
	count, _ = fmt.Scanf("%s", &inputKey)
	if count <= 0 {
		fmt.Scanf("%s", &inputKey)
	}
	fmt.Println("Bye")
}

func factorial(x uint) uint {

	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
