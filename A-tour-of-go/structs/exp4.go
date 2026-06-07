package structs

import "fmt"

type Corners struct {
	X, Y int
}

var (
	v1 = Corners{1, 2}

	v2 = Corners{X: 1}

	v3 = Corners{}
	p  = &Vertex{1, 2}
)

func exp4() {
	fmt.Println(v1, p, v2, v3)
}
