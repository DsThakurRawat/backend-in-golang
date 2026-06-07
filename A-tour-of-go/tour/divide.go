package tour
// you have to return quotient and reminder 

// first take input and then 
import(
       "errors"
    
       
)

 func divide(a,b int) (int,int,error ){
	  if b == 0{
		 return 0,0,errors.New("division by zero is not possible")
		 
	  }
	  return a/b,a%b,nil
 }