package views

import tb "gopkg.in/tucnak/telebot.v2"

func ShowMainMenu(b *tb.Bot, m *tb.Message) {
	menu := &tb.ReplyMarkup{}
	menu.Inline(
		menu.Row(
			menu.Data("▶️ Play", "play"),
		),
		menu.Row(
			menu.Data("👤 View Profile", "profile"),
			menu.Data("🏦 View Balance", "balance"),
			menu.Data("🏅 View Ranking", "ranking"),
		),
		menu.Row(
			menu.Data("📚 How To Play", "how_to_play"),
			menu.URL("💻 Contribute Code", "https://github.com/RaphGL/XPMan"),
		),
	)
	b.Send(m.Sender, `
	Welcome to the <b>XPManBot</b>.
	What would you like to do?
	`, menu, "html")
}
