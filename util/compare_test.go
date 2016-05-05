package util

import "testing"

func TestCompare(t *testing.T) {
    array := [10]byte {
        1, 2, 3, 4, 5, 6, 7, 8, 9,
    }
    data := []struct {
        in [][]byte, want int
    } {
        {{array[0:1], array[0:2}, -1},
        {},
    }

}
