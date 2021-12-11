package game

import (
	"database/sql"
	"errors"
	"fmt"

	tb "gopkg.in/tucnak/telebot.v2"

	_ "github.com/go-sql-driver/mysql"
)

// TODO function needs to return to tell if a game is current
func Create(c *tb.Callback, db *sql.DB) {
	// checks if game is already created
	if db.QueryRow(`
		SELECT host_id, chat_id FROM active_games
		WHERE host_id = ? AND chat_id = ?;
	`, c.Sender.ID, c.Message.Chat.ID).Err() != nil {
		return
	}

	// if game not create, create it
	_, err := db.Exec(`
		INSERT INTO active_matches (chat_id, participant_uid, part_username)
		VALUES (?, ?, ?);
	`, c.Message.Chat.ID, c.Sender.ID, c.Sender.Username)
	if err != nil {
		//fmt.Println("Couldn't create active_matches entry")
		fmt.Println(err)
	}
	_, err = db.Exec(`
		INSERT INTO active_games (chat_id, host_id)
		VALUES (?, ?);
	`, c.Message.Chat.ID, c.Sender.ID)
	if err != nil {
		//fmt.Println("Couldn't create active_games entry")
		fmt.Println(err)
	}
}

// TODO Rewrite Registration
func RegisterParticipant(c *tb.Callback, db *sql.DB) error {
	var exists bool
	// checks if player already exists in database
	err := db.QueryRow(`
		SELECT EXISTS(SELECT chat_id, participant_uid FROM active_matches 
		WHERE chat_id = ? AND participant_uid = ?);
	`, c.Message.Chat.ID, c.Sender.ID).Scan(&exists)
	if err != nil {
		return err
		// if user doesn't exist it adds them to the active_matches table
	} else if !exists {
		_, err = db.Exec(`
			INSERT INTO active_matches (chat_id, part_username, participant_uid) 
			VALUES (?, ?, ?);
		`, c.Message.Chat.ID, c.Sender.Username, c.Sender.ID)
		if err != nil {
			return err
		}
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

// Removes user from active match
func RemoveParticipant(c *tb.Callback, db *sql.DB) error {
	_, err := db.Exec(`
		DELETE FROM active_matches WHERE chat_id = ? AND participant_uid = ?;
	`, c.Message.Chat.ID, c.Sender.ID)
	if err != nil {
		return err
	}
	return nil
}

// Picks a word at to be played
func Start(c *tb.Callback, db *sql.DB) error {
	var hostID int
	err := db.QueryRow(`
		SELECT host_id FROM active_games WHERE chat_id = ?;
	`, c.Message.Chat.ID).Scan(&hostID)
	if err != nil {
		return err
	}

	if hostID == c.Sender.ID {
		var word string
		err = db.QueryRow(`
			SELECT word FROM game_dictionary ORDER BY RAND() LIMIT 1;
		`).Scan(&word)
		if err != nil {
			return err
		}

		_, err = db.Exec(`
			UPDATE active_games SET word = ? WHERE chat_id = ?;
		`, word, c.Message.Chat.ID)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("Callback is not from host")
}

func SetGameScore() {}
