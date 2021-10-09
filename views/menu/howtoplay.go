package menu

import tb "gopkg.in/tucnak/telebot.v2"

type howtoplay struct {
	BaseMenu
}

func NewTutorial() MenuHandle {
	// game instruction message
	gi := &tb.ReplyMarkup{}
	gi.Inline(
		gi.Row(
			gi.Data("◀️ Back", "back"),
		),
	)
	htp := howtoplay{}
	htp.msg = `<b>How To Play</b>
	XPManBot is a hangman inspired social game where players try to gain as much XP as possible by guessing the letters of a given word. Players can use coins, timing and perks in order to gain advantage over the adversaries. The player with the most XP at the end of the game wins.`
	htp.msgID = "how_to_play"
	return htp
}
