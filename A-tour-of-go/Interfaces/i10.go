/*
Stringers

One of the most ubiquitous interfaces is Stringer defined by the fmt package.

type Stringer interface {
    String() string
}

A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
*/
package interfaces

import "fmt"

type Person_10 struct {
	Name string
	Age  int
}

func (p Person_10) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func I10() {
	a := Person_10{"Arthur Dent", 42}
	z := Person_10{"Zaphod Beeblebrox", 9001}
	fmt.Println(a)
	fmt.Println(z)
}
