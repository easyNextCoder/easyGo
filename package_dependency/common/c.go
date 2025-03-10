package common

import "fmt"

type C struct {
	val int
}

func init() {
	fmt.Println("c init")
}
