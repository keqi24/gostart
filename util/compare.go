package util

import "fmt"

type Embed struct  {
    Name string
    privateParam string
}

func (e *Embed)SayHello() {
    fmt.Println(e.Name, "say hello")
}

func (e *Embed)Hello() {
    fmt.Println("hello")
}

func (e *Embed)SetName(name string)  {
    e.Name = name
}

func (e *Embed)SetPrivateParam(str string)  {
    e.privateParam = str
}

func (e *Embed)PrivateParam() string {
    return e.privateParam
}


func TakePlace() {

}

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