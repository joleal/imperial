package models

import (
	"encoding/json"
	"log"
	"strings"
)

//SaveGame represents a saved game
type SaveGame struct {
	ID             string
	Name           string        `json:"name"`
	GameSettings   gameSettings  `json:"settings"`
	Players        []*gamePlayer `json:"players"`
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
	NumberOfPlayers      int    `json:"numberOfPlayers"`
	InvestorCard         bool   `json:"investorCard"`
	TaxIncreaseOnlyBonus bool   `json:"taxIncreaseBonus"`
	Random               bool   `json:"random"`
	StartingMode         string `json:"startingMode"`
}

type gamePlayer struct {
	Nickname string `json:"nickname"`
	PlayerID string `json:"user_id"`
	Email    string `json:"email"`
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

//GetActiveGames gets active games for user
func GetActiveGames(user string) (*[]*SaveGame, error) {
	db := CreateDb()
	defer db.closeConnection()

	sql := `

		SELECT 
			game_id AS id,
			name
		FROM game g
		INNER JOIN game_player gp 
			ON gp.fk_game = g.game_id
		WHERE 
			g.game_finish = 0
			AND g.game_start = 1
			AND gp.fk_user = ?;`

	rows, err := db.Query(sql, user)
	if err != nil {
		log.Println(err)
	}

	games := make([]*SaveGame, 0)
	for rows.Next() {
		game := &SaveGame{}
		err := rows.Scan(&game.ID, &game.Name)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		games = append(games, game)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return &games, nil
}

//GetOpenGames gets games without all players
func GetOpenGames() (*[]*SaveGame, error) {
	db := CreateDb()
	defer db.closeConnection()

	//games
	sql := `

		SELECT 
			game_id AS id,
			name,
			CONCAT('[', GROUP_CONCAT(JSON_OBJECT("user_id", user_id, "email", email, "nickname", nickname)), ']') AS players
		FROM game g
		INNER JOIN game_player gp 
			ON gp.fk_game = g.game_id
		INNER JOIN user u 
			ON u.user_id = gp.fk_user
		WHERE 
			g.game_start = 0
		GROUP BY 
			game_id, name;`

	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
	}

	games := make([]*SaveGame, 0)
	for rows.Next() {
		game := &SaveGame{}
		players := "[]"
		err := rows.Scan(&game.ID, &game.Name, &players)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		err = json.Unmarshal([]byte(players), &game.Players)
		if err != nil {
			panic(err) // This would normally be a normal Error http response but I've put this here so it's easy for you to test.
		}

		games = append(games, game)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return &games, nil
}

//CreateGame creates a new game in the database
func (g *SaveGame) CreateGame() error {
	db := CreateDb()
	defer db.closeConnection()

	//creates game
	sql := `
		CALL create_game(?, ?, ?, ?, ?, ?);`

	row, err := db.Query(sql,
		g.Name,
		g.GameSettings.Version,
		g.GameSettings.NumberOfPlayers,
		g.GameSettings.InvestorCard,
		g.GameSettings.TaxIncreaseOnlyBonus,
		g.GameSettings.StartingMode)
	if err != nil {
		return err
	}
	var gameID int64
	row.Next()
	err = row.Scan(&gameID)
	if err != nil {
		log.Println(err)
		return err
	}

	//add players
	players := make([]interface{}, len(g.Players))
	for i, p := range g.Players {
		players[i] = p.PlayerID
	}
	repeat := ""
	if len(players) > 1 {
		repeat = strings.Repeat(",?", len(players)-1)
	}
	sql = `
		INSERT INTO game_player (fk_game, fk_user)
		SELECT 
			?, 
			user_id  
		FROM 
			user
		WHERE
			user_id IN (?` + repeat + `)`
	args := []interface{}{gameID}
	args = append(args, players...)
	_, err = db.InsertAndReturn(sql, args...)
	if err != nil {
		return err
	}

	return nil
}
