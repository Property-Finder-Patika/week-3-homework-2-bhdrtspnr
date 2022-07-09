// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Refactor the game store to funcs - step #1
//
//  Remember the game store program from the structs exercises?
//  Now, it's time to refactor it to funcs.
//
//  Create games.go file
//
//     1- Move the structs there
//
//     2- Add a func that creates and returns a game.
//
//        Name  : newGame
//        Inputs: id, price, name and genre
//        Output: game
//
//     3- Add a func that adds a `game` to `[]game` and returns `[]game`:
//
//        Name  : addGame
//        Inputs: []game, game
//        Output: []game
//
//     4- Add a func that uses newGame and addGame funcs to load up the game
//        store:
//
//        Name  : load
//        Inputs: none
//        Output: []game
//
//     5- Add a func that indexes games by ID:
//
//        Name  : indexByID
//        Inputs: []game
//        Output: map[int]game
//
//     6- Add a func that prints the given game:
//
//        Name  : printGame
//        Inputs: game
//
//
//  Go back to main.go and change the existing code with
//  the new funcs that you've created. There are hints
//  inside the code.
//
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	games := []game{
		{
			item:  item{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		},
		{item: item{id: 2, name: "x-com 2", price: 40}, genre: "strategy"},
		{item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
	}

	fmt.Println("1 to list")
	fmt.Println("2 to index games by id")
	fmt.Println("3 to add game")
	fmt.Println("0 to quit")

	var menu int
	fmt.Scanln(&menu)
	//parameters are missing from the functions
	switch menu {
	case 0:
		fmt.Println("Bye")
		return
	case 1:
		listGames()
	case 2:
		indexByID()
	case 3:
		addGame()
	}

}
