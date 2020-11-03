package game

import (
	"fmt"
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
