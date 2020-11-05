package game

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
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
	running    bool
	Player     Character
	difficulty difficultyStrategy
	gameMap    [][]rune
	characters []Character
	loot       []Loot
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
	lootLocation := make([][]int, g.difficulty.getNumberOfLoots())
	for i, loot := range g.loot {
		x, y := loot.getLocation()
		lootLocation[i] = make([]int, 2)
		lootLocation[i][0] = x
		lootLocation[i][1] = y
	}
	//TODO Miras, fix rendering here, please
	for a, row := range g.gameMap {
		for b, cell := range row {
			for i, enemy := range enemyLocation {
				if enemy[0] == a && enemy[1] == b {
					fmt.Print(string(g.characters[i].skin))
				}
			}
			for i, loot := range lootLocation {
				if loot[0] == a && loot[1] == b {
					fmt.Print(string(g.loot[i].skin))
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
	return g.running
}

var once sync.Once

func randomGen(min, max int) int {
	once.Do(func() {
		rand.Seed(time.Now().Unix())
	})
	return rand.Intn(max-min) + min
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
	g.gameMap = level.getLevel()
	g.difficulty = level
	g.Player = level.getPlayerStats()
	numberOfEnemies := level.getNumberOfEnemies()
	for i := 0; i < numberOfEnemies; i++ {
		var x, y int = 0, 0
		for g.gameMap[x][y] != '*' { //TODO also check if there is enemy in this cell
			x = randomGen(0, len(g.gameMap))
			y = randomGen(0, len(g.gameMap))
		}
		g.characters = append(g.characters, Character{
			skin:        'E',
			x:           x,
			y:           y,
			stamina:     3,
			health:      5,
			power:       1,
			attackRange: 3,
		})
	}
	numberOfLoots := level.getNumberOfLoots()
	for i := 0; i < 1; i++ {
		var x, y int = 0, 0
		for g.gameMap[x][y] != '*' { //TODO also check if there is enemy in this cell
			x = randomGen(0, len(g.gameMap))
			y = randomGen(0, len(g.gameMap))
		}
		g.loot = append(g.loot, Loot{
			skin: 'S',
			name: "sword",
			x:    x,
			y:    y,
		})
	}
	for i := 0; i < numberOfLoots-1; i++ {
		var x, y int = 0, 0
		for g.gameMap[x][y] != '*' { //TODO also check if there is enemy in this cell
			x = randomGen(0, len(g.gameMap))
			y = randomGen(0, len(g.gameMap))
		}
		g.loot = append(g.loot, Loot{
			skin: 'A',
			name: "armor",
			x:    x,
			y:    y,
		})
	}

	g.running = true
}
