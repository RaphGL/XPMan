package menu

import tb "gopkg.in/tucnak/telebot.v2"

// Basic functions of all menu views
type MenuHandle interface {
	HCallback(*tb.Callback)
	GetMsgID() string
	GetMsgContent() string
	GetReplyMarkup() *tb.ReplyMarkup
}

// Attributes of a menu state
type BaseMenu struct {
	// Message content
	msg string
	// Message reply markup
	rm *tb.ReplyMarkup
	// embed the menuhandle interface
	MenuHandle
}

// Get message text
func (b BaseMenu) GetMsgContent() string {
	return b.msg
}

// Get the message's reply markup
func (b BaseMenu) GetReplyMarkup() *tb.ReplyMarkup {
	return b.rm
}
