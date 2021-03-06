package menu

import (
	"fmt"

	"github.com/RaphGL/XPMan/views"
	tb "gopkg.in/tucnak/telebot.v2"
)

func NewPlay(gameHost string, participants []string) views.MenuHandle {
	// join game message
	p := &tb.ReplyMarkup{}
	p.Inline(
		p.Row(
			p.Data("▶️ Start game", "mainmenu", "start_game"),
		),
		p.Row(
			p.Data("🕹️ Join game", "mainmenu", "join_game"),
			p.Data("🚶 Quit game", "mainmenu", "quit_game"),
		),
		p.Row(
			p.Data("◀️ Back", "mainmenu", "back"),
		),
	)

	pm := views.BaseMenu{}
	pm.Rm = p
	pm.Msg = fmt.Sprintf("<b>GAME IS STARTING</b>\n Join the game to be able to play.\nWhen ready, <b>%s</b> should click the Play button.", gameHost)
	pm.Msg += "\n\nParticipants:\n"
	for _, part := range participants {
		pm.Msg += "- <b>" + part + "</b>\n"
	}

	return pm
}
