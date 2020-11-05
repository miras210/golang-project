package main

import (
	game3 "final-project/game"
	"fmt"
)

/*
Available commands:
- move [left|right|up|down]
- attack x y (attack the cell)
- inventory heal if potions were found
- loot current cell (random chance of getting buffs and healing potions)
*/
var game game3.Game

func main() {
	fmt.Println("Welcome to Dungeon Master RPG!")
	fmt.Println("Choose your difficulty! [easy, medium, hard]")
	var difficulty string
	fmt.Scanln(&difficulty)
	game.Init(difficulty)
	game.Display("")
	for game.IsRunning() {
		game.Display(game.TurnStart())
	}
}
