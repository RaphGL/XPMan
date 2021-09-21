package views

import tb "gopkg.in/tucnak/telebot.v2"

func ShowGameInstructions(b *tb.Bot, m *tb.Message) {
	gi := &tb.ReplyMarkup{}
	gi.Inline(
		gi.Row(
			gi.Data("◀️ Back", "back"),
		),
	)

	b.Send(m.Sender, `
	<b>How To Play</b>

	XPManBot is a hangman inspired social game where players try to gain as much XP as possible by guessing the letters of a given word. Players can use coins, timing and perks in order to gain advantage over the adversaries. The player with the most XP at the end of the game wins.
	`, gi, "html")
}
