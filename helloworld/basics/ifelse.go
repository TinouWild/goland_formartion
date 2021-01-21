package basics

import "fmt"

var Count = 5

func IfElse() {
	if Count > 10 {
		fmt.Printf("We have enough count, got=%d\n", Count)
	} else {
		fmt.Printf("Not enough, got=%d\n", Count)
	}
}