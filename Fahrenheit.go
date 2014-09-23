package main

import "fmt"

func main() {

	var fahrenheit float32
	var waittoUserINteract string
	fmt.Println("Enter Fahrenheit value")
	fmt.Scanf("%f", &fahrenheit)
	centigrade := (fahrenheit - 32) * (5 / 9)
	fmt.Println("Centigrade value is %f", centigrade)
	fmt.Println("Press any key to exit")
	fmt.Scanf("%s", &waittoUserINteract)
}
