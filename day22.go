package main

type spell struct {
	cost   int
	damage int
	heal   int
}

type game struct {
	playerHitPoints int
	playerMama      int

	bossHitPoints int
	bossDamage    int

	effects []string
}

// Day22 is
func Day22() {

	// bossHitPoint := 51
	// bossDamage := 9

	init := make([]game, 0)
	start := append(init, initialGameStateTest1())
	tick(start)
}

func initialGameStateTest1() game {
	effects := make([]string, 0)
	return game{10, 250, 13, 8, effects}
}

func tick(games []game) []game {

	next := make([]game, 0)
	effect := make([]string, 0)

	for _, g := range games {

		// Apply each possible
		// Magic Missile costs 53 mana. It instantly does 4 damage.
		n := game{g.playerHitPoints, g.playerMama - 53, g.bossHitPoints - 4, g.bossDamage, effect}
		next = append(next, n)
		//  playerHitPoints  int
		// 	playerMama       int
		// 	bossHitPoints     int
		// 	bossDamage       int
		// 	effects          []string

		// Drain costs 73 mana. It instantly does 2 damage and heals you for 2 hit points.
		// Shield costs 113 mana. It starts an effect that lasts for 6 turns. While it is active, your armor is increased by 7.
		// Poison costs 173 mana. It starts an effect that lasts for 6 turns. At the start of each turn while it is active, it deals the boss 3 damage.
		// Recharge costs 229 mana. It starts an effect that lasts for 5 turns. At the start of each turn while it is active, it gives you 101 new mana.

	}

	return next

}

// playerHitPoints int
// 	playerMama      int
// 	bosHitPoints int
// 	bossDamage   int
// 	effects []string
