package reversenumber

import "testing"

func TestReverseNumber(t *testing.T) {
	rev, err := ReverseNumber(432423)
	if err != nil {
		t.Error(err.Error(), rev)
	} else if rev != 324234 {
		t.Error("Un Expected Reverse")
	} else {
		t.Log(rev)
	}
}
