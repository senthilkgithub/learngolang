package main

import (
	//"bufio"
	"fmt"
	"github.com/senthilkgithub/learngolang/factorial"
	"github.com/senthilkgithub/learngolang/reversenumber"
	"github.com/senthilkgithub/learngolang/swapnumber"
	//"os"
	//"strconv"
)

var (
	availableOperations [6]string
	selection           int
	yesOrNo             string
)

func main() {
	Manipulator()
}

func Manipulator() int {
	availableOperations := [8]string{
		"1.Factorial",
		"2.FaherenHeit To Centigrade",
		"3.Reverse Number",
		"4.Swap Numbers",
		"5.Fibonacci Series",
		"6.Star Pyramid",
		"7.Triangle",
		"8.Exit",
	}
	for _, value := range availableOperations {
		fmt.Println(value)
	}

	fmt.Println("Select your operation :")
	n, err := fmt.Scanf("%d\n", &selection)
	if err != nil || n != 1 {
		// handle invalid input
		fmt.Println(n, err)
		return 0
	}

	if selection > 0 {
		switch selection {
		case 1:
			findFactorial()
		case 2:
			faherenheittocelcius()
		case 3:
			math_reversenumber()
		case 4:
			swapNumbers()
		case 5:
			GetFibonacciSeries()
		case 6:
			StarPyramid()
		case 7:
			math_Triangle()

		}
		fmt.Println("Do you want to continue again Y/N :")
		n, err := fmt.Scanf("%s\n", &yesOrNo)
		if err != nil || n != 1 {
			// handle invalid input
			fmt.Println(n, err)
			return 0
		}

		if yesOrNo == "y" || yesOrNo == "Y" {
			Manipulator()
		}

	}
	return 0
}

func faherenheittocelcius() int {
	var fahrenheit float32
	fmt.Println("Enter Fahrenheit value")
	n, err := fmt.Scanf("%f\n", &fahrenheit)
	if err != nil || n != 1 {
		// handle invalid input
		fmt.Println(n, err)
		return 0
	}
	centigrade := (fahrenheit - 32.0) * (5.0 / 9.0)
	fmt.Println("Centigrade value is :", centigrade)
	return 0
}

func findFactorial() int {
	var input float64
	fmt.Println("Please Enter value to find out factorial")
	n, err := fmt.Scanf("%f\n", &input)
	if err != nil || n < 0 {
		fmt.Println(n, err)
		return 0
	}
	fact, err := factorial.Factorial(input)
	fmt.Println("The factorial value of ", input, " is ", fact)

	return 0
}

func math_reversenumber() float64 {
	var input uint
	fmt.Println("Enter a number to reverse\n")
	n, err := fmt.Scanf("%d\n", &input)
	if err != nil || n < 0 {
		// handle invalid input
		fmt.Println(n, err)
		return 0
	}
	rev, err := reversenumber.ReverseNumber(input)
	fmt.Println("Reverse of entered number is ", rev)
	return 0
}

func swapNumbers() float64 {

	var (
		first  float64
		second float64
	)
	fmt.Printf("Enter two integers to add\n")
	n, err := fmt.Scanf("%d%d\n", &first, &second)
	if err != nil || n != 1 {
		fmt.Println(err)
	}

	fmt.Printf("Value before swap first = %f second = %f\n", first, second)
	first, second, err = swapnumber.Swap(first, second)
	fmt.Printf("Value after swap first = %f second = %f\n", first, second)
	return 0
}

func GetFibonacciSeries() int {
	var input int
	fmt.Printf("Enter extend of series:")
	n, err := fmt.Scanf("%d\n", &input)
	if err != nil || n != 1 {
		fmt.Println(err)
	}
	fab(input)
	return 0
}

func fab(num int) int {
	num1 := 0.0
	num2 := 1.0
	var (
		sum float64
		i   int
	)
	fmt.Printf("fab series is\n%d\n%d", num1, num2)
	for i = 0; i < num-2; i++ {
		sum = num1 + num2
		num1 = num2
		num2 = sum
		fmt.Println(sum)
	}
	return 0
}

func StarPyramid() int {
	var (
		rows int
	)
	k := 0
	fmt.Printf("Enter the number of rows: ")
	n, err := fmt.Scanf("%d\n", &rows)
	if err != nil || n != 1 {
		fmt.Println(err)
	}
	for i := 1; i <= rows; i++ {
		for space := 1; space <= rows-i; space++ {
			fmt.Printf("  ")
		}
		for k != 2*i-1 {
			fmt.Printf("* ")
			k++
		}
		k = 0
		fmt.Printf("\n")
	}
	return 0
}

func math_Triangle() int {
	var input int
	fmt.Println("Enter number of rows: ")
	n, err := fmt.Scanf("%d\n", &input)
	if err != nil || n < 0 {
		fmt.Println(n, err)
		return 0
	}
	for i := 1; i <= input; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d\t", j)
		}
		fmt.Println("")
	}
	return 0
}
