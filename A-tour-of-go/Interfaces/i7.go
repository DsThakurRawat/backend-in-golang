package interfaces

import "fmt"

func I7(){
	var i interface{}
	describe_7(i)

	i = 42
	 describe_7(i)

}
func describe_7(i interface{}){
	fmt.Println("(%v,%T)\n",i,i)
}