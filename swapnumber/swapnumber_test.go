package swapnumber

import "testing"

func TestSwap(t *testing.T) {
	var v float64
	first, second, err := Swap(543, 8)
	if err != nil {
		t.Error(err.Error(), v)
	} else if first != 8 || second != 543 {
		t.Error("Un Expected Swap")
	}
}
