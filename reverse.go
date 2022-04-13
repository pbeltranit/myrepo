package reverse

import (

	"fmt"
	//"strings"
)

func Rev2(word string) string {
	var flipped string
	for idx := len(word)-1; idx >= 0; idx-- {
	
		flipped += string(word[idx])
	}
	return flipped

}


func Rev(word string) string {
	var flipped string
	//var sb strings.Builder
	for _,val := range word {
		flipped = string(val) + flipped
		fmt.Println(val)
		//sb.WriteByte(word[i])
		
	}
	return flipped
	//return sb.string
}

