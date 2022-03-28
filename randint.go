package randint

import (
	"math/rand"
	"time"
)

func Randint(min int, max int) int {
	max++
	max = max - min
	rand.Seed(time.Now().UnixNano())
	r := rand.Int()
	return (r % max) + min
}
