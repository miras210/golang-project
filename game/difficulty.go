package game

type difficultyStrategy interface {
	getLevel() [][]rune
	getNumberOfEnemies() int
	getPlayerStats() Character
	getNumberOfLoots() int
}

type easyLevel struct {
}

func (e *easyLevel) getLevel() [][]rune {
	return [][]rune{
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
}

func (e *easyLevel) getNumberOfEnemies() int {
	return 3
}

func (e *easyLevel) getNumberOfLoots() int {
	return 3
}

func (e *easyLevel) getPlayerStats() Character {
	return Character{
		skin:        'P',
		x:           1,
		y:           1,
		baseStamina: 5,
		curStamina:  5,
		health:      10,
		power:       2,
		attackRange: 3,
	}
}

type mediumLevel struct {
	//TODO implement mediumLevel thoroughly, i just copied easylevel
}

func (ml *mediumLevel) getNumberOfLoots() int {
	panic("implement me")
}

func (ml *mediumLevel) getLevel() [][]rune {
	return [][]rune{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '#', '#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '#', '#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '#', '#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '#', '#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
	}
}

func (ml *mediumLevel) getNumberOfEnemies() int {
	return 5
}

func (ml *mediumLevel) getPlayerStats() Character {
	return Character{
		skin:        'P',
		x:           1,
		y:           1,
		baseStamina: 5,
		curStamina:  5,
		health:      10,
		power:       2,
		attackRange: 3,
	}
}

type hardLevel struct {
	//TODO implement hardLevel thoroughly, i just copied easylevel
}

func (hl *hardLevel) getNumberOfLoots() int {
	panic("implement me")
}
func (hl *hardLevel) getLevel() [][]rune {

	return [][]rune{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '#', '#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '#', '#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '#', '#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '#', '#', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
	}
}

func (hl *hardLevel) getNumberOfEnemies() int {
	return 8
}

func (hl *hardLevel) getPlayerStats() Character {
	return Character{
		skin:        'P',
		x:           1,
		y:           1,
		baseStamina: 5,
		curStamina:  5,
		health:      10,
		power:       2,
		attackRange: 3,
	}
}
