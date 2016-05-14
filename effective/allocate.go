/*
in go there are make and new
new(T) will return *T which is zeroed
make(T, args) used to create slice, maps and channels and it returns an initialized value of tyep
The reason for the distinction is that these three types must be initialized before use

notice: new(T) and &T{} are equivalent in go
 */
package effective

func NewT() {

}