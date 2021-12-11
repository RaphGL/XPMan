package menu

import (
	"github.com/RaphGL/XPMan/views"
	tb "gopkg.in/tucnak/telebot.v2"
)

func NewTutorial() views.MenuHandle {
	// game instruction message
	gi := &tb.ReplyMarkup{}
	gi.Inline(
		gi.Row(
			gi.Data("◀️ Back", "mainmenu", "back"),
		),
	)
	htp := views.BaseMenu{}
	htp.Rm = gi
	htp.Msg = `<b>HOW TO PLAY</b>
	XPMan is a hangman like game. You're given a word to guess. People will take turns in guessing a letter of the word. For each correct answer you earn a XP and coins and are allowed to either continue playing or pay coins to skip your turn, player with the most XP in the end wins.
	
	<b><i>Coins</i></b>
	The coins are used to gain advantage over others by buying perks.
	Some of the available perks are: skipping turns, target a penalty to a user, get tips...

	<b><i>Experience Points</i></b>
	One gains or loses XP based on where they got an answer right or wrong or whether they skipped a turn. Managing how one acts in-game will affect one's ranking.
	`
	return htp
}
