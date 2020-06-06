package main

import "fmt"

type effect struct {
	name string
	ttl  int
}

type game struct {
	playerHitPoints int
	playerMama      int

	bossHitPoints int
	bossDamage    int

	effects []effect

	winner string
}

// Day22 is
func Day22() {

	init := make([]game, 0)
	start := append(init, initialGameStateTest1())
	results := tick(start)
	fmt.Println(results)
}

func initialGameStateTest1() game {
	effects := make([]effect, 0)
	return game{10, 250, 13, 8, effects, ""}
}

func missle(g game) game {
	// Magic Missile costs 53 mana. It instantly does 4 damage.

	applyEffects(g)
	g = checkWin(game{g.playerHitPoints, g.playerMama - 53, g.bossHitPoints - 4, g.bossDamage, g.effects, ""})

	if len(g.winner) > 0 {
		printIfWon(g)
		return g
	}

	applyEffects(g)
	g = checkWin(game{g.playerHitPoints - g.bossDamage, g.playerMama, g.bossHitPoints, g.bossDamage, g.effects, ""})
	printIfWon(g)
	return g

}

func poison(g game) game {

	// Poison costs 173 mana. It starts an effect that lasts for 6 turns. At the start of each turn while it is active, it deals the boss 3 damage.
	applyEffects(g)
	g.effects = append(g.effects, effect{"poison", 6})
	g = checkWin(game{g.playerHitPoints, g.playerMama - 173, g.bossHitPoints - 4, g.bossDamage, g.effects, ""})

	if len(g.winner) > 0 {
		printIfWon(g)
		return g
	}

	applyEffects(g)
	g = checkWin(game{g.playerHitPoints - g.bossDamage, g.playerMama, g.bossHitPoints, g.bossDamage, g.effects, ""})
	printIfWon(g)
	return g
}

func applyEffects(g game) game {
	n := g
	n.effects = make([]effect, 0)

	for _, spell := range g.effects {
		if spell.name == "poison" {

			n.bossHitPoints -= 3
			if spell.ttl > 0 {
				n.effects = append(n.effects, effect{"poison", spell.ttl - 1})
			}
		}

	}

	return n
}

func printIfWon(g game) {
	if g.winner != "" {
		fmt.Println(g.winner)
	}
}

func tick(games []game) []game {

	next := make([]game, 0)

	for _, g := range games {

		if len(g.winner) > 0 {
			continue
		}

		next = append(next, missle(g))

		if canCast(g, "poison") {
			next = append(next, poison(g))
		}

		// Drain costs 73 mana. It instantly does 2 damage and heals you for 2 hit points.
		// Shield costs 113 mana. It starts an effect that lasts for 6 turns. While it is active, your armor is increased by 7.
		// Recharge costs 229 mana. It starts an effect that lasts for 5 turns. At the start of each turn while it is active, it gives you 101 new mana.

	}

	if allWon(next) {
		return next
	}

	return tick(next)

}

func canCast(g game, effect string) bool {
	for _, s := range g.effects {
		if s.name == effect {
			return false
		}
	}
	return true
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
