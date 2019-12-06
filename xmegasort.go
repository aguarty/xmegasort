package xmegasort

import (
	"math/rand"
	"time"
)

// CreateSlice create slice
func CreateTestSlice(count int) []int {
	s := make([]int, 0)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < count; i++ {
		s = append(s, r1.Intn(9999)+1)
	}
	return s
}
