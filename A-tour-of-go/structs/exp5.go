package structs

import "fmt"

/*

Struct fields can be accessed through a struct pointer.
To access the field X of a struct when we have the struct pointer p we could write (*p).X.
However, that notation is cumbersome, so the language permits us instead to write just p.X,
without the explicit dereference.




*/

type exp5 struct {
	X int
	Y int
}

func EXP5() {
	v := exp5{1, 2}
	p := &v //  // p holds the address of v
	p.X = 1e9
	fmt.Println(v)
}
