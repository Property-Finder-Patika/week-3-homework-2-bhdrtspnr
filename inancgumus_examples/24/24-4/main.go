// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"encoding/json"
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
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
	type jsonGame struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Genre string `json:"genre"`
		Price int    `json:"price"`
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
		fmt.Println("3 to save as json")
		fmt.Println("0 to quit")
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
			var encodable []jsonGame
			for _, g := range games {
				encodable = append(encodable,
					jsonGame{g.id, g.name, g.genre, g.price})
			}

			out, err := json.MarshalIndent(encodable, "", "\t")
			if err != nil {
				fmt.Println("Sorry:", err)
				continue
			}

			fmt.Println(string(out))
			return

		case 0:
			fmt.Println("Bye")
			return
		default:
			fmt.Println("Please enter a valid input")
		}
	}

}
