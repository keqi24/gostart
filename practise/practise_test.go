package tmp

import (
    "testing"
    "fmt"
)

func TestPAppend(t *testing.T) {
    //append elements
    s := []int {1, 2, 3,}
    x := append(s, 4, 5, 6)
    fmt.Println(x)

    //append slice to slice
    y := []int {4, 5, 6}
    x = append(s, y...)
    fmt.Println(x)
}

func TestPTwoInterface(t *testing.T) {
    PTwoInterface()
}

func TestPMulMethodInt(t *testing.T) {
    PMulMethodInt()
}

func TestPPointerVsValue(t *testing.T) {
    PPointerVsValue()
}

func TestPPinterMethod(t *testing.T) {
    PPinterMethod()
}

func TestPType(t *testing.T) {
    PType()
}

func TestPEmbed(t *testing.T) {
    PEmbed()
}

func TestPChannel(t *testing.T) {
    PChannel()
}

func TestPrpcServer(t *testing.T) {
    PrpcServer()
}

func TestPFlag(t *testing.T) {
    PFlag()
}