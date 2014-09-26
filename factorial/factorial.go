package factorial

import "errors"

func Factorial(x float64) (float64, error) {

	if x < 0 {
		return 0, errors.New("Negetive value is not allowed")
	}
	if x > 45 {
		return 0, errors.New("Number above 45 is not allowed,The system is will not process output")
	}
	if x == 0 {
		return 1, nil
	}
	fac, _ := Factorial(x - 1)
	return x * fac, nil
}
