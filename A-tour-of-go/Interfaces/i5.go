/*
Interface values with nil underlying values

If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception, 
but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)

Note that an interface value that holds a nil concrete value is itself non-nil.
*/
package interfaces

import "fmt"

type I_5 interface {
	M()
}

type T_5 struct {
	S string
}

func (t *T_5) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func I5() {
	var i I_5
	var t *T_5
	i = t
	describe_5(i)
	i.M()

	i = &T_5{"hello"}
	describe_5(i)
	i.M()
}

func describe_5(i I_5) {
	fmt.Printf("(%v, %T)\n", i, i)
}