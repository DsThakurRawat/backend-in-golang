/*


Interfaces are implemented implicitly
A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.




*/

package interfaces

import "fmt"

type I_2 interface{
	M()
}

type T_2 struct{
	S string
}

//// This method means type T_2 implements the interface I_2,
////// but we don't need to explicitly declare that it does so

func (t T_2 ) M(){
	fmt.Println(t.S)
}
func I2(){
	var i I_2 = T_2{"hello"}
	i.M()
}