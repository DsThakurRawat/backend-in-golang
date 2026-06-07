// Case 1 — & and * (the two operators)
package pointers

import "fmt"

func case1() {
	x := 10
	p := &x // & = "address of": p holds the address of x

	fmt.Println("x  =", x)  // 10
	fmt.Println("p  =", p)  // 0xc0000140a0 (some address)
	fmt.Println("*p =", *p) // 10  — * = "dereference": value at that address

	*p = 20                // store 20 at the address p points to
	fmt.Println("x  =", x) // 20  — x changed, because p pointed at x
}

//Takeaway: & gets an address, * follows it. Writing through *p changes the original.
