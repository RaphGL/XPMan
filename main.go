package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		views.ShowMainMenu(b, m)
		b.Handle(tb.OnCallback, func(c *tb.Callback) {
			b.Respond(c, &tb.CallbackResponse{})
		})
	})

	b.Start()
}
