package game

type loot struct {
	name string
}
type SwordDecorator struct {
	power float64
	loot
}

type ArmorDecorator struct {
	armor float64
	loot
}

func (c *Character) loot(l loot) {
	if l.name == "sword" {
		s := SwordDecorator{power: 3}
		c.power = s.power
	}
	if l.name == "armor" {
		a := ArmorDecorator{armor: 2}
		c.health += a.armor
	}
}
