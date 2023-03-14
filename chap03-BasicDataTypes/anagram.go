package main 

import "fmt"
import "os"

func main() {
	argLength := len(os.Args)
	if  argLength < 3 || argLength > 3 {
		panic("Just need 2 strings to compare")
	}

	string1, string2 := os.Args[1], os.Args[2]
	strlength1 := len(string1)
	strlength2 := len(string2)

	if strlength2 != strlength1 {
		fmt.Println("Not an anagram")
	} else {
		for i,j := 0, strlength2-1; i < strlength1 && j >= 0; i, j  = i+1, j-1 {
			if string1[i] !string2 string2[j] {
				fmt.Println("Not an anagram")
			} else {
				fmt.Println("Anagram")
			}
 		}
	}

}
