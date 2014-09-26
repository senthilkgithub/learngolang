package main

import (
	"fmt"
	"github.com/senthilkgithub/learngolang/swapnumber"
)

func main() {
	var (
		first  float64
		second float64
		errval error
	)
	fmt.Printf("Enter two integers to add\n")
	_, err := fmt.Scanf("%f\n%f\n", &first, &second)
	if err != nil {
		fmt.Println(err)
		getChar()
	} else {
		fmt.Printf("Value before swap first = %f second = %f\n", first, second)
		first, second, errval = swapnumber.Swap(first, second)
		if errval != nil {
			fmt.Println(errval)
			getChar()
		} else {
			fmt.Printf("Value after swap first = %f second = %f\n", first, second)
			getChar()
			return
		}
	}

}
func getChar() {
	var inputKey string
	count, err := fmt.Scanf("%s\n", &inputKey)
	if count != 1 {
		fmt.Println(inputKey, err)
	}
}
