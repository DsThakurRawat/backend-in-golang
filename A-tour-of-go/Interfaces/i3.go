/*
Interface values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type
(value, type)
An interface value holds a value of a specific underlying concrete type.
Calling a method on an interface value executes the method of the same name on its underlying type.
*/
package interfaces

import (
	"fmt"
)

type I_3 interface {
	M()
}
type T_3 struct {
	S string
}

func (t *T_3) M() {
	fmt.Println(t.S)
}

type F_3 float64

func (f F_3) M() {
	fmt.Println((f))
}

func I3() {
	var i I_3
	i = &T_3{"Hello"}
	describe_3(i)
	i.M()

}
func describe_3(i I_3) {
	fmt.Println("(%v,%T)\n", i, i)
}
