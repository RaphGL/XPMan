package menu

import (
	"fmt"
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

func NewProfile(username string, XP int, gamesWon int, gamesLost int, lastPlayed string) MenuHandle {
	p := BaseMenu{}
	p.rm = &tb.ReplyMarkup{}
	p.msg = fmt.Sprintln(
		fmt.Sprintf("%s%s%s", "<b>", strings.ToUpper(username), "</b>"), "\n",
		"<b>Level:</b>", XP/100, "\n",
		"<b>Games Won:</b>", gamesWon, "\n",
		"<b>Games Lost</b>:", gamesLost, "\n",
		"<b>Last Played:</b>", lastPlayed)
	return p
}
