package main

import "fmt"

func main() {
	var (
		inputKey string
		n        int
	)
	reverse := new(int)
	fmt.Println("Enter a number to reverse\n")
	fmt.Scanf("%d", &n)
	fmt.Println("Reverse of entered number is ", reversenumber(&n, reverse))

	count := 0
	count, _ = fmt.Scanf("%s", &inputKey)
	if count <= 0 {
		fmt.Scanf("%s", &inputKey)
	}

}

func reversenumber(num, rev *int) int {

	for *num != 0 {
		*rev = *rev * 10
		*rev = *rev + *num%10
		*num = *num / 10
	}
	return *rev

}
