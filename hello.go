package main

import "fmt"

// this is a comment

func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World")
	fmt.Println("Hello " + "World")

	x := ""
	x += "new string"
	fmt.Println(x)
	const const_x = "jksdbhfjksgjkh"
	var inputFloat float32
	fmt.Println("Enter some float number")
	fmt.Scanf("%f", &inputFloat)
	result := inputFloat * 10
	fmt.Println(result)
	var userresponse string
	fmt.Scanf("%f", &userresponse)
	//fmt.Println("x value is :" + const_x + " and y value is: " + var_y + "and their comparisition is " + equlityComparer(const_x ,var_y))

}

/*func equlityComparer(x,y){
	if(x==y)
	return true;
}*/
