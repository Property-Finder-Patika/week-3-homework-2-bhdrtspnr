// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. I don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
// ---------------------------------------------------------

func main() {
	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item
		genre string
	}

	games := []game{
		{
			item:  item{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		},
		{item: item{id: 2, name: "x-com 2", price: 40}, genre: "strategy"},
		{item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
	}

	// index the games by id
	byID := make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}

	for {
		fmt.Println("1 to list")
		fmt.Println("2 to query by id")
		fmt.Println("3 to quit")
		var menu int
		fmt.Scanln(&menu)
		switch menu {
		case 1:
			fmt.Printf("%v", games)
		case 2:
			fmt.Println("Enter id:")
			var qId int
			foundFlag := false
			fmt.Scanln(&qId)
			for _, game := range games {
				if game.id == qId {
					fmt.Printf("Game found: %v", game)
					foundFlag = true
				}
			}
			if !foundFlag {
				fmt.Printf("No game found by id %d", qId)
			}
		case 3:
			fmt.Println("Bye")
			return
		default:
			fmt.Println("Please enter a valid input")
		}
	}

}
