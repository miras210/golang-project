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
	difficulty Difficulty
	gameMap    [][]rune
	characters []Character
}

func (g *Game) Init() {

}
