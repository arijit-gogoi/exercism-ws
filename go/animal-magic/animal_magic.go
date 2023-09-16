package chance

import (
	"math/rand"
	"time"
)

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return 1 + r1.Intn(20)
}

// GenerateWandEnergy returns a r1om float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return r1.Float64() * float64(12)
}

// ShuffleAnimals returns a slice with all eight animal strings in r1om order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

	r1.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})

	return animals
}
