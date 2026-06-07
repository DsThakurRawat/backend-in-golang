/*

Map literals continued
If the top-level type is just a type name, 
you can omit it from the elements of the literal.
*/

package maps

import "fmt"



var m3 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func mp3() {
	fmt.Println(m3)
}
