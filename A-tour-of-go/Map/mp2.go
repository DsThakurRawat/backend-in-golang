/*
Map literals
Map literals are like struct literals, but the keys are required.
*/
package maps

import "fmt"



var m2 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func mp2() {
	fmt.Println(m2)
}
