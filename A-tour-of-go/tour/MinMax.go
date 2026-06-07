package tour
 import (
	"fmt"
	"math"
 )
func findMinAndMax(){

	var n int
    fmt.Scan(&n)  
    nums := make([]int ,n)
	for i := range nums{
		fmt.Scan(&nums[i])
	}
	mini := math.MaxInt
	maxi := math.MinInt 

	for _, p := range nums{
           if p < mini{
			mini = p
		   }
		   if p > maxi{
			maxi = p
		   }
	} 
	fmt.Println(mini,maxi)



}