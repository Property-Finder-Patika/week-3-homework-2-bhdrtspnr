package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
)

var DEBUG_FLAG bool = true //change it to false to avoid debug messages
var RANDOM_MIN int = 1     //it can use the characters between RANDOM_MIN and RANDOM_MAX
var RANDOM_MAX int = 9
var DIGIT_NUM int = 4 //change it to play the game with 5 characters

func generateRandomNumber() int { //basic random number generator between two given min and max parameters, used it while  generating non repeating random numbers
	return rand.Intn(RANDOM_MAX-RANDOM_MIN) + RANDOM_MIN
}

func generateRandomNumberNonRepeating() int { //generates a random number between RANDOM_MIN and RANDOM_MAX
	m := make(map[int]bool) //create map to store characters
	count := 0

	for count < DIGIT_NUM { //if we wanted the play the game with 5 digits all we have to do is change DIGIT_NUM to 5
		randomNumber := generateRandomNumber()
		if _, ok := m[randomNumber]; ok { //if random number has been previously created skip
			if DEBUG_FLAG {
				fmt.Printf("DEBUG: %v is in map\n", randomNumber)
			}
		} else { //else add it to the random number map
			if DEBUG_FLAG {
				fmt.Printf("DEBUG: %v is not in map\n", randomNumber)
			}
			m[randomNumber] = true //add element to map
			count++                //increase counter to keep track of digits
		}
	}
	if DEBUG_FLAG {
		fmt.Printf("DEBUG: Printing MAP:  ") //check if you've mapped your chars correctly
		bs, _ := json.Marshal(m)
		fmt.Println(string(bs))
	}

	map2Slice := make([]int, 0, len(m)) //create a slice of keys

	for k := range m {
		map2Slice = append(map2Slice, k) //append keys to slice
	}
	var returnStr string = ""

	for i := range map2Slice {
		returnStr += strconv.Itoa(map2Slice[i]) //append characters to return str
	}

	returnInt, _ := strconv.Atoi(returnStr) //convert str to int
	return returnInt                        //return int

}
