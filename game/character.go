package game

import (
	"github.com/eiannone/keyboard"
	"math"
)

type CharacterI interface {
	getLocation() (int, int)
	getPreviousLocation() (int, int)
	getDamaged(character Character)
	getDistance(character Character) float64
	move() bool                      // move to a single cell costs 1 stamina
	attack(character Character) bool // attack by default costs 2 stamina
	isDead() bool
	loot()
	// TODO other moves like Buffs
}

type Character struct {
	skin        rune
	x           int // x coordinate
	y           int // y coordinate
	baseStamina int // TODO ? num of cells he can move in a turn
	curStamina  int
	health      float64 // num of health
	power       float64 // num of damage that he can produce
	attackRange float64 // attack range as a radius
}

func (c *Character) move(g *Game) bool {
	if c.curStamina == 0 {
		return false
	}

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	_, key, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}
	prevX, prevY := c.x, c.y
	if key == keyboard.KeyArrowLeft {
		c.y -= 1
	} else if key == keyboard.KeyArrowRight {
		c.y += 1
	} else if key == keyboard.KeyArrowUp {
		c.x -= 1
	} else if key == keyboard.KeyArrowDown {
		c.x += 1
	}
	if key == keyboard.KeyEsc {
		return false
	}
	correctLocation := true
	if g.gameMap[c.x][c.y] == '*' {
		for _, character := range g.characters {
			x, y := character.getLocation()
			if x == c.x && y == c.y {
				correctLocation = false
			}
		}
		if correctLocation {
			c.curStamina -= 1
			return true
		}
	}
	c.x, c.y = prevX, prevY
	return false
}

/*func (c *Character) Move() bool {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	_, key, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}
	for {
		if key == keyboard.KeyArrowLeft {
			c.y -= 1
			break
		} else if key == keyboard.KeyArrowRight {
			c.y += 1
			break
		} else if key == keyboard.KeyArrowUp {
			c.x -= 1
			break
		} else if key == keyboard.KeyArrowDown {
			c.x += 1
			break
		}
		if key == keyboard.KeyEsc {
			return false
		}
	}

	return true
}*/

func (c *Character) getDistance(character Character) float64 {
	return math.Sqrt(float64(c.x-character.x)*float64(c.x-character.x) + float64(c.y-character.y)*float64(c.y-character.y))
}

func (c *Character) getLocation() (int, int) {
	return c.x, c.y
}

func (c *Character) getDamaged(character Character) {
	c.health -= character.power
}

func (c *Character) isDead() bool {
	if c.health <= 0 {
		return true
	}
	return false
}

func (c *Character) attack(character Character) bool {
	if c.getDistance(character) <= c.attackRange {
		character.health -= c.power
		return true
	}
	return false
}

func (c *Character) loot(l Loot) {
	if l.name == "sword" {
		s := SwordDecorator{power: 3}
		c.power = s.power
	}
	if l.name == "armor" {
		a := ArmorDecorator{armor: 2}
		c.health += a.armor
	}
}
