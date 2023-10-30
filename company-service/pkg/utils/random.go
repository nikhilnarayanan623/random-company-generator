package utils

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.NewString()
}

// To get random index. out will within(0 to len-1)
func GetRandomIndex(len int) int {
	return rand.Intn(len)
}

// To get an int between start and end(start and end is inclusive)
func GetIntBetween(start, end int) int {

	diff := (end - start) + 1

	return start + (rand.Intn(diff))
}

// To get a random time between start and end
func GetTimeBetween(start, end time.Time) time.Time {
	// get the difference of time duration between start and end
	diff := end.Sub(start)

	// select and random duration and add it to start.
	return start.Add(time.Duration(rand.Intn(int(diff))))
}

func GetFloatBetween(start, end float64) float64 {

	diff := (end - start) + 1

	return start + (rand.Float64() * diff)
}
