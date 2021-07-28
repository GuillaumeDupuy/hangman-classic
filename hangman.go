package main

import (
	"fmt"
	"math/rand"
	"time"
)

func isRevealOK(toReveal *[]int, random *int) bool {

	for elem := range *toReveal {
		if elem == *random {
			return false
		}
	}

	return true
}

func main() {
	// get word - open file to get word
	word := "HELLO"

	// Number of letter to reveal
	reveal := len(word)/2 - 1

	rand.Seed(time.Now().UnixNano())

	if reveal > 0 {
		toRev := []int{}
		var randInt int

		for i := 0; i < reveal; i++ {
			randInt = rand.Intn(len(word))

			for elem := range toRev {
				if elem == randInt {

				}
				toRev = append(toRev)
			}
		}
		fmt.Println(toRev)
	}

	var myInput string

	fmt.Println("Good luck, you have 10 attempts.")

	// for index, char := range word {
	// 	if index ==
	// }

	for attempts := 0; attempts < 10; attempts++ {
		fmt.Print("Choose :")
		fmt.Scanln(&myInput)
		fmt.Println(attempts, word, reveal, myInput)
	}
}
