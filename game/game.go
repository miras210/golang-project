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
	player     Character
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

	x, y := g.player.getLocation()
	g.gameMap[x][y] = g.player.skin
	for _, enemy := range g.characters {
		x, y := enemy.getLocation()
		g.gameMap[x][y] = enemy.skin
	}

	for _, row := range g.gameMap {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}

	// JUST FOR FUN TEST
	g.gameMap[g.player.x][g.player.y] = '*'
	g.player.x += 1

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
	g.player = level.getPlayerStats()
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
