package models

import (
	"database/sql"
	"encoding/csv"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Creates tables used by the bot
// db = variable with the opened database
func SetupDB(db *sql.DB) {
	db.Exec(`
		CREATE TABLE IF NOT EXISTS players (
			telegram_uid VARCHAR(25) NOT NULL,
			xp_points INT NOT NULL DEFAULT 0,
			games_won INT NOT NULL DEFAULT 0,
			games_lost INT NOT NULL DEFAULT 0,
			last_played DATE DEFAULT CURRENT_DATE(),
			PRIMARY KEY ( telegram_uid )
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
	if err != nil {
		panic("Invalid csv file")
	}

	// add table rows to database
	for _, row := range csvData {
		_, err := db.Exec("INSERT IGNORE INTO game_dictionary (word, tip1, tip2, tip3, difficulty) VALUES (?, ?, ?, ?, ?)",
			row[0], row[1], row[2], row[3], row[4])
		if err != nil {
			panic(err)
		}
	}
}
