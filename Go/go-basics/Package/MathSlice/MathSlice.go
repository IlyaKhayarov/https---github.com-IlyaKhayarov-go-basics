package mathslice

type Slice []El
type El int

func SumSlice(s Slice) (res El) {
	for _, v := range s {
		res += v
	}
	return res
}

func MapSlice(s Slice, val int, op func(El, int) El) {
	for i, v := range s {
		s[i] = op(v, val)
	}
}
