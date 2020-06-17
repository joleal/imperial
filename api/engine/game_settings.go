package engine

//Settings define the settings for starting a game
type Settings struct {
	Version              Version
	NumberOfPlayers      int
	StartingMode         StartingMode
	InvestorCard         bool
	TaxIncreaseOnlyBonus bool
}

//config define the configurations of a game
type config struct {
	Version  string
	Setup    configSetup
	Rules    configRules
	BoardMap configMap
}

type configRules struct {
	TurnOrder     []string
	Tax           tax
	Rondel        []string
	RondelPenalty rondelPenalty
}

type rondelPenalty struct {
	base       int
	powerLevel bool
}

type configSetup struct {
	Countries       []configCountry
	Shares          []configShares
	Flags           []flag
	MoneyBasic      []int
	MoneyAdvanced   []int
	AllocationBasic [][][]string
	Factories       []string
}

type tax struct {
	TaxPerFactory int
	TaxPerFlag    int
	CostPerUnit   int
	Track         []trackSpace
	Bonus         bonusConfig
}

type trackSpace struct {
	TaxMin        int
	TaxMax        int
	Bonus         int
	PowerIncrease int
}

type bonusConfig struct {
	FromCountry  bool
	IncreaseOnly bool
}

type configCountry struct {
	Colour string
	Name   string
	Tanks  int
	Ships  int
	Flags  int
}

type configShares struct {
	Interest int
	Price    int
}

type flag struct {
	Colour      string
	Higher      string
	Lower       string
	HigherShare int
	LowerShare  int
}

type configMap struct {
	Areas []area
	Links []link
}

type area struct {
	Name     string
	AreaType string
	Factory  string
}

type link struct {
	From       string
	To         string
	Permission string
}
