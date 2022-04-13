package myrepo

import (

	"fmt"

)

//get dec number then divide by base your going into
func DecToBase(dec int, base int) string {
	
	var result string
	const charset = "0123456789ABCDF"
	for dec > 0 {
	
		
		remainder := dec % base
		fmt.Println("val->", remainder)
		
		//convert to string
		result = string(charset[remainder])+ result
		fmt.Println("newReslt", result)
		
		//divide decimal by base
		dec = dec / base
		fmt.Println("newDec", dec)
		fmt.Println("\n")
	}
	


		
	return result
}	


