package game

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

/*

	##########
	#********#
	#********#
	#********#
	#********#
	#********#
	#********#
	#********#
	##########

	# - wall
	* - empty space
	P - player
	E - enemy
*/

type Game struct {
	Player     Character
	difficulty difficultyStrategy
	gameMap    [][]rune
	characters []Character
}

func (g *Game) Display() {
	//TODO clear the previous screen

	// WORKS ONLY IN TERMINALS BASH / CMD / .exe file
	// Clearing the console
	clear := make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	value, ok := clear[runtime.GOOS]

	if ok {
		value()
	}

	x, y := g.Player.getLocation()
	enemyLocation := make([][]int, g.difficulty.getNumberOfEnemies())
	for i, enemy := range g.characters {
		x, y := enemy.getLocation()
		enemyLocation[i] = make([]int, 2)
		enemyLocation[i][0] = x
		enemyLocation[i][1] = y
	}

	for a, row := range g.gameMap {
		for b, cell := range row {
			for i, enemy := range enemyLocation {
				if enemy[0] == a && enemy[1] == b {
					fmt.Print(string(g.characters[i].skin))
				}
			}
			if a == x && b == y {
				fmt.Print(string(g.Player.skin))
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}

}

func (g *Game) IsRunning() bool {
	return true
	//TODO check if game is running
}

func (g *Game) Init(difficulty string) {
	var level difficultyStrategy
	switch difficulty {
	case "easy":
		level = &easyLevel{}
	case "medium":
		level = &mediumLevel{}
	case "hard":
		level = &hardLevel{}
	default:
		level = &easyLevel{}
	}
	level.setLevel(g)
	g.difficulty = level
	g.Player = level.getPlayerStats()
	numberOfEnemies := level.getNumberOfEnemies()
	for i := 0; i < numberOfEnemies; i++ {
		g.characters = append(g.characters, Character{
			skin:        'E',
			x:           8, // TODO: random spawn method for enemies
			y:           8,
			stamina:     3,
			health:      5,
			power:       1,
			attackRange: 3,
		})
	}
}
