package util

func Compare(a, b []byte) int {
    for i := 0; i < len(a) && i < len(b); i++ {
        switch {
        case a[i] > b[i]:
            return 1
        case a[i] < b[i]:
            return -1
        }
    }

    if len(a) > len(b) {
        return 1
    }
    if len(a) < len(b) {
        return -1
    }
    return 0
}