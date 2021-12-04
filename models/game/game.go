package game

import (
	"database/sql"
	"errors"
	"fmt"

	tb "gopkg.in/tucnak/telebot.v2"

	_ "github.com/go-sql-driver/mysql"
)

// TODO function needs to return to tell if a game is current
func CreateGame(c *tb.Callback, db *sql.DB) {
	if db.QueryRow(`
		SELECT host_id, chat_id FROM active_games
		WHERE host_id = ? AND chat_id = ?;
	`, c.Sender.ID, c.Message.Chat.ID).Err() != nil {
	}

	_, err := db.Exec(`
		INSERT INTO active_matches (chat_id, participant_uid, part_username)
		VALUES (?, ?, '?');
	`, c.Message.Chat.ID, c.Sender.ID, c.Sender.Username)
	if err != nil {
		fmt.Println("Couldn't create active_matches entry")
	}
	_, err = db.Exec(`
		INSERT INTO active_games (chat_id, host_id)
		VALUES (?, ?);
	`, c.Message.Chat.ID, c.Sender.ID)
	if err != nil {
		fmt.Println("Couldn't create active_games entry")
	}
}

func RegisterParticipant(c *tb.Callback, db *sql.DB) error {
	_, err := db.Exec(`
	INSERT INTO active_matches (part_username, chat_id, participant_uid)
	VALUES ('?', ?, ?);
	`, c.Sender.Username, c.Message.Chat.ID, c.Sender.ID)
	if err != nil {
		return errors.New("Could not register new participant.")
	}
	return nil
}

func GetParticipants(c *tb.Callback, db *sql.DB) []string {
	rows, err := db.Query(`
		SELECT part_username FROM active_matches WHERE chat_id = ?;
	`, c.Message.Chat.ID)
	if err != nil {
		return nil
	}

	var part string
	var parts []string
	for rows.Next() {
		rows.Scan(&part)
		parts = append(parts, part)
	}
	return parts
}

func RemoveParticipant() {}
func StartGame()         {}
