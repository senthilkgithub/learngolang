package main

import (
	//"bufio"
	"fmt"
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
	var input uint
	fmt.Println("Please Enter value to find out factorial")
	n, err := fmt.Scanf("%d\n", &input)
	if err != nil || n < 0 {
		// handle invalid input
		fmt.Println(n, err)
		return 0
	}
	fmt.Println("The factorial value of ", input, " is ", factorial(input))

	return 0
}

func factorial(x uint) uint {

	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func math_reversenumber() int {
	var input int
	reverse := new(int)
	fmt.Println("Enter a number to reverse\n")
	n, err := fmt.Scanf("%d\n", &input)
	if err != nil || n < 0 {
		// handle invalid input
		fmt.Println(n, err)
		return 0
	}

	fmt.Println("Reverse of entered number is ", reversenumber(&input, reverse))
	return 0
}

func reversenumber(num, rev *int) int {

	for *num != 0 {
		*rev = *rev * 10
		*rev = *rev + *num%10
		*num = *num / 10
	}
	return *rev
}

func swapNumbers() int {

	var (
		first  int
		second int
	)
	fmt.Printf("Enter two integers to add\n")
	n, err := fmt.Scanf("%d%d\n", &first, &second)
	if err != nil || n != 1 {
		fmt.Println(err)
	}

	fmt.Printf("Value before swap first = %d second = %d\n", first, second)
	swap(&first, &second)
	fmt.Printf("Value after swap first = %d second = %d\n", first, second)
	return 0
}

func swap(first, second *int) int {
	*first = *first + *second
	*second = *first - *second
	*first = *first - *second
	return 1
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
