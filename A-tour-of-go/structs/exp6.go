package structs

import "fmt"

/*
A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

The special prefix & returns a pointer to the struct value.
*/
type exp6 struct {
	X, Y int
}

var (
	p1 = exp6{1, 2}  // has type exp6
	p2 = exp6{X: 1}  //Y:0 is implicit
	p3 = exp6{}      // X:0 and Y:0
	q  = &exp6{1, 2} // has type *Vertex

)

func Exp6() {
	fmt.Println(p1, q, p2, p3)
}
