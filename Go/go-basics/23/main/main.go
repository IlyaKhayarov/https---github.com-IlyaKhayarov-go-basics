package main

import "fmt"

type MyType int
type IntSlice []int

const (
	MyType1 MyType = 1
	MyType2 MyType = 2
)

func (m MyType) IsValid() bool {
	switch m {
	case MyType1, MyType2:
		return true
	}
	return false
}
func (m MyType) String() string {
	return fmt.Sprintf("Hello from %d and is it valid? %v", m, m.IsValid())
}
func (m *MyType) Set(i int) {
	*m = MyType(i)
}
func main() {
	var m MyType = 2
	m.Set(4)
	fmt.Println(m.String())
	Slice := make(IntSlice, 0)
	Slice.Add(1)
	fmt.Println(Slice)
}
func (s *IntSlice) Add(v int) {
	*s = append(*s, v)
}
