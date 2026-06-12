/*
Images

Package image defines the Image interface:

package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}

Note: the Rectangle return value of the Bounds method is actually an image.Rectangle, as the declaration is inside package image.
*/
package interfaces

import (
	"fmt"
	"image"
)

func I13() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
