package main

import (
	"fmt"
	"goland_formartion/helloworld/basics"
)

func main() {
	fmt.Println("variable exportee", basics.Name)

	// On déclare lowerName et l
	lowerName, l := basics.ToLoverStr("ALICE")
	fmt.Printf("%s (len=%d)\n", lowerName, l)
	// GO est sympa, la déclaration sur l est ignorée car bobLower est là
	bobLower, l := basics.ToLoverStr("BOB")
	fmt.Printf("%s (len=%d)\n", bobLower, l)

	basics.Tableaux()
	basics.Slices()

	data, err := basics.ReadFile("basics/test.txt")
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}
	fmt.Println("File content :")
	fmt.Println(data)
}
