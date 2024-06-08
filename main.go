package main

import (
	"fmt"
	a1 "geekhunter/answer1"
	a2 "geekhunter/answer2"
	a3 "geekhunter/answer3"
)

func main() {
	fmt.Println("----------------------------------------")
	fmt.Println("--------------- ANSWER 1 ---------------")
	fmt.Println("----------------------------------------")
	s := "abbcccd"
	queries := []int{1, 3, 9, 8}
	results := a1.WeightedStrings(s, queries)
	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Println("----------------------------------------")
	fmt.Println("--------------- ANSWER 2 ---------------")
	fmt.Println("----------------------------------------")
	fmt.Println(a2.IsBalanced("{[()]}"))
	fmt.Println(a2.IsBalanced("{[(])}"))
	fmt.Println(a2.IsBalanced("{(([])[[]])[]}"))
	fmt.Println("----------------------------------------")
	fmt.Println("--------------- ANSWER 3 ---------------")
	fmt.Println("----------------------------------------")
	fmt.Println(a3.HighestPalindrome("3943", 1))
	fmt.Println(a3.HighestPalindrome("932239", 2))
	fmt.Println(a3.HighestPalindrome("12321", 1))
	fmt.Println(a3.HighestPalindrome("12345", 1))
}
