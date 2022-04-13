package basetodec


import (


	"fmt"


)




func BaseToDec(base int, val string) int {


	const charset = "0123456789ABCDEF"
	total := 0
	mult := 1
	
	//reverse string
	for idx := len(val) - 1; idx >= 0; idx-- {
	
		chrindex := -1
		
		//map char to int using a custom charset and its index
		for j, char := range charset {
		
			if char == rune(val[idx]) {
				chrindex = j
				break
			}
		}
		

		
		total += mult*chrindex
		
		mult *= base
		
		

	}
	return total
}


