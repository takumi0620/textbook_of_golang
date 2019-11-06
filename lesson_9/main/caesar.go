package main

import "fmt"

func main() {
	question := "L fdph, L vdz, L frqtxhuhg."
	for i := 0; i < len(question); i++ {
		if (question[i] >= 'a' && question[i] <= 'z') || (question[i] >= 'A' && question[i] <= 'Z') {
			fmt.Printf("%c", question[i]-3)
		} else {
			fmt.Printf("%c", question[i])
		}

	}
}
