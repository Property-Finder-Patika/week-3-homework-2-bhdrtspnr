package main

import "fmt"

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item  item
	genre string
}

func newGame(id int, price int, name string, genre string) game {
	return game{item: item{id: id, name: name, price: price}, genre: genre}
}

func addGame(game game, games []game) []game {
	games = append(games, game)
	return games
}

func load() []game {
	//predefined load
	var games []game
	return addGame(newGame(1, 15, "Bahadir", "strategy"), games)
}

func indexByID(games []game) map[int]game {
	byID := make(map[int]game)

	for _, g := range games {
		byID[g.item.id] = g
	}
	return byID
}

func printGame(game game) {
	fmt.Println(game)
}

func listGames(games []game) {
	for _, g := range games {
		fmt.Printf("#%d: %-15q %-20s $%d\n",
			g.item.id, g.item.name, "("+g.genre+")", g.item.price)
	}
}
