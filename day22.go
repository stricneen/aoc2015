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

	winner    string
	manaSpent int

	//pastMoves []game
}

func (g game) Hello() {

}

// Day22 is
func Day22() {

	init := make([]game, 0)
	start := append(init, initialGameStateTest1())

	// Recharge
	// Shield
	// Drain
	// Poison
	// Magic Missile
	// g1 := recharge(initialGameStateTest1())
	// g2 := shield(g1)
	// g3 := drain(g2)
	// g4 := poison(g3)
	// g5 := missle(g4)

	// fmt.Println(g5.manaSpent)
	// return

	results := tick(start)

	minMana := 100000
	for _, v := range results {
		if v.winner == "wizard" && v.manaSpent < minMana {
			minMana = v.manaSpent
		}
	}

	//fmt.Println(results)
	fmt.Println("Lowest", minMana)
}

func initialGameStateTest1() game {
	effects := make([]effect, 0)
	//return game{50, 500, 51, 9, effects, "", 0}
	return game{10, 250, 14, 8, effects, "", 0}

	//return game{10, 250, 13, 8, effects, "", 0}
}

func missle(g game) game {
	// Magic Missile costs 53 mana. It instantly does 4 damage.

	g = applyEffects(g)
	if len(g.winner) > 0 {
		printIfWon(g)
		return g
	}

	g = checkWin(game{g.playerHitPoints, g.playerMama - 53, g.bossHitPoints - 4, g.bossDamage, g.effects, "", g.manaSpent + 53})
	if len(g.winner) > 0 {
		printIfWon(g)
		return g
	}

	g = applyEffects(g)
	if len(g.winner) > 0 {
		printIfWon(g)
		return g
	}

	g = checkWin(game{g.playerHitPoints - g.bossDamage + armor(g), g.playerMama, g.bossHitPoints, g.bossDamage, g.effects, "", g.manaSpent})
	printIfWon(g)
	return g

}

func drain(g game) game {
	// Drain costs 73 mana. It instantly does 2 damage and heals you for 2 hit points.

	g = applyEffects(g)
	if len(g.winner) > 0 {
		printIfWon(g)
		return g
	}

	g = checkWin(game{g.playerHitPoints + 2, g.playerMama - 73, g.bossHitPoints - 2, g.bossDamage, g.effects, "", g.manaSpent + 73})
	if len(g.winner) > 0 {
		printIfWon(g)
		return g
	}

	g = applyEffects(g)
	if len(g.winner) > 0 {
		printIfWon(g)
		return g
	}

	g = checkWin(game{g.playerHitPoints - g.bossDamage + armor(g), g.playerMama, g.bossHitPoints, g.bossDamage, g.effects, "", g.manaSpent})
	printIfWon(g)
	return g

}

func shield(g game) game {
	// Shield costs 113 mana. It starts an effect that lasts for 6 turns. While it is active, your armor is increased by 7.

	g = applyEffects(g)
	if won(g) {
		return g
	}

	g.effects = append(g.effects, effect{"shield", 6})
	g = checkWin(game{g.playerHitPoints, g.playerMama - 113, g.bossHitPoints, g.bossDamage, g.effects, "", g.manaSpent + 113})
	if won(g) {
		return g
	}

	g = applyEffects(g)
	if won(g) {
		return g
	}

	g = checkWin(game{g.playerHitPoints - g.bossDamage + armor(g), g.playerMama, g.bossHitPoints, g.bossDamage, g.effects, "", g.manaSpent})
	printIfWon(g)
	return g
}

func recharge(g game) game {
	// Recharge costs 229 mana. It starts an effect that lasts for 5 turns. At the start of each turn while it is active, it gives you 101 new mana.

	g = applyEffects(g)
	if len(g.winner) > 0 {
		printIfWon(g)
		return g
	}

	g.effects = append(g.effects, effect{"recharge", 5})
	g = checkWin(game{g.playerHitPoints, g.playerMama - 229, g.bossHitPoints, g.bossDamage, g.effects, "", g.manaSpent + 229})
	if len(g.winner) > 0 {
		printIfWon(g)
		return g
	}

	g = applyEffects(g)
	if len(g.winner) > 0 {
		printIfWon(g)
		return g
	}

	g = checkWin(game{g.playerHitPoints - g.bossDamage + armor(g), g.playerMama, g.bossHitPoints, g.bossDamage, g.effects, "", g.manaSpent})
	printIfWon(g)
	return g
}

func poison(g game) game {

	// Poison costs 173 mana. It starts an effect that lasts for 6 turns. At the start of each turn while it is active, it deals the boss 3 damage.
	g = applyEffects(g)
	if len(g.winner) > 0 {
		printIfWon(g)
		return g
	}

	g.effects = append(g.effects, effect{"poison", 6})
	g = checkWin(game{g.playerHitPoints, g.playerMama - 173, g.bossHitPoints, g.bossDamage, g.effects, "", g.manaSpent + 173})
	if len(g.winner) > 0 {
		printIfWon(g)
		return g
	}

	g = applyEffects(g)
	if len(g.winner) > 0 {
		printIfWon(g)
		return g
	}

	g = checkWin(game{g.playerHitPoints - g.bossDamage + armor(g), g.playerMama, g.bossHitPoints, g.bossDamage, g.effects, "", g.manaSpent})
	printIfWon(g)
	return g
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

// tl 671

func tick(games []game) []game {

	next := make([]game, 0)

	for _, g := range games {

		if won(g) {
			next = append(next, g)
			continue
		}

		next = append(next, missle(g))

		if canCast(g, "recharge", 229) {
			next = append(next, recharge(g))
		}

		if canCast(g, "shield", 113) {
			next = append(next, shield(g))
		}

		next = append(next, drain(g))

		if canCast(g, "poison", 173) {
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

var min int = 100000

func won(g game) bool {
	if len(g.winner) > 0 {
		printIfWon(g)
		return true
	}
	return false
}

func printIfWon(g game) {
	if g.manaSpent < min && g.winner == "wizard" {
		min = g.manaSpent
		fmt.Println(min)
	}

	// if g.winner != "" {
	// 	fmt.Println(g.winner, g.manaSpent)
	// }
}

func canCast(g game, effect string, cost int) bool {
	for _, s := range g.effects {
		if s.name == effect && g.playerMama < cost {
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
