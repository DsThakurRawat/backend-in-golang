/*
Nil interface values
A nil interface value holds neither value nor concrete type.
Calling a method on a nil interface is a run-time error because there is no type inside
the interface tuple to indicate which concrete method to call.
*/
package interfaces

import "fmt"

type I_6 interface {
	M()
}

func I6() {
	var i I_6 // i = nil (no type, no value)
	describe_6(i) // prints (<nil>, <nil>) — runs fine
	
	// Uncommenting the next line would cause a panic:
	// i.M()  // PANIC — nil pointer, there's no concrete type to call M() on

	// Let's use the second example instead of a panic:
	runExample2_6()
}

func describe_6(i I_6) {
	fmt.Printf("(%v, %T)\n", i, i)
}


// --- Second part of the file ---

// 1. Concrete type that implements I_6
type MyType_6 struct {
	Name string
}

// 2. MyType_6 implements M() so it satisfies interface I_6
func (m MyType_6) M() {
	fmt.Println("M called by:", m.Name)
}

func describe2_6(i I_6) {
	fmt.Printf("(value: %v, type: %T)\n", i, i)
}

func runExample2_6() {
	var i I_6             // i is nil
	i = MyType_6{"Divyansh"} // now i has a concrete type

	describe2_6(i)         // (value: {Divyansh}, type: main.MyType_6)
	i.M()               // M called by: Divyansh
}