package utils

import (
	"math/rand"
	"reflect"
	"time"
)

var seed = rand.NewSource(time.Now().UnixNano())
var rando = rand.New(seed)

func Roll(count int, die int) int {
	return roll(count, die)
}

func Pick(values interface{}) int {
	val := reflect.ValueOf(values)
	if val.Kind() != reflect.Slice {
		return 1
	}
	return roll(1, val.Len()) - 1
}

func roll(count int, size int) int {
	number := 0
	for d := 0; d < count; d++ {
		die := make([]int, size)

		for i := 0; i < size; i++ {
			die[i] = i + 1
		}
		number += die[rando.Intn(size)]
	}
	return number
}
