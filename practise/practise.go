package tmp

import (
    "fmt"
    "log"
    "os"
    "github.com/keqi24/gostart/util"
    "time"
    "flag"
)

/*
 each source file can define it own init(it can own multiple init)
 init is called after all the variable decalration

 */
var (
    home   = os.Getenv("HOME")
    user   = os.Getenv("USER")
    gopath = os.Getenv("GOPATH")
)

func init() {
    if user == "" {
        log.Fatal("$USER not set")
    }
    if home == "" {
        home = "/home/" + user
    }
    if gopath == "" {
        gopath = home + "/go"
    }
    // gopath may be overridden by --gopath flag on command line.
    //flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
    fmt.Println(user, home, gopath)
}

func init()  {
    fmt.Println("I am second init")
}

func PAppend() {

}


/*
 one method implement two interface
 */
type inA interface {
    Hello()
}

type inB interface  {
    Hello()
}

type child int16

func (v child)Hello() {
    fmt.Println("hello")
}

func testAHello(v inA)  {
    v.Hello()
}

func testBHello(v inB)  {
    v.Hello()
}

func PTwoInterface()  {
    var c child
    testAHello(c)
    testBHello(c)
}

/*
 implement part method of interface
 if want to cast to interface, child must implement all the methods declared by parent interface
 */

type mulMethodInt interface  {
    HelloA()
    HelloB()
}

type child2 struct {
    a int
    b int
}

func (c child2)HelloA() {
    fmt.Println("Hello A")
}

func (c child2)HelloB() {

}

func useMulMethodInt(m mulMethodInt)  {
    m.HelloA()
}

func PMulMethodInt()  {
    var c child2
    useMulMethodInt(c)
}

/*
 Pointer vs Value receiver
 The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values,
 but pointer methods can only be invoked on pointers. (this rule used to when implemented a interface)
 */

type byteSlice []byte

func (b *byteSlice)append1(s []byte)  {
    *b = append(*b, s...)
}

func (b byteSlice)append2(s []byte) (newByte []byte) {
    newByte = append(b, s...)
    return
}

func PPointerVsValue()  {
    bs := byteSlice {10, 12,}
    p := &bs
    s := []byte("1234")
    p.append1(s)
    fmt.Println("pointer:", p)

    newB := bs.append2(s)
    fmt.Println("value append2: ", newB)
    newB = p.append2(s)
    fmt.Println("pointer append2:", newB)
}

type testInt interface {
    Test()
}

type str struct {
    a int
    b string
}

func (p *str)Test()  {
    fmt.Println(p.b)
}

func useTestInt(t testInt)  {
    t.Test()
}

func PPinterMethod() {
    //s := str{1, "kk"}
    var s str
    s.b = "kk"
    //this is ok
    s.Test()
    //error here
    //useTestInt(s)
    //should use
    useTestInt(&s)
}

/*
    .(type)
 */
func PType() {
    var st str
    var val interface{}
    val = st
    switch s := val.(type) {
    case string:
        fmt.Println("case string:", s)
    case str:
        fmt.Println("case str", s)
    }

    data, ok := val.(string)
    if ok {
        fmt.Println(data)
    } else {
        fmt.Println("no string ")
    }
}

/*
 use _ to ignore unused import and variable
 sometimes this is useful when you just want to use the package's side effect (like things done by it's init function)

 */

var _ = util.Compare

func testUnusedVariable() {
    var a int
    //TODO use a
    _ = a
}

/*
 interface check
 some time we need to guarantee the implementation is correct can user black identifier to static check interface
 */

type needInt interface {
    Hello()
}

type child3 int

func (c child3)Hello() {

}

var _ needInt = (*child3)(nil)


/*
 test embedding
 */
type outer struct {
    status string
    *util.Embed
}

func PEmbed() {
    o := outer{"status", &util.Embed{Name:"embed"}}
    //o.SayHello()

    //var o outer
    fmt.Println(o.Name)
    o.SetName("outer")
    o.SayHello()
    o.SetPrivateParam("private")
    fmt.Println(o.PrivateParam())
}

/*
 channel
 use channel to control goroutine count
 */



type Request struct  {
    name string
    param string
}

func (r *Request)String() string {
    return fmt.Sprintf("%v : %v", r.name, r.param)
}

func Serve(queue chan *Request)  {
    var sem = make(chan int, 3)
    for req := range queue  {
        r := req
        sem <- 1
        go func() {
            fmt.Println("start to handle:", r)
            time.Sleep(10 * time.Second)
            fmt.Println("finish:", r)
            <-sem
        }()
    }
}

func PChannel() {

    //queue := make(chan *Request)
    //
    //go func() {
    //    for i := 0; i < 10; i++ {
    //        queue <- &Request{"hello", fmt.Sprintf("param:%d", i)}
    //    }
    //    close(queue)
    //}()
    //
    //Serve(queue)

}

/*
 Basic framework for RPC
 */

type rpcRequest struct {
    args []int
    f func([]int)int
    resultChan chan int
}

func sum(data []int) int {
    sum := 0
    for i := range data {
        sum += i
    }
    return sum
}

func client(queue chan *rpcRequest)  {
    req := &rpcRequest{[]int{1, 2, 3, 7, 8, 9}, sum, make(chan int)}
    queue <- req
    fmt.Println("result=", <-req.resultChan)
}

func server(queue chan *rpcRequest) {
    req := <- queue
    sum := req.f(req.args)
    req.resultChan <- sum
}

func PrpcServer() {
    queue := make(chan *rpcRequest)
    go server(queue)
    client(queue)
}


/*
 test flag
 */
var data = flag.String("name", "hello", "some thing else")

func PFlag() {
    flag.Parse()
    fmt.Println(data)
}