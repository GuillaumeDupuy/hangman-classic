package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func isInside(entier *int, tab *[]int) bool {
	for _, elem := range *tab {
		if elem == *entier {
			return true
		}
	}
	return false
}

// printWordProgress prints the progess of finding the word
func printWordProgress(wordToFind *string, toRev *[]int) {
	for index, char := range *wordToFind {
		if isInside(&index, toRev) {
			fmt.Print(string(char))
		} else {
			fmt.Print("_")
		}

		if index != len(*wordToFind)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	fmt.Println()
}

// checkError checks if the error is different from nil otherwise displays error
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

//ReadFile returns an array of string which is the same as the file (line = line)
func ReadFile(Filename string) []string {
	var source []string
	file, _ := os.Open(Filename)      // opens the .txt
	scanner := bufio.NewScanner(file) // scanner scans the file
	scanner.Split(bufio.ScanLines)    // sets-up scanner preference to read the file line-by-line
	for scanner.Scan() {              // loop that performs a line-by-line scan on each new iteration
		if scanner.Text() != "" {
			source = append(source, scanner.Text()) // adds the value of scanner (that contains the characters from StylizedFile) to source
		}
	}
	file.Close() // closes the file
	return source
}

// readFile reads files
func readFile() string {
	// read file words.txt

	// data, err := ioutil.ReadFile("words.txt")
	// checkError(err)
	// return string(data)

	// read any file
	input_files := os.Args[1:]

	// if len(input_files) < 1 {
	// 	fmt.Println("Not detected files.")
	// } else {
	fmt.Println("File_name is : ", input_files[0])
	data, err := ioutil.ReadFile(input_files[0])
	checkError(err)
	return string(data)
	//}
}

func main() {
	// Get hangman
	hangman := ReadFile("hangman.txt")

	// get word - open file to get word
	rand.Seed(time.Now().UnixNano())
	randINT := rand.Intn(len(strings.Split(readFile(), "\n")))
	wordToFind := strings.Split(readFile(), "\n")[randINT]

	// Number of letter to reveal
	reveal := len(wordToFind)/2 - 1

	toRev := []int{}
	if reveal > 0 {
		var randInt int

		for i := 0; i < reveal; i++ {
			randInt = rand.Intn(len(wordToFind))

			if !isInside(&randInt, &toRev) {
				toRev = append(toRev, randInt)
			} else {
				i--
			}
		}
	}

	var myInput string

	fmt.Println("Good luck, you have 10 attempts.")

	printWordProgress(&wordToFind, &toRev)

	attempts := 0
	inside := false

	for attempts != 10 {
		inside = false
		// Get user input
		fmt.Print("Choose: ")
		fmt.Scanln(&myInput)

		// Check is choosen letter is in the word
		for index, char := range wordToFind {
			if strings.EqualFold(strings.ToUpper(string(char)), strings.ToUpper(myInput)) {
				if !isInside(&index, &toRev) {
					toRev = append(toRev, index)
				}
				inside = true
			}
		}

		if !inside {

			fmt.Println("Not present in the word,", 10-attempts, "attempts remaining")

			// Print pendu ici
			for i := (attempts * 7); i < (attempts*7)+7; i++ {
				fmt.Println(hangman[i])
			}
			attempts++
		} else {
			// Print word progess
			printWordProgress(&wordToFind, &toRev)
		}

		// Si mot trouvé
		if len(toRev) == len(wordToFind) {
			fmt.Println("Congrats !")
			break
		}
	}

	if attempts == 10 {
		fmt.Println("Failed")
	}
}
