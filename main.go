package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/raphgl/telegrambot/models"
	"github.com/raphgl/telegrambot/views"
	tb "gopkg.in/tucnak/telebot.v2"
)

// JSON for the config file
type Config struct {
	APIKey        string `json:"api_key"`
	DatabaseSetup string `json:"database_setup"`
}

func main() {
	// read json file and load it to struct
	jsonFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
	}

	var configData Config
	err = json.Unmarshal(jsonFile, &configData)
	if err != nil {
		panic(err)
	}

	// initialize database
	db, err := sql.Open("mysql", configData.DatabaseSetup)
	if err != nil {
		panic(err)
	}

	// initialize telegram bot
	b, err := tb.NewBot(tb.Settings{
		Token:  configData.APIKey,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		panic(err)
	}

	// set up database
	models.SetupDB(db)
	models.PopulateDictDB(db)

	// show start menu
	b.Handle("/start", func(m *tb.Message) {
		//views.ShowMainMenu(b, m)
		views.ShowGameInstructions(b, m)
		b.Handle(tb.OnCallback, func(c *tb.Callback) {
			var res string
			// cleans up LF and CR characters from c.Data
			switch strings.TrimSpace(c.Data) {
			case "play":
				res = "Loading Game"
			case "profile":
				res = "Loading User Profile"
			case "balance":
				res = "Checking Balance"
			case "ranking":
				res = "Loading Ranking"
			case "how_to_play":
				res = "Loading Tutorial"
			}
			b.Respond(c, &tb.CallbackResponse{
				Text: res,
			})
		})
	})

	b.Start()
}
