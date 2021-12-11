package views

import tb "gopkg.in/tucnak/telebot.v2"

// Basic functions of all menu views
type MenuHandle interface {
	HCallback(*tb.Callback)
	MsgContent() string
	ReplyMarkup() *tb.ReplyMarkup
}

// Attributes of a menu state
type BaseMenu struct {
	// Message content
	Msg string
	// Message reply markup
	Rm *tb.ReplyMarkup
	// embed the menuhandle interface
	MenuHandle
}

// Get message text
func (b BaseMenu) MsgContent() string {
	return b.Msg
}

// Get the message's reply markup
func (b BaseMenu) ReplyMarkup() *tb.ReplyMarkup {
	return b.Rm
}
