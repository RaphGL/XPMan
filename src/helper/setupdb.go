package helper

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type dictData struct {
	word       string
	tip1       string
	tip2       string
	tip3       string
	difficulty string
}

func SetupDB(db *sql.DB) {
	db.Exec(`
		CREATE TABLE players (
			player_id INT NOT NULL AUTO_INCREMENT,
			telegram_uid VARCHAR(25) NOT NULL,
			level INT NOT NULL,
			xp_points INT NOT NULL,
			games_played INT NOT NULL,
			PRIMARY KEY ( player_id )
		);
	`)
	db.Exec(`
		CREATE TABLE active_games (
			chat_id INT NOT NULL, 
			num_participants INT NOT NULL,
			PRIMARY KEY ( chat_id )
		);
	`)
	db.Exec(`
		CREATE TABLE active_matches (
			chat_id INT NOT NULL AUTO_INCREMENT,
			participant_uid VARCHAR(25) NOT NULL,
			attempts_available INT NOT NULL,
			ncoins INT NOT NULL,
			guessed_right VARCHAR(25),
			guessed_wrong VARCHAR(25),
			PRIMARY KEY ( chat_id )
		);
	`)
}

func PopulateDictDB() {

}
