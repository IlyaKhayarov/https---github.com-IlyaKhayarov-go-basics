package main

import (
	"fmt"
	mathslice "go/go-basics/Package/MathSlice"
)

func main() {
	testSlice := mathslice.Slice{2, 4}
	fmt.Printf("Сумма слайса %v: = %d\n", testSlice, mathslice.SumSlice(testSlice))

	testVal := 3

	fmt.Printf("Умножение слайса %v на %v\n", testSlice, testVal)
	mathslice.MapSlice(testSlice, testVal, func(i mathslice.El, val int) mathslice.El {
		return i * mathslice.El(testVal)
	})
	fmt.Printf("Равно %v\n", testSlice)
}
