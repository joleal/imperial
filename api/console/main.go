package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joleal/imperial/api/engine"
)

func main() {

	console()

	/*
		helloHandler := func(w http.ResponseWriter, req *http.Request) {
			io.WriteString(w, "Imperial\n")
		}

		http.HandleFunc("/", helloHandler)
		log.Println("Listing for requests at http://localhost:3333/hello")
		log.Fatal(http.ListenAndServe(":3333", nil))
	*/
}

func console() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Tester")
	fmt.Println("---------------------")

	var game *engine.Game
	settings := engine.Settings{
		Version:         engine.I2030,
		NumberOfPlayers: 4,
		StartingMode:    engine.Random,
		InvestorCard:    true,
	}
	users := []engine.User{
		{Name: "Joao"},
		{Name: "Pedro"},
		{Name: "Luos"},
		{Name: "Paulo"},
	}
Loop:
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

		i, err := strconv.Atoi(text)
		if err == nil {
			settings.NumberOfPlayers = i
		}

		switch text {
		case "game":
			game = &engine.Game{}
			game.CreateGame(settings)
			game.StartGame(users)
			fmt.Println(game.Status())
		case "country":
			fmt.Println(game.Country())
		case "map":
			fmt.Println(game.Map())
		case "player":
			fmt.Println(game.NextPlayer())
		case "status":
			fmt.Println(game.Status())
		case "random":
			rand.Seed(time.Now().UnixNano())
			fmt.Println(rand.Intn(10))
		case "exit":
			fmt.Println("bye...")
			break Loop
		}
	}
}
