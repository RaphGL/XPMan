package menu

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

type mainMenu struct {
	BaseMenu
}

func NewMainMenu() MenuHandle {
	menu := &tb.ReplyMarkup{}
	menu.Inline(
		menu.Row(
			menu.Data("▶️ Play", "mainmenu", "play"),
		),
		menu.Row(
			menu.Data("👤 View Profile", "mainmenu", "profile"),
			menu.Data("🏦 View Balance", "mainmenu", "balance"),
			menu.Data("🏅 View Ranking", "mainmenu", "ranking"),
		),
		menu.Row(
			menu.Data("📚 Learn How To Play", "mainmenu", "how_to_play"),
			menu.URL("💻 Contribute Code", "https://github.com/RaphGL/XPMan"),
		),
	)

	m := mainMenu{}
	m.rm = menu
	m.msg = `
	Welcome to the <b>XPManBot</b>.
	What would you like to do?
	`
	m.msgID = "mainmenu"
	return m
}
