package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/raphgl/telegrambot/src/helper"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

	// initialize telegram bot
	bot, err := tgbotapi.NewBotAPI(configData.APIKey)
	bot.Debug = false
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	// initialize database
	db, err := sql.Open("mysql", configData.DatabaseSetup)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	helper.SetupDB(db)

	// message receiver
	for u := range updates {
		if u.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(u.Message.Chat.ID, u.Message.Text)
		msg.ReplyToMessageID = u.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}

}
