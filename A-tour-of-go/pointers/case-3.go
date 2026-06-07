//Case 3 — pass by value vs pass by pointer
package pointers

import "fmt"

func tryToZero(x int) {
    x = 0 // only the local copy
}

func actuallyZero(x *int) {
    *x = 0 // the original, via its address
}

func case3() {
    n := 5

    tryToZero(n)
    fmt.Println("after tryToZero:   ", n) // 5  — unchanged

    actuallyZero(&n)
    fmt.Println("after actuallyZero:", n) // 0  — changed
}
//Takeaway: Go copies arguments, so to mutate the caller's variable you pass &n and dereference with *x. This is why pointers exist.