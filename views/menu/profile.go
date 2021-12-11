package menu

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/RaphGL/XPMan/models"
	"github.com/RaphGL/XPMan/views"
	tb "gopkg.in/tucnak/telebot.v2"
)

func NewProfile(c *tb.Callback, db *sql.DB) (views.MenuHandle, error) {
	res, err := models.UserInfo(c.Sender.ID, db)
	if err != nil {
		return nil, fmt.Errorf("Could not get profile for %s", c.Sender.Username)
	}
	p := views.BaseMenu{}
	p.Rm = &tb.ReplyMarkup{}
	p.Msg = fmt.Sprintln(
		fmt.Sprintf("%s%s%s", "<b>", strings.ToUpper(c.Sender.Username), "</b>"), "\n",
		"<b>Level:</b>", res.XPPoints/100, "\n",
		"<b>Games Won:</b>", res.GamesWon, "\n",
		"<b>Games Lost</b>:", res.GamesLost, "\n",
		"<b>Last Played:</b>", res.LastPlayed)
	return p, nil
}
