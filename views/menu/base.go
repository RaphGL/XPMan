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
	msg   string
	rm    *tb.ReplyMarkup
	msgID string
	MenuHandle
}

func (b BaseMenu) GetMsgContent() string {
	return b.msg
}

func (b BaseMenu) GetMsgID() string {
	return b.msgID
}

func (b BaseMenu) GetReplyMarkup() *tb.ReplyMarkup {
	return b.rm
}
