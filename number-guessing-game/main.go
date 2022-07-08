package main

import (
	"fmt"
	"strconv"
)

//configs
//var LOG_LEVEL int = 3 //couldn't make it work
var DEBUG_FLAG bool = true //change it to false to avoid debug messages
var DIGIT_NUM int = 4      //change it to play the game with 5 characters
//-----//

var secretNumber int                    //init secret number here to avoid declaring it everytime we create a secret number
var secretNumberParsed = make([]int, 0) // parsed secret number for comparison
var previousGuesses = make([]int, 0)    //init slice to keep track of the previous guesses
var guessCount int = 0                  //counter
var appLogger = newLogger()

func main() {
	initLogger()
	initSecret() //initilizes a secret number
	initMenu()   //initilizes the menu function to guess and see the results
}

func initLogger() {
	appLogger.Info().Println("Program Started")
	if DEBUG_FLAG {
		appLogger.Info().Println("Debug mode is ON!")
	} else {
		appLogger.Info().Println("Debug mode is OFF!")
	}
	appLogger.Info().Printf("DIGIT NUM IS: %d", DIGIT_NUM)
}

func initMenu() {
	appLogger.Info().Println("----------------NEW RUN INITIATED----------------")
	menuFlag := true    //menu flag to show menu again and again after every guess
	var randomGuess int //user's guess
	var result []int    //variable to store the results

	for menuFlag { //keep menu on
		appLogger.Info().Println("----------------STARTING THE MENU----------------")
		if guessCount != 0 { //if guess count is 0 it means there are no previous guesses

			fmt.Printf("Your previous guesses are: %v \n", previousGuesses)
		} else {
			appLogger.Info().Println("First run initiated")
		}
		fmt.Println("Please enter a number to check: ")
		fmt.Scanln(&randomGuess) //get input
		appLogger.Info().Printf("User guessed %d", randomGuess)
		if len(strconv.Itoa(randomGuess)) != DIGIT_NUM { //check if input has 4 chars
			appLogger.Warning().Printf("User entered a number with digitc count not equal to %d\n", DIGIT_NUM)
			fmt.Printf("ERROR: Please enter a %d character number!\n", DIGIT_NUM)
			continue //skip loop 1 time
		}
		appLogger.Info().Println("User input is valid.")
		guessCount++ //if input is valid, increase the counter by one
		appLogger.Info().Printf("Increasing guess count by one, current value: %d\n", guessCount)

		previousGuesses = append(previousGuesses, randomGuess) //add it to the previous guesses' slice
		appLogger.Info().Printf("Appending guess to previous guesses slice... current slice:  %v\n", previousGuesses)
		result = checkNumber(randomGuess) //get result, [x y] 1x2 array
		fmt.Println(result)               //print result
		if result[0] == 4 {               //if fist element (1x1) of the result is 4 it means the game ends and user won
			appLogger.Info().Printf("User found the correct secret key in %d tries ...\n", guessCount)
			appLogger.Info().Printf("User won the game, exiting...\n")
			fmt.Printf("You have guessed the number right in %d tries. Congrats!", guessCount)
			fmt.Printf("Your previous guesses were: %v\n", previousGuesses)
			return //exit
		}
		appLogger.Info().Printf("User guessed wrong, initiating another run\n")
	}
}

func checkNumber(guess int) []int {
	parsedGuess := parseDigits(guess) //get user guess and parse it to its digits
	appLogger.Info().Printf("Parsing the user guess: %v \n", parsedGuess)
	returnSlice := make([]int, 0) //init return slice
	positive, negative := 0, 0    //init positive and negative to return

	for i, item := range parsedGuess {
		if item == secretNumberParsed[i] { //if index and number equals to each other, increase positive by one
			appLogger.Info().Printf("User input: (%v) and secret key: (%v) first digits (%d) are equal, increasing the positive count by one... current positive %d\n", parsedGuess, secretNumberParsed, secretNumberParsed[i], positive)
			positive++
		} else if containsInt(secretNumberParsed, item) { //else if items are equal to each other but not their index, increase negative by one
			appLogger.Info().Printf("Secret key contains %dth item: %d  in a different position (%v), increasing the negative count by one... current negative %d\n", i, item, secretNumberParsed, negative)
			negative--
		}
	}
	returnSlice = append(returnSlice, positive)
	returnSlice = append(returnSlice, negative)
	appLogger.Info().Printf("Returning postive and negative values %v \n", returnSlice)
	return returnSlice //return pos and neg
}

func initSecret() {
	secretNumber = generateRandomNumberNonRepeating() //initlize a number between 1000-9999 without any repeating chars
	appLogger.Info().Printf("Secret number generated:  %d \n", secretNumber)
	secretNumberParsed = parseDigits(secretNumber) //parse it to compare
	appLogger.Info().Printf("Secret number parsed:  %v \n", secretNumberParsed)
	if DEBUG_FLAG {
		fmt.Printf("DEBUG: Secret number is: %d \n", secretNumber) //check if your secret number is ok
	}
	if DEBUG_FLAG {
		fmt.Printf("DEBUG: Secret number parsed: %v \n", secretNumberParsed) //check if it's parsed correctly
	}
}
