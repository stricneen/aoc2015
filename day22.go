package main

import "fmt"

type effect struct {
	name string
	ttl  int
}

type effectFn func(game) game

type game struct {
	playerHitPoints int
	playerMama      int

	bossHitPoints int
	bossDamage    int

	effects []effect

	winner    string
	manaSpent int
}

// Day22 is
func Day22() {

	init := make([]game, 0)
	start := append(init, initialGameStateTest1())

	results := tick(start)

	minMana := 100000
	for _, v := range results {
		if v.winner == "wizard" {
			fmt.Println(v)
		}
		if v.winner == "wizard" && v.manaSpent < minMana {
			minMana = v.manaSpent
		}
	}
	fmt.Println("Lowest", minMana)
}

func initialGameStateTest1() game {
	effects := make([]effect, 0)
	return game{50, 500, 51, 9, effects, "", 0}
	//return game{10, 250, 14, 8, effects, "", 0}
	//return game{10, 250, 13, 8, effects, "", 0}
}

func turn(g game, spell effectFn) game {

	g = hard(g)
	if won(g) {
		return g
	}

	g = applyEffects(g)
	if won(g) {
		return g
	}

	g = spell(g)
	g = checkWin(g)
	if won(g) {
		return g
	}

	g = applyEffects(g)
	if won(g) {
		return g
	}

	g = checkWin(game{g.playerHitPoints - g.bossDamage + armor(g), g.playerMama, g.bossHitPoints, g.bossDamage, g.effects, "", g.manaSpent})
	return g

}

func hard(g game) game {
	//return g
	r := g
	//r.playerHitPoints--
	return checkWin(r)
}

func armor(g game) int {
	for _, v := range g.effects {
		if v.name == "shield" {
			return 7
		}
	}
	return 0
}

func applyEffects(g game) game {
	n := g
	n.effects = make([]effect, 0)

	for _, spell := range g.effects {
		if spell.name == "poison" {
			n.bossHitPoints -= 3
			if spell.ttl > 1 {
				n.effects = append(n.effects, effect{"poison", spell.ttl - 1})
			}
		}

		if spell.name == "recharge" {
			n.playerMama += 101
			if spell.ttl > 1 {
				n.effects = append(n.effects, effect{"recharge", spell.ttl - 1})
			}
		}

		if spell.name == "shield" {
			if spell.ttl > 1 {
				n.effects = append(n.effects, effect{"shield", spell.ttl - 1})
			}
		}
	}

	return checkWin(n)
}

// th 1242

func tick(games []game) []game {

	next := make([]game, 0)

	for _, g := range games {

		if won(g) {
			next = append(next, g)
			continue
		}

		if canCast(g, "missle", 53) {
			next = append(next, turn(g, func(g game) game {
				return game{g.playerHitPoints, g.playerMama - 53, g.bossHitPoints - 4, g.bossDamage, g.effects, "", g.manaSpent + 53}
			}))
		}

		if canCast(g, "recharge", 229) {
			next = append(next, turn(g, func(g game) game {
				g.effects = append(g.effects, effect{"recharge", 5})
				return game{g.playerHitPoints, g.playerMama - 229, g.bossHitPoints, g.bossDamage, g.effects, "", g.manaSpent + 229}
			}))
		}

		if canCast(g, "shield", 113) {
			next = append(next, turn(g, func(g game) game {
				g.effects = append(g.effects, effect{"shield", 6})
				return game{g.playerHitPoints, g.playerMama - 113, g.bossHitPoints, g.bossDamage, g.effects, "", g.manaSpent + 113}
			}))
		}

		if canCast(g, "drain", 73) {
			next = append(next, turn(g, func(g game) game {
				return game{g.playerHitPoints + 2, g.playerMama - 73, g.bossHitPoints - 2, g.bossDamage, g.effects, "", g.manaSpent + 73}
			}))
		}

		if canCast(g, "poison", 173) {
			next = append(next, turn(g, func(g game) game {
				g.effects = append(g.effects, effect{"poison", 6})
				return game{g.playerHitPoints, g.playerMama - 173, g.bossHitPoints, g.bossDamage, g.effects, "", g.manaSpent + 173}
			}))
		}
	}

	if allWon(next) {
		return next
	}

	return tick(next)

}

func won(g game) bool {
	if len(g.winner) > 0 {
		return true
	}
	return false
}

func canCast(g game, effect string, cost int) bool {
	for _, s := range g.effects {
		if s.name == effect {
			return false
		}
	}
	return g.playerMama >= cost
}

func allWon(games []game) bool {
	for _, x := range games {
		if len(x.winner) == 0 {
			return false
		}
	}
	return true
}

func checkWin(g game) game {
	r := g
	if g.playerHitPoints < 1 {
		r.winner = "boss"
	}
	if g.bossHitPoints < 1 {
		r.winner = "wizard"
	}
	return r
}
