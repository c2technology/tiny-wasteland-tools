package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/c2technology/tiny-wasteland-tools/utils"
)

// Rolls N X-sided dice
func main() {
	input := os.Args[1:][0]

	if !strings.Contains(input, "d") {
		help()
		return
	}
	s := strings.Split(input, "d")
	if len(s) != 2 {
		help()
		return
	}
	count, err := strconv.Atoi(s[0])
	if err != nil {
		panic(err)
	}
	sides, err := strconv.Atoi(s[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(utils.Roll(count, sides))
}

func help() {
	fmt.Println("Roll an X N-sided dice")
	fmt.Println("2 6-sided dice: 2d6")
	fmt.Println("20 42-sided dice: 20d42")
}
