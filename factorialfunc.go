package main

import (
	"fmt"
	"github.com/senthilkgithub/learngolang/factorial"
	"github.com/senthilkgithub/learngolang/getChar"
)

func main() {
	var input uint
	fmt.Println("Please Enter value to find out factorial")
	fmt.Scanf("%d\n", &input)
	fact, err := factorial.Factorial(float64(input))
	if err != nil {
		fmt.Println(err.Error())
		getChar.GetChar()
	}
	fmt.Println("The factorial value of %d\t", input, " is %f\t", fact)
	fmt.Println("Press any key to exit")
	getChar.GetChar()
}
