package menu

import (
	"github.com/RaphGL/XPMan/views"
	tb "gopkg.in/tucnak/telebot.v2"
)

func NewMainMenu() views.MenuHandle {
	menu := &tb.ReplyMarkup{}
	menu.Inline(
		menu.Row(
			menu.Data("▶️ Play", "mainmenu", "play"),
		),
		menu.Row(
			menu.Data("👤 View Profile", "mainmenu", "profile"),
			menu.Data("🏅 View Ranking", "mainmenu", "ranking"),
		),
		menu.Row(
			menu.Data("📚 Learn How To Play", "mainmenu", "how_to_play"),
			menu.URL("💻 Contribute Code", "https://github.com/RaphGL/XPMan"),
		),
	)

	m := views.BaseMenu{}
	m.Rm = menu
	m.Msg = `
	Welcome to the <b>XPManBot</b>. 
	This bot is currently in beta and may not function properly. 
	If you wish to report any issues with the bot or suggest new gaming modes. Please click on <i>Contribute Code</i>

	<b>What would you like to do?</b>
	`
	return m
}
