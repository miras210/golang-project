package game

type Loot struct {
	skin rune
	name string
	x    int // x coordinate
	y    int // y coordinate
}
type SwordDecorator struct {
	power float64
	Loot
}

type ArmorDecorator struct {
	armor float64
	Loot
}

func (l *Loot) getLocation() (int, int) {
	return l.x, l.y
}
