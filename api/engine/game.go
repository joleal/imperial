package engine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

//Version represents the version of the game being played
type Version string

//comment
const (
	Original Version = "Original"
	I2030    Version = "Imperial 2030"
)

//StartingMode represents how the game starts
type StartingMode string

//comment
const (
	Random      StartingMode = "Random"
	Experienced StartingMode = "Experienced"
)

//RondelAction represents an action in the rondel
type RondelAction int

//actions of the rondel
const (
	Investor    RondelAction = 0
	Import      RondelAction = 1
	Production1 RondelAction = 2
	Maneuver1   RondelAction = 3
	Factory     RondelAction = 4
	Production2 RondelAction = 5
	Maneauver2  RondelAction = 6
)

// Unit represents a unit (tank or ship)
type Unit struct {
	country  string
	unitType string
}

// Game represents the game
type Game struct {
	id            int
	version       Version
	startingMode  StartingMode
	players       []*Player
	gameMap       *BoardMap
	countries     map[string]*country
	config        *config
	currentTurn   *country
	currentPlayer *Player
	gameStart     bool
	gameFinish    bool
}

// User represents a user in the app
type User struct {
	Name string
}

//StartGame creates a new game
func (g *Game) StartGame(users []User) error {

	//Add Players
	g.players = make([]*Player, len(users))
	for i := 0; i < len(users); i++ {
		g.players[i] = &Player{
			name:  users[i].Name,
			order: i,
			game:  g,
		}
	}
	//Add starting Money
	var startMoney int
	if g.startingMode == Random {
		startMoney = g.config.Setup.MoneyBasic[len(g.players)-2] //the allocation is an array, the first position is for 2 players
	} else if g.startingMode == Experienced {
		startMoney = g.config.Setup.MoneyAdvanced[len(g.players)-2]
	} else {
		return fmt.Errorf("Unknown starting mode (%s)", g.startingMode)
	}
	for _, p := range g.players {
		p.money = startMoney
	}

	if g.startingMode == Random {
		allocation := g.config.Setup.AllocationBasic[len(g.players)-2]
		rand.Seed(time.Now().UnixNano())
		//shuffle allocation
		for i := 0; i < len(allocation); i++ {
			r := rand.Intn(i + 1)
			if i != r {
				allocation[i], allocation[r] = allocation[r], allocation[i]
			}
		}
		for i, p := range g.players {
			for _, a := range allocation[i] {
				for _, f := range g.config.Setup.Flags {
					if f.Colour == a {
						p.buyShare(g.countries[f.Higher], f.HigherShare)
						p.buyShare(g.countries[f.Lower], f.LowerShare)
					}
				}
			}
		}
		g.determineStart()
	} else {
		g.currentPlayer = g.players[0]
		g.currentTurn = g.countries[g.config.Rules.TurnOrder[0]]
	}

	return nil
}

//CreateGame creates the game
func (g *Game) CreateGame(settings Settings) error {
	g.version = settings.Version
	g.startingMode = settings.StartingMode

	var fileName string
	switch g.version {
	case I2030:
		fileName = "config/game_config/imperial2030.json"
	case Original:
		fileName = "config/game_config/imperialOriginal.json"
	}
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return err
	}

	var config config
	err = json.Unmarshal([]byte(file), &config)
	if err != nil {
		return fmt.Errorf("Failed to parse config file (%s)", fileName)
	}

	g.config = &config

	//Create Map - Areas
	m := &BoardMap{}
	for _, r := range config.BoardMap.Areas {
		m.AddNode(&MapArea{
			name:     r.Name,
			areaType: r.AreaType,
			factory:  r.Factory != "",
		})
	}

	//Create Map - Factories
	for _, r := range config.Setup.Factories {
		if val, ok := m.areas[r]; ok {
			if val.areaType != "land" && val.areaType != "sea" {
				val.factoryBuild = true
			} else {
				return fmt.Errorf("Factory in area (%s) that doesn't support factories", r)
			}
		} else {
			return fmt.Errorf("Factory in area (%s) that doesn't exist", r)
		}
	}

	//Create Map - Add Links
	for _, r := range config.BoardMap.Links {
		if _, ok := m.areas[r.From]; !ok {
			return fmt.Errorf("Link from area (%s) that doesn't exist", r)
		} else if _, ok := m.areas[r.To]; !ok {
			return fmt.Errorf("Link to area (%s) that doesn't exist", r)
		} else {
			m.AddLink(m.areas[r.From], m.areas[r.To], r.Permission)
		}
	}
	g.gameMap = m

	//Add countries
	g.countries = make(map[string]*country)
	for _, c := range config.Setup.Countries {
		country := &country{
			name:       c.Name,
			colour:     c.Colour,
			money:      0,
			tanksTotal: c.Tanks,
			tanksUsed:  0,
			shipsTotal: c.Ships,
			shipsUsed:  0,
			flagsTotal: c.Flags,
			flagsUsed:  0,
			shares:     make(map[int]*share),
			action:     -1, //to represent it's not on the rondel
			score:      0,
		}

		for _, s := range config.Setup.Shares {
			country.shares[s.Interest] = &share{
				interest: s.Interest,
				price:    s.Price,
				country:  country,
			}
		}

		g.countries[c.Colour] = country
	}

	return nil
}

func (g *Game) determineStart() {
	//The first country to act is the first country which has a government, based on the turn order
	for _, i := range g.config.Rules.TurnOrder {
		c := g.countries[i]
		if c.leader != nil {
			g.currentTurn = c
			g.currentPlayer = c.leader
			g.gameStart = true
			return
		}
	}
}

//LoadGame loads a game from save file
func (g *Game) LoadGame(settings Settings) {

}

//Country returns the status of the countries
func (g *Game) Country() string {
	status := "\n---Countries---\n"
	for _, c := range g.countries {
		status += c.printCountry()
	}
	return status
}

//Map returns the map of the game
func (g *Game) Map() string {
	return fmt.Sprintln(g.gameMap)
}

//Status returns the status of the game
func (g *Game) Status() string {
	status := "Status of the game:\n"
	status += fmt.Sprintf("GameId = %d", g.id)
	status += "\n---Players---\n"
	for i, s := range g.players {
		status += fmt.Sprintf("=> Player %d - %s\n", i+1, s.name)
		status += fmt.Sprintf("Money: %d\n", s.money)
		status += fmt.Sprintf(s.printShares())
	}
	if g.gameStart {
		status += "***Game has started***"
	} else {
		status += "***Game hasn't started***\n"
	}
	status += fmt.Sprintf("Current Turn: %s\n", g.currentTurn.name)
	status += fmt.Sprintf("Current Player: %s\n", g.currentPlayer.name)
	return status
}

//NextPlayer returns the next player to play
func (g *Game) NextPlayer() string {
	return "player"
}
