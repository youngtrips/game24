package main

import (
	"fmt"
	"game24/calculator"
)

/*
 3 3 6 6
 3 5 6 6
 3 2 3 4
 7 5 11 13
*/
func main() {
	nums := make([]int, 4)
	fmt.Scanf("%d %d %d %d\n", &nums[0], &nums[1], &nums[2], &nums[3])

	c := calculator.New(nums, 24)
	if exps := c.Calc(); len(exps) > 0 {
		for _, exp := range exps {
			fmt.Println(exp)
		}
	} else {
		fmt.Println("no found.")
	}
}
