package engine

import "fmt"

type country struct {
	name       string
	colour     string
	shares     map[int]*share
	money      int
	tanksTotal int
	tanksUsed  int
	shipsTotal int
	shipsUsed  int
	flagsTotal int
	flagsUsed  int
	action     RondelAction
	leader     *Player
	score      int
}

type share struct {
	country  *country
	interest int
	price    int
	owner    *Player
}

func (c *country) refreshLeader() {
	newLeader := c.leader

	//determine players who have more shares
	for _, s := range c.shares {
		if s.owner != nil && (c.leader == nil || (s.owner != c.leader && s.owner.totalShares(c) > c.leader.totalShares(c))) {
			//check if there are ties for new leaders
			if newLeader == nil || s.owner.totalShares(c) > newLeader.totalShares(c) {
				newLeader = s.owner
			} else if s.owner.totalShares(c) == newLeader.totalShares(c) && s.owner.investorDelta() < newLeader.investorDelta() {
				newLeader = s.owner
			}
		}
	}
	c.leader = newLeader
}

func (c *country) printCountry() string {
	result := ""

	result += fmt.Sprintf("%s\n", c.name)
	result += fmt.Sprintf("Tanks: Used-%d, Total-%d\n", c.tanksUsed, c.tanksTotal)
	result += fmt.Sprintf("Ships: Used-%d, Total-%d\n", c.shipsUsed, c.shipsTotal)
	result += fmt.Sprintf("Flags: Used-%d, Total-%d\n", c.flagsUsed, c.flagsTotal)
	result += fmt.Sprintf("Money: %d\n", c.money)
	result += fmt.Sprintf("Shares:\n")
	for _, s := range c.shares {
		owner := "No owner"
		if s.owner != nil {
			owner = fmt.Sprintf("Owner: %s", s.owner.name)
		}
		result += fmt.Sprintf("  Interest: %d, %s\n", s.interest, owner)
	}
	return result
}
