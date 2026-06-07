//Case 6 — slices and maps don't need pointers (plus the append trap)
package pointers

import "fmt"

func doubleElements(nums []int) {
    for i := range nums {
        nums[i] *= 2 // caller sees this — no & needed
    }
}

func tryAppend(nums []int) {
    nums = append(nums, 999) // caller may NOT see this
}

func addEntry(m map[string]int) {
    m["added"] = 1 // caller sees this
}

func case6() {
    s := []int{1, 2, 3}

    doubleElements(s)
    fmt.Println("after double:", s) // [2 4 6]

    tryAppend(s)
    fmt.Println("after append:", s) // [2 4 6] — 999 did NOT appear

    m := map[string]int{"start": 0}
    addEntry(m)
    fmt.Println("after addEntry:", m) // map[added:1 start:0]
}

/*
Takeaway: slices and maps already hold an internal pointer, so element edits are visible without &. 
But append can relocate the backing array, so a reassignment inside the function doesn't reach the caller — that's the trap.
*/