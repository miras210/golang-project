package game

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
	difficulty int
	gameMap    [][]rune
	characters []Character
}

func (g *Game) Init() {

}

type difficultyStrategy interface {
	setLevel(g *Game)
}

type easyLevel struct {
}

func (e *easyLevel) setLevel(g *Game) {
	g.gameMap = [][]rune{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '#', '#', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '#', '#', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '#', '#', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '#', '#', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
	}
	player := Character{
		skin:        'P',
		x:           1,
		y:           1,
		stamina:     5,
		health:      10,
		power:       2,
		attackRange: 3,
	}
	g.characters = append(g.characters, player)
	for i := 0; i < 1; i++ {
		g.characters = append(g.characters, Character{
			skin:        'E',
			x:           9, // TODO: random spawn method for enemies
			y:           9,
			stamina:     3,
			health:      5,
			power:       1,
			attackRange: 3,
		})
	}
}

type mediumLevel struct {
}

type hardLevel struct {
}
