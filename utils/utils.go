package utils

import (
	"math/rand"
	"reflect"
	"time"
)

//Result of rolls
type Result struct {
	Sum   int
	Rolls []int
}

var seed = rand.NewSource(time.Now().UnixNano())
var rando = rand.New(seed)

//Roll number of sided dice
func Roll(number int, sided int) Result {
	return roll(number, sided)
}

//Pick a value from given values returning the index of the given values
func Pick(values interface{}) int {
	val := reflect.ValueOf(values)
	if val.Kind() != reflect.Slice {
		return 1
	}
	return roll(1, val.Len()).Sum - 1
}

func roll(count int, size int) Result {
	result := Result{}
	for d := 0; d < count; d++ {
		die := make([]int, size)

		for i := 0; i < size; i++ {
			die[i] = i + 1
		}
		roll := die[rando.Intn(size)]
		result.Rolls = append(result.Rolls, roll)
		result.Sum += roll
	}
	return result
}
