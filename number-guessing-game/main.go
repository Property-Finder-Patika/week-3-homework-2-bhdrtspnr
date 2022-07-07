package main

import (
	"fmt"
	"strconv"
)

var secretNumber int                    //init secret number here to avoid declaring it everytime we create a secret number
var secretNumberParsed = make([]int, 0) // parsed secret number for comparison
var previousGuesses = make([]int, 0)    //init slice to keep track of the previous guesses
var guessCount int = 0                  //counter

func main() {
	initSecret() //initilizes a secret number
	initMenu()   //initilizes the menu function to guess and see the results
}

func initMenu() {
	menuFlag := true    //menu flag to show menu again and again after every guess
	var randomGuess int //user's guess
	var result []int    //variable to store the results

	for menuFlag { //keep menu on
		if guessCount != 0 { //if guess count is 0 it means there are no previous guesses
			fmt.Printf("Your previous guesses are: %v \n", previousGuesses)
		}
		fmt.Println("Please enter a number to check: ")
		fmt.Scanln(&randomGuess)                         //get input
		if len(strconv.Itoa(randomGuess)) != DIGIT_NUM { //check if input has 4 chars
			fmt.Printf("ERROR: Please enter a %d character number!\n", DIGIT_NUM)
			continue //skip loop 1 time
		}
		guessCount++                                           //if input is valid, increase the counter by one
		previousGuesses = append(previousGuesses, randomGuess) //add it to the previous guesses' slice
		result = checkNumber(randomGuess)                      //get result, [x y] 1x2 array
		fmt.Println(result)                                    //print result
		if result[0] == 4 {                                    //if fist element (1x1) of the result is 4 it means the game ends and user won
			fmt.Printf("You have guessed the number right in %d tries. Congrats!", guessCount)
			fmt.Printf("Your previous guesses were: %v\n", previousGuesses)
			return //exit
		}
	}
}

func checkNumber(guess int) []int {
	parsedGuess := parseDigits(guess) //get user guess and parse it to its digits
	returnSlice := make([]int, 0)     //init return slice
	positive, negative := 0, 0        //init positive and negative to return

	for i, item := range parsedGuess {
		if item == secretNumberParsed[i] { //if index and number equals to each other, increase positive by one
			positive++
		} else if containsInt(secretNumberParsed, item) { //else if items are equal to each other but not their index, increase negative by one
			negative--
		}
	}
	returnSlice = append(returnSlice, positive)
	returnSlice = append(returnSlice, negative)
	return returnSlice //return pos and neg
}

func initSecret() {
	secretNumber = generateRandomNumberNonRepeating() //initlize a number between 1000-9999 without any repeating chars
	secretNumberParsed = parseDigits(secretNumber)    //parse it to compare
	if DEBUG_FLAG {
		fmt.Printf("DEBUG: Secret number is: %d \n", secretNumber) //check if your secret number is ok
	}
	if DEBUG_FLAG {
		fmt.Printf("DEBUG: Secret number parsed: %v \n", secretNumberParsed) //check if it's parsed correctly
	}
}
