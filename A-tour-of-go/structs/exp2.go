package structs

import "fmt"

type Edges struct {
	X int
	Y int
}

func exp2() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

}
