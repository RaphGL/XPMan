package views

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func ShowUserPolling(b *tb.Bot, m *tb.Message) (*tb.Message, error) {
	p := &tb.ReplyMarkup{}
	p.Inline(
		p.Row(
			p.Data("📥 Join", "userpolling", "join_game"),
		),
		p.Row(
			p.Data("📤 Quit", "userpolling", "quit_game"),
		),
	)

	return b.Send(m.Chat, "Click on `Join` to join the game.", p, tb.ModeMarkdown)
}
