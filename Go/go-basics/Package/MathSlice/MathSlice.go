package mathslice

func SumSlice(s []int) int {
	var res int
	for _, v := range s {
		res += v
	}
	return res
}
