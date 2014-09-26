package factorial

import "testing"

// func TestFactorial(t *testing.T) {
// 	_, err := Factorial(-321)
// 	outputPrinter(t, 0, err)
// 	_, err = Factorial(200)
// 	outputPrinter(t, 0, err)
// 	factvalue, _ := Factorial(2)

// 	outputPrinter(t, factvalue, nil)
// }

// func outputPrinter(t *testing.T, out float64, err error) {
// 	if err != nil {
// 		t.Error(err.Error())
// 	} else {
// 		t.Log(out)
// 	}
// }

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fact, _ := Factorial(42)
		b.Log(fact)
	}
}
