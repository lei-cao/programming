package utils

import (
	"math/rand"
	"time"
)

func Shuffle(size int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := []int{}
	for i := 0; i < size; i++ {
		numbers = append(numbers, i+1)

	}
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	return numbers
}