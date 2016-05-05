package main

import (
	"fmt"
	"github.com/keqi24/gostart/util"
)

func main() {
	fmt.Println("Hello Go!")
	a := []byte {
		10,
		5,
		4,
		10,
	}
	b := []byte {
		10,
		5,
		4,
	}
	fmt.Println(util.Compare(a, b))
}
