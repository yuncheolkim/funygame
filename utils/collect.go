package utils

import (
	"math/rand"
	"time"
)

func Shuffle(slice []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}
func DeleteSlice(a []int, v int) []int {
	j := 0
	for i, val := range a {
		if val == v {
			j = i
			break
		}
	}
	return append(a[:j], a[j+1:]...)
}