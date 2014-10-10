//Reversing uint type number
package reversenumber

import "errors"

//Accepts uint type number of more than single digit and returns reverse of that number and error if it failed by matching criteria
func ReverseNumber(num uint) (uint, error) {

	if CountDigitsInNumber(num) <= 1 {
		return 0, errors.New("Atleaset two digit values expected")
	} else {
		var rev uint
		for num != 0 {
			rev = rev * 10
			rev = rev + num%10
			num = num / 10
		}
		return rev, nil
	}
}

//Short method for counting number of digits presented in given number
// Purpose of the method is to findout More than one digit numbers
func CountDigitsInNumber(numb uint) (digit int) {
	digit = 0
	for numb != 0 && digit != 2 {
		numb = numb / 10
		digit++

	}
	return digit
}
