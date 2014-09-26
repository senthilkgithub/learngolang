package swapnumber

import (
	"errors"
	"reflect"
)

func Swap(first, second float64) (float64, float64, error) {
	if reflect.ValueOf(first).Kind() == reflect.Float64 && reflect.ValueOf(second).Kind() == reflect.Float64 {
		first = first + second
		second = first - second
		first = first - second
		return first, second, nil
	} else {
		return 0, 0, errors.New("Input type mismatch")
	}
}
