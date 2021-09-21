package views

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

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
 	XPManBot is a hangman inspired social game where players try to gain as much XP as possible by guessing the letters of a given word. Players can use coins, timing and perks in order to gain advantage over the adversaries. The player with the most XP at the end of the game wins.

 	<i>You can learn more about how the game works by clicking on How To Play.</i>
	`, menu, "html")
}
