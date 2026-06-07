//Case 7 — returning a pointer to a local (safe in Go)

package pointers

import "fmt"

type People struct {
	Name string
}

func newUser(name string) *People {
	u := People{Name: name} // local variable
	return &u               // safe — escape analysis keeps it alive
}

func case7() {
	a := newUser("Alice")
	b := newUser("Bob")

	fmt.Println(a.Name, b.Name) // Alice Bob
	fmt.Println(a == b)         // false — distinct objects
}
//Takeaway: unlike C, returning the address of a local is fine. 
// The compiler detects the pointer escapes and keeps the data alive; 
// the GC frees it later. No manual alloc/free, no pointer arithmetic.