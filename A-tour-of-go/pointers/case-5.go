//Case 5 — creating pointers: &T{} vs new
package pointers

import "fmt"

type Person struct {
    Name string
    Age  int
}

func case5() {
    // Idiomatic: &T{} makes a pointer AND sets fields at once
    u := &Person{Name: "Bob", Age: 25}
    fmt.Println(u)      // &{Bob 25}
    fmt.Println(u.Name) // Bob

    // new(T): allocates a zeroed T, returns *T
    p := new(int)
    fmt.Println(*p)     // 0
    *p = 42
    fmt.Println(*p)     // 42

    u2 := new(Person)     // zeroed struct, then fill in
    u2.Name = "Carol"
    fmt.Println(u2)     // &{Carol 0}
}

//Takeaway: reach for &User{...} (flexible, set fields inline); you'll read new occasionally but rarely write it.