package a

import (
	b "easyGo/package_dependency/common/b"
	"fmt"
)

func init() {
	fmt.Println("a init")
	b.B()
}

func A() {
	fmt.Println("A do")
}
