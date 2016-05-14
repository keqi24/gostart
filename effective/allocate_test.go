package effective

import (
    "testing"
    "fmt"
)

func TestNewT(t *testing.T) {
    //s := &[]string {"hello"}
    //fmt.Printf("%T\n", s)

    //different new and make
    var p *[]int = new([]int)  // *p is empty
    var v []int = make([]int, 100) //v is 100 elements slice

    fmt.Println(*p)
    fmt.Println(v)
}
