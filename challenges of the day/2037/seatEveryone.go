package main

import (
	"fmt"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	moves := 0

	sort.Slice(seats, func(i, j int) bool {
		return seats[i] < seats[j]
	})
	sort.Slice(students, func(i, j int) bool {
		return students[i] < students[j]
	})

	for i := 0; i < len(seats); i++ {
		movesTmp := seats[i] - students[i]

		if movesTmp < 0 {
			movesTmp *= -1
		}

		moves += movesTmp

	}

	return moves
}

func main() {
	seats := []int{3, 1, 5}
	students := []int{2, 7, 4}

	fmt.Println(minMovesToSeat(seats, students))
}
