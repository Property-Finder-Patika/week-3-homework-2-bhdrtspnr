// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
// ---------------------------------------------------------

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item  item
	genre string
}

var games = []game{
	{
		item:  item{id: 1, name: "god of war", price: 50},
		genre: "action adventure",
	},
	{item: item{id: 2, name: "x-com 2", price: 40}, genre: "strategy"},
	{item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
}

func main() {
	r := bufio.NewReader(os.Stdin)
	menu := true
	var menuInput int
	for menu {
		fmt.Println("Enter 1 to search for a game")
		fmt.Println("Enter 2 to list all the games")
		fmt.Println("Enter 3 to quit.")
		fmt.Scanln(&menuInput)
		switch menuInput {
		case 1:
			fmt.Println("Please enter name to search")
			searchName, err := r.ReadString('\n')
			_ = err
			returnval := searchGame(searchName)
			if returnval == nil {
				fmt.Printf("No game found with name: %s\n", searchName)
			} else {
				fmt.Printf("Found a game here are the details \n %v \n", returnval)
			}
		case 2:
			fmt.Printf("List of games: \n %v\n", games)
		case 3:
			return
		default:
			fmt.Println("Please enter a valid input")
		}

	}

}

func searchGame(s string) *game {
	for _, game := range games {
		if game.item.name == s {
			return &game
		}
	}
	return nil
}
