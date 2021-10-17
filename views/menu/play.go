package menu

import tb "gopkg.in/tucnak/telebot.v2"

func NewPlay() MenuHandle {
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

	pm := BaseMenu{}
	pm.rm = p
	pm.msg = "<b>GAME IS STARTING</b>\n Join the game to be able to play.\nWhen ready click the Play button."
	return pm
}
