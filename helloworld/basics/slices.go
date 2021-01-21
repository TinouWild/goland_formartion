package basics

import "fmt"

// A RETENIR: si len du slice > cap du tableau initial, cap x 2
func Slices() {
	nums := make([]int, 2, 3)
	nums[0] = 10
	nums[1] = -1
	fmt.Printf("%v (len=%v, cap=%v)\n", nums, len(nums), cap(nums))
	nums = append(nums, 64)
	fmt.Printf("%v (len=%v, cap=%v)\n", nums, len(nums), cap(nums))
	nums = append(nums, -8)
	fmt.Printf("%v (len=%v, cap=%v)\n", nums, len(nums), cap(nums))

	letters := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Printf("%v (len=%v, cap=%v)\n", letters, len(letters), cap(letters))

	sub1 := letters[:2]
	sub2 := letters[2:]
	fmt.Printf("%v (len=%v, cap=%v)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("%v (len=%v, cap=%v)\n", sub2, len(sub2), cap(sub2))
}