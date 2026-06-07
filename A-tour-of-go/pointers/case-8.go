//Case 8 — when a pointer earns its place (nil = "absent")

package pointers

import "fmt"

type Config struct {
    Timeout int
}

// *Config lets us signal "not found" with nil,
// which is distinct from a zeroed Config{Timeout: 0}.
func loadConfig(found bool) *Config {
    if !found {
        return nil
    }
    return &Config{Timeout: 30}
}

func case8() {
    c := loadConfig(false)
    if c == nil {
        fmt.Println("no config — using defaults")
    }

    c = loadConfig(true)
    if c != nil {
        fmt.Println("timeout:", c.Timeout) // 30
    }
    // Output:
    // no config — using defaults
    // timeout: 30
}
//Takeaway: the three reasons to use a pointer, 
// all shown across these examples — modify the caller's value (Case 3, 4),
//  avoid copying a big struct (Case 4), or represent "absent/nil" as a real state (Case 8). 
// Otherwise prefer plain values.
