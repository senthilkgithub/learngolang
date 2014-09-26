package reversenumber

import "errors"

func ReverseNumber(num uint) (uint, error) {

	if countDigitsInNumber(num) <= 1 {
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

func countDigitsInNumber(numb uint) (digit int) {
	digit = 0
	for numb != 0 && digit != 2 {
		numb = numb / 10
		digit++

	}
	return digit
}
