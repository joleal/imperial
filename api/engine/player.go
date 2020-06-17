package engine

import (
	"fmt"
)

//Player represents a player
type Player struct {
	playerID int
	name     string
	money    int
	shares   []*share
	order    int //order of play
	investor bool
	game     *Game
}

func (player *Player) buyShare(country *country, interest int) error {
	if interest < 1 || interest > 9 {
		return fmt.Errorf("Invalid share (%d)", interest)
	}

	share := country.shares[interest]
	if share.owner != nil {
		return fmt.Errorf("Share is not for sale, owner is %s", share.owner.name)
	} else if share.price > player.money {
		return fmt.Errorf("Player doesn't have enough money to buy share")
	} else {
		share.owner = player
		player.shares = append(player.shares, share)
		player.money -= share.price
		country.money += share.price
		country.refreshLeader()
	}

	return nil
}

func (player *Player) totalShares(country *country) int {

	result := 0
	for _, s := range player.shares {
		if s.country == country {
			result += s.price
		}
	}
	return result
}

func (player *Player) investorDelta() int {
	for _, p := range player.game.players {
		if p.investor {
			numbPlayers := len(player.game.players)
			return ((p.order - player.order) + numbPlayers) % numbPlayers
		}
	}
	return 0
}

func (player *Player) printShares() string {
	s := ""
	for _, share := range player.shares {
		s += fmt.Sprintf("Share: %s, Interest: %d\n", share.country.name, share.interest)
	}
	return s
}
