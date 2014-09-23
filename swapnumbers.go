package main

import "fmt"

func main() {
	var inputKey string
	var (
		first  int
		second int
	)
	fmt.Printf("Enter two integers to add\n")
	fmt.Scanf("%d%d", &first, &second)

	fmt.Printf("Value before swap first = %d second = %d\n", first, second)
	swap(&first, &second)
	fmt.Printf("Value after swap first = %d second = %d\n", first, second)
	count := 0
	count, _ = fmt.Scanf("%s", &inputKey)
	if count <= 0 {
		fmt.Scanf("%s", &inputKey)
	}
}
func swap(first, second *int) int {
	*first = *first + *second
	*second = *first - *second
	*first = *first - *second
	return 1
}
