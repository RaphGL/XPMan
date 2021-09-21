package models

import (
	"database/sql"
	"encoding/csv"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type dictData struct {
	word       string
	tip1       string
	tip2       string
	tip3       string
	difficulty string
}

// Creates tables used by the bot
// db = variable with the opened database
func SetupDB(db *sql.DB) {
	db.Exec(`
		CREATE TABLE IF NOT EXISTS players (
			player_id INT NOT NULL AUTO_INCREMENT,
			telegram_uid VARCHAR(25) NOT NULL,
			level INT NOT NULL,
			xp_points INT NOT NULL,
			games_played INT NOT NULL,
			PRIMARY KEY ( player_id )
		);
	`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS active_games (
			chat_id INT NOT NULL, 
			num_participants INT NOT NULL,
			PRIMARY KEY ( chat_id )
		);
	`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS active_matches (
			chat_id INT NOT NULL AUTO_INCREMENT,
			participant_uid VARCHAR(25) NOT NULL,
			attempts_available INT NOT NULL,
			ncoins INT NOT NULL,
			guessed_right VARCHAR(25),
			guessed_wrong VARCHAR(25),
			PRIMARY KEY ( chat_id )
		);
	`)
	db.Exec(`
		CREATE TABLE IF NOT EXISTS game_dictionary (
			word VARCHAR(20) NOT NULL,
			tip1 VARCHAR(50) NOT NULL,
			tip2 VARCHAR(50) NOT NULL,
			tip3 VARCHAR(50) NOT NULL,
			difficulty CHAR(5) NOT NULL,
			PRIMARY KEY ( word )
		)
	`)
}

// Populates game_dictionary with all the playable words
// db = variable with opened database
func PopulateDictDB(db *sql.DB) {
	// open csv file with the game's dictionary
	csvFile, err := os.Open("game_dictionary.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	csvData, err := csv.NewReader(csvFile).ReadAll()

	// add table rows to database
	for _, row := range csvData {
		_, err := db.Exec("INSERT IGNORE INTO game_dictionary (word, tip1, tip2, tip3, difficulty) VALUES (?, ?, ?, ?, ?)",
			row[0], row[1], row[2], row[3], row[4])
		if err != nil {
			panic(err)
		}
	}
}