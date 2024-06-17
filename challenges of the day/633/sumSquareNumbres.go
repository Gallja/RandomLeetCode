package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	squareOk := false
	cnt := 1
	num2 := cnt + 1

	for cnt < c {
		num2 = cnt + 1

		for num2 < c {
			if math.Pow(float64(cnt), 2)+math.Pow(float64(num2), 2) == float64(c) {
				return true
			}
			// fmt.Println(math.Pow(float64(cnt), 2) + math.Pow(float64(num2), 2))

			num2++
		}

		cnt++
	}

	return squareOk
}

func main() {
	fmt.Println(judgeSquareSum(5))
}
