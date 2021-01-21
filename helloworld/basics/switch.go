package basics

import "fmt"

func Main() {
	weekDay := 1
	fmt.Printf("day = %d\n", weekDay)

	switch weekDay {
	case 1:
		fmt.Println("C'est lundi !")
	case 5:
		fmt.Println("C'est vendredi !")
	case 6, 7:
		fmt.Println("C'est le week-end !")
	default:
		fmt.Println("On s'enfout du jour qu'on est")
	}
}