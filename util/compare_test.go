package util

import (
    "testing"
)

func TestCompare(t *testing.T) {
    array := [10]byte {
        1, 2, 3, 4, 5, 6, 7, 8, 9,
    }
    //data := []struct {
    //    in byte, want int
    //} {
    //    {1, 2},
    //}
    //cases := []struct {
    //    in, want string
    //}{
    //    {"Hello, world", "dlrow ,olleH"},
    //    {"Hello, 世界", "界世 ,olleH"},
    //    {"", ""},
    //}

    //test := [][]byte {
    //    array[0:1], array[1:2],
    //}
    //
    //fmt.Println(test)

    data := []struct {
        in [][]byte
        want int
    } {
        {[][]byte{array[0:1], array[0:2],}, 1},
    }
    //fmt.Println(cases)
    //data := []struct {
    //    in byte, want int
    //}{
    //    {1, 2},
    //}
    for _, c := range data {
        got := Compare(c.in[0], c.in[1])
        if got != c.want {
            t.Errorf("Compare (%v, %v) == %v, want %v", c.in[0], c.in[1], got, c.want)
        }
    }
}
