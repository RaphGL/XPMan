package models

import (
	"database/sql"
)

type Player struct {
	XPPoints   int
	GamesWon   int
	GamesLost  int
	LastPlayed string
}

func CheckAndAddPlayerAcc(userID int, db *sql.DB) {
	if db.QueryRow("SELECT telegram_uid FROM players where telegram_uid = ?", userID).Err() != nil {
		db.Exec("INSERT INTO players (telegram_uid) VALUES (?)", userID)
	}
}

func UserInfo(userID int, db *sql.DB) (Player, error) {
	CheckAndAddPlayerAcc(userID, db)
	player := Player{}
	db.QueryRow(`
		SELECT xp_points, games_won, games_lost, last_played FROM players WHERE telegram_uid = ?;
	`, userID).Scan(&player.XPPoints, &player.GamesWon, &player.GamesLost, &player.LastPlayed)
	return player, nil
}
