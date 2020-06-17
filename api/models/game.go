package models

//SaveGame represents a saved game
type SaveGame struct {
	ID             int
	GameSettings   gameSettings
	Players        []*gamePlayer
	Countries      []*gameCountries
	Shares         []*gameShares
	GameStart      bool
	GameFinish     bool
	CurrentCountry string
	CurrentPlayer  string
	CurrentAction  string
	Actions        []*gameAction
}

type gameSettings struct {
	Version              string
	NumberOfPlayers      int
	InvestorCard         bool
	TaxIncreaseOnlyBonus bool
}

type gamePlayer struct {
	Name     string
	PlayerID int
	Money    int
	Order    int
	Investor bool
}

type gameCountries struct {
	Colour       string
	Money        int
	TanksTotal   int
	TanksUsed    int
	ShipsTotal   int
	ShipsUsed    int
	FlagsTotal   int
	FlagsUsed    int
	Leader       string
	RondelAction string
	Score        int
}

type gameShares struct {
	Colour   string
	Owner    string
	Interest string
}

type gameAction struct {
}
