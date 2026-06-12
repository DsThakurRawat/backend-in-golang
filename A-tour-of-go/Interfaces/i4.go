/*
Interface values with nil underlying values

If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception,
but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)

Note that an interface value that holds a nil concrete value is itself non-nil.
*/
package interfaces

import "fmt"

type I_4 interface{
	M()
}
type T_4 struct{
	S string
}
func (t *T_4) M() {
	if t == nil{
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
func I4()  {
	var i I_4 
	var t *T_4 
	i = t 
	describe_4(i) 
	i.M()
	i = &T_4{"hello"}
	describe_4(i)
	i.M()

}

func describe_4(i I_4) {
	fmt.Printf("(%v, %T)\n", i, i)
}