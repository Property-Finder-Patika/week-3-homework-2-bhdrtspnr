// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Students
//
//  Create a program that returns the students by the given
//  Hogwarts house name (see the data below).
//
//  Print the students sorted by name.
//
//  "bobo" doesn't belong to Hogwarts, remove it by using
//  the delete function.
//
//
// RESTRICTIONS
//
//  + Add the following data to your map as is.
//    Do not sort it manually and do not modify it.
//
//  + Slices in the map shouldn't be sorted (changed).
//    HINT: Copy them.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//  Please type a Hogwarts house name.
//
//
//  go run main.go bobo
//
//  Sorry. I don't know anything about "bobo".
//
//
//  go run main.go hufflepuf
//
//  ~~~ hufflepuf students ~~~
//
//        + diggory
//        + helga
//        + scamander
//        + wenlock
//
// ---------------------------------------------------------

func main() {
	// House        Student Name
	// ---------------------------

	houseStudentMap := make(map[string][]string)
	//go excel copy all add to notepad++
	//replace [ with  [", replace ] with "], replace , with ," replace ) with ")
	houseStudentMap["gryffindor"] = append(houseStudentMap["gryffindor"], "weasley")
	houseStudentMap["gryffindor"] = append(houseStudentMap["gryffindor"], "hagrid")
	houseStudentMap["gryffindor"] = append(houseStudentMap["gryffindor"], "dumbledore")
	houseStudentMap["gryffindor"] = append(houseStudentMap["gryffindor"], "lupin")
	houseStudentMap["hufflepuf"] = append(houseStudentMap["hufflepuf"], "wenlock")
	houseStudentMap["hufflepuf"] = append(houseStudentMap["hufflepuf"], "scamander")
	houseStudentMap["hufflepuf"] = append(houseStudentMap["hufflepuf"], "helga")
	houseStudentMap["hufflepuf"] = append(houseStudentMap["hufflepuf"], "diggory")
	houseStudentMap["ravenclaw"] = append(houseStudentMap["ravenclaw"], "flitwick")
	houseStudentMap["ravenclaw"] = append(houseStudentMap["ravenclaw"], "bagnold")
	houseStudentMap["ravenclaw"] = append(houseStudentMap["ravenclaw"], "wildsmith")
	houseStudentMap["ravenclaw"] = append(houseStudentMap["ravenclaw"], "montmorency")
	houseStudentMap["slytherin"] = append(houseStudentMap["slytherin"], "horace")
	houseStudentMap["slytherin"] = append(houseStudentMap["slytherin"], "nigellus")
	houseStudentMap["slytherin"] = append(houseStudentMap["slytherin"], "higgs")
	houseStudentMap["slytherin"] = append(houseStudentMap["slytherin"], "scorpius")
	houseStudentMap["bobo"] = append(houseStudentMap["bobo"], "wizardry")
	houseStudentMap["bobo"] = append(houseStudentMap["bobo"], "unwanted")
	delete(houseStudentMap, "bobo")

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no arg given")
		return
	}

	house, students := args[0], houseStudentMap[args[0]]
	if students == nil {
		fmt.Printf("no house found:  %q.\n", house)
		return
	}

	copied := append([]string(nil), students...)
	sort.Strings(copied)

	fmt.Printf(" %s students \n", house)
	for _, student := range copied {
		fmt.Printf("\t+ %s\n", student)
	}

}
