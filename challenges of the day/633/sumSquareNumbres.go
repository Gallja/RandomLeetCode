package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	cnt := 0
	num2 := cnt + 1

	for cnt <= c {
		num2 = cnt

		for num2 <= c {
			if math.Pow(float64(cnt), 2)+math.Pow(float64(num2), 2) == float64(c) {
				return true
			}

			num2++
		}

		cnt++
	}

	return false
}

func judgeSquareSum2(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))

	for left <= right {
		sum := left*left + right*right

		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left++
		}
	}

	return false
}

func main() {
	fmt.Println(judgeSquareSum(5))
	fmt.Println(judgeSquareSum(3))
	fmt.Println(judgeSquareSum(4))
	fmt.Println(judgeSquareSum(2))

	fmt.Println("Versione 2")

	fmt.Println(judgeSquareSum2(5))
	fmt.Println(judgeSquareSum2(3))
	fmt.Println(judgeSquareSum2(4))
	fmt.Println(judgeSquareSum2(2))
}
