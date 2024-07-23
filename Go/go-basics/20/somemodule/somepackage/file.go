package somepackage

import (
	"fmt"
	"reflect"
)

type Empty struct{}

func Func() {
	fmt.Printf("Hello from: %v\n", reflect.TypeOf(Empty{}).PkgPath())
}
