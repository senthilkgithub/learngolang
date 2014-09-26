package getChar

import "fmt"

func GetChar() {
	var inputKey string
	count, err := fmt.Scanf("%s\n", &inputKey)
	if count != 1 {
		fmt.Println(inputKey, err)
	}
}
