package game

type Difficulty interface {
	GetNumberOfEnemies() int
}

type EasyDifficulty struct {
}

func (e *EasyDifficulty) GetNumberOfEnemies() int {
	return 2
}

type MediumDifficulty struct {
}

func (m *MediumDifficulty) GetNumberOfEnemies() int {
	return 4
}

type HardDifficulty struct {
}

func (h *HardDifficulty) GetNumberOfEnemies() int {
	return 6
}
