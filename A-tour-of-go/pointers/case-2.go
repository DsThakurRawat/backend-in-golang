
//Case 2 — nil and the nil-pointer panic
package pointers

import "fmt"

func case2() {
    var p *int // declared, points nowhere -> nil

    fmt.Println("p =", p)            // <nil>
    fmt.Println("is nil?", p == nil) // true

    if p != nil {
        fmt.Println(*p)
    } else {
        fmt.Println("p points to nothing — skipping deref")
    }
    // Output:
    // p = <nil>
    // is nil? true
    // p points to nothing — skipping deref
}
/*
If you delete the if guard and just write fmt.Println(*p), the program crashes:
panic: runtime error: invalid memory address or nil pointer dereference
Takeaway: a pointer that points nowhere is nil; check != nil before dereferencing.

*/