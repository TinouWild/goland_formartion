package basics

import "fmt"

// allocation en un seul
// 1er index toujours à 0
// la taille d'un tableau est définitive
// initialisé avec la valeur a vide du type si aucune valeur

func Tableaux() {
	var names[3]string
	fmt.Printf("names=%v, len=%v\n", names, len(names))
	names[0] = "Bob"
	names[2] = "Alice"
	fmt.Printf("names=%v, len=%v\n", names, len(names))

	odds := [4]int{1, 3, 5, 7}
	fmt.Printf("names=%v, len=%v\n", odds, len(odds))
}