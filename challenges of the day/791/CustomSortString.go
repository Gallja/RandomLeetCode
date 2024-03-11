package main

import (
	"fmt"
)

/*
You are given two strings order and s. All the characters of order are unique and were sorted in some custom order previously.

Permute the characters of s so that they match the order that order was sorted. More specifically, if a character x occurs before a character y in order, then x should occur before y in the permuted string.

Return any permutation of s that satisfies this property.
*/
func customSortString(order string, s string) string {
	strOut := ""

	for i := 0; i < len(order); i++ {
		for j := 0; j < len(s); j++ {
			if order[i] == s[j] {
				strOut += string(order[i])
			}
		}
	}

	strOut += checkS(s, order)

	return strOut
}

/*
This function check if there are different characters between str1 and str2.
*/
func checkS(str1, str2 string) string {
	isIn := false
	strOut := ""

	for i := 0; i < len(str1); i++ {
		isIn = false
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				isIn = true
			}
		}
		if !isIn {
			strOut += string(str1[i])
		}
	}

	return strOut
}

func main() {
	fmt.Println(customSortString("cba", "cbad"))

	fmt.Println(customSortString("bcafg", "abcd"))
}
