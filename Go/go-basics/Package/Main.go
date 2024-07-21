package main

import (
	"fmt"
	mathslice "go/go-basics/Package/MathSlice"
)

func main() {
	testSlice := mathslice.Slice{2, 4, 3}
	fmt.Printf("Сумма слайса %v: = %d\n", testSlice, mathslice.SumSlice(testSlice))

	testVal := 3
	copiedSlice := make([]mathslice.El, len(testSlice))
	copy(copiedSlice, testSlice)

	fmt.Printf("Умножение слайса %v на %v ", testSlice, testVal)
	mathslice.MapSlice(testSlice, testVal, func(i mathslice.El, val int) mathslice.El {
		return i * mathslice.El(testVal)
	})
	fmt.Printf("Равно %v\n", testSlice)

	fmt.Printf("Свёртка слайса %v умножением = %v\n", copiedSlice, mathslice.FoldSlice(copiedSlice, func(a mathslice.El, b mathslice.El) mathslice.El {
		return a * b
	}, 1))
	fmt.Printf("Свёртка слайса %v сложением = %v\n", copiedSlice, mathslice.FoldSlice(copiedSlice, func(a mathslice.El, b mathslice.El) mathslice.El {
		return a + b
	}, 0))
}
