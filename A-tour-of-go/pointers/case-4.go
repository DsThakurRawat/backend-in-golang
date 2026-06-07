//Case 4 — pointers with structs (automatic dereference)

package pointers

import "fmt"

type User struct {
    Name string
    Age  int
}

func birthday(u *User) {
    u.Age++ // Go auto-dereferences — no (*u).Age needed
}

func case4() {
    u := User{Name: "Alice", Age: 30}

    birthday(&u)
    fmt.Println(u.Name, u.Age) // Alice 31

    p := &u
    p.Age++    // auto-deref
    (*p).Age++ // explicit — identical effect
    fmt.Println(u.Name, u.Age) // Alice 33
}
// Takeaway: for struct pointers, p.Field and (*p).Field mean the same thing — Go does the deref for you. This is your most common pointer use in real code.