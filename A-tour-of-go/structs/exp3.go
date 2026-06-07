package structs

import "fmt"

type Vertices struct {
	X int
	Y int
}

func exp3() {
	v := Vertices{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

}
