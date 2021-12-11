package controllers

import (
	"container/list"
	"database/sql"
	"fmt"
	"strings"

	"github.com/RaphGL/XPMan/models/game"
	"github.com/RaphGL/XPMan/views/menu"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Menu message struct with state management methods
type menuMsg struct {
	senderMsg *tb.Message     // Received message struct
	bot       *tb.Bot         // Running bot instance
	cView     menu.MenuHandle // Current view
	pView     *list.List      // Previous views in stack
}

// Retuns a menuMsg struct
func newMenuMsg(b *tb.Bot, m *tb.Message) menuMsg {
	return menuMsg{
		senderMsg: m,
		bot:       b,
		pView:     list.New(),
		cView:     menu.NewMainMenu(),
	}
}

// Retrieve the current menu view
func (mm *menuMsg) currState() menu.MenuHandle {
	return mm.cView
}

// Set current menu view struct
func (mm *menuMsg) setCurrState(ms menu.MenuHandle) {
	mm.pView.PushFront(mm.currState())
	mm.cView = ms
}

// Go back to a previous menu view state
func (mm *menuMsg) setPrevState() {
	s := mm.pView.Front()
	mm.pView.PushFront(mm.currState())
	defer mm.pView.Remove(s)

	pState, ok := (s.Value).(menu.MenuHandle)
	if ok {
		mm.setCurrState(pState)
	}
}

// Updates message content when state is changed
func (mm *menuMsg) updateMsg(c *tb.Callback) {
	mm.bot.Edit(c.Message, mm.cView.MsgContent(), mm.cView.ReplyMarkup(), tb.ModeHTML)
}

// Sends message to the user's chat
func (mm *menuMsg) sendMenuMsg() {
	mm.bot.Send(mm.senderMsg.Chat, mm.cView.MsgContent(), mm.cView.ReplyMarkup(), tb.ModeHTML)
}

// control how menu msg is displayed
func ControlMenuDisplay(b *tb.Bot, db *sql.DB) {
	b.Handle("/start", func(m *tb.Message) {
		msg := newMenuMsg(b, m)
		msg.sendMenuMsg()
		var host string

		b.Handle(tb.OnCallback, func(c *tb.Callback) {
			// cleans up LF and CR characters from c.Data
			switch strings.TrimSpace(c.Data) {
			case "mainmenu|back":
				msg.setPrevState()

			case "mainmenu|play":
				host = c.Sender.Username
				msg.setCurrState(menu.NewPlay(c.Sender.Username, []string{}))
				game.Create(c, db)

			case "mainmenu|profile":
				p, err := menu.NewProfile(c, db)
				if err != nil {
					errmsg := fmt.Sprintf("Internal error: no user `%s` in database. Please report error to developers.", c.Sender.Username)
					b.Send(c.Message.Chat, errmsg)
				}
				b.Send(c.Message.Chat, p.MsgContent(), p.ReplyMarkup(), tb.ModeHTML)

			case "mainmenu|ranking":

			case "mainmenu|how_to_play":
				msg.setCurrState(menu.NewTutorial())

			case "mainmenu|join_game":
				game.RegisterParticipant(c, db)
				// make sure that message doesn't get stuck on updating participants
				msg.setPrevState()
				msg.setCurrState(menu.NewPlay(host, game.GetParticipants(c, db)))

			case "mainmenu|quit_game":
				game.RemoveParticipant(c, db)
				msg.setPrevState()
				msg.setCurrState(menu.NewPlay(host, game.GetParticipants(c, db)))

			case "mainmenu|start_game":
				game.Start(c, db)
			}

			b.Respond(c)
			msg.updateMsg(c)
		})
	})
}
