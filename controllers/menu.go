package controllers

import (
	"container/list"
	"strings"

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
func (mm *menuMsg) getCurrState() menu.MenuHandle {
	return mm.cView
}

// Set current menu view struct
func (mm *menuMsg) setCurrState(ms menu.MenuHandle) {
	mm.pView.PushFront(mm.getCurrState())
	mm.cView = ms
}

// Go back to a previous menu view state
func (mm *menuMsg) setPrevState() {
	s := mm.pView.Front()
	mm.pView.PushFront(mm.getCurrState())
	defer mm.pView.Remove(s)

	pState, ok := (s.Value).(menu.MenuHandle)
	if ok {
		mm.setCurrState(pState)
	}
}

func ControlMenuDisplay(b *tb.Bot) {
	b.Handle("/start", func(m *tb.Message) {
		msg := newMenuMsg(b, m)
		b.Send(m.Chat, msg.cView.GetMsgContent(), msg.cView.GetReplyMarkup(), tb.ModeHTML)

		b.Handle(tb.OnCallback, func(c *tb.Callback) {
			// cleans up LF and CR characters from c.Data
			switch strings.TrimSpace(c.Data) {
			case "ranking":
				b.Send(m.Chat, "fuck you")
			case "how_to_play":
				b.Send(m.Chat, "test")
			case "back":
				msg.setPrevState()
			}
		})
	})
}
