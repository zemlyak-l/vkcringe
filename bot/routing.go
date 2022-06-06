package bot

import "github.com/zemlyak-l/vkgottle/object"

type TextRoutes struct {
	AllRoutes     map[string]func(object.NewMessage)
	PrivateRoutes map[string]func(object.NewMessage)
	ChatRoutes    map[string]func(object.NewMessage)
	AllHandler    func(message object.NewMessage)
}

func NewTextRoutes() *TextRoutes {
	return &TextRoutes{
		AllRoutes:     make(map[string]func(object.NewMessage)),
		PrivateRoutes: make(map[string]func(object.NewMessage)),
		ChatRoutes:    make(map[string]func(object.NewMessage)),
	}
}

func (bot *Bot) OnAll(f func(message object.NewMessage)) {
	bot.TextRoutes.AllHandler = f
}

func (bot *Bot) OnMessage(text string, f func(message object.NewMessage)) {
	bot.TextRoutes.AllRoutes[text] = f
}

func (bot *Bot) OnPrivateMessage(text string, f func(message object.NewMessage)) {
	bot.TextRoutes.PrivateRoutes[text] = f
}

func (bot *Bot) OnChatMessage(text string, f func(message object.NewMessage)) {
	bot.TextRoutes.ChatRoutes[text] = f
}
