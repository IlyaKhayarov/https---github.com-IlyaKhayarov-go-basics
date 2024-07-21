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

func FoldSlice(s Slice, op func(El, El) El, init El) (res El) {
	for i, v := range s {
		switch i {
		case 0:
			res = op(init, s[0])
		default:
			res = op(res, v)
		}
	}
	return res
}
