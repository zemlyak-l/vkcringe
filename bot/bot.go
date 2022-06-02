package bot

import (
	"github.com/zemlyak-l/vkgottle/api"
	"github.com/zemlyak-l/vkgottle/object"
	"github.com/zemlyak-l/vkgottle/polling"
)

type Bot struct {
	Api           *api.Api
	Longpoll      *polling.Longpoll
	Routes        map[string]func(object.NewMessage)
	PrivateRoutes map[string]func(object.NewMessage)
	ChatRoutes    map[string]func(object.NewMessage)
	AllHandler    func(object.NewMessage)
}

func NewBot(token string) (*Bot, error) {
	api := api.NewApi(token)

	group := &object.GroupResponseSlice{}
	req := object.NewGetByID{Fields: "description"}
	if err := api.GroupsGetByID(req, group); err != nil {
		return nil, err
	}
	groupID := group.Response[0].ID

	lp, err := polling.NewLongpoll(api, groupID)
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		Api:           api,
		Longpoll:      lp,
		Routes:        make(map[string]func(object.NewMessage)),
		PrivateRoutes: make(map[string]func(object.NewMessage)),
		ChatRoutes:    make(map[string]func(object.NewMessage)),
	}
	bot.Longpoll.Routes.MessageNew = bot.messageHandler

	return bot, nil
}

func (bot *Bot) messageHandler(message object.NewMessage) {
	f, ok := bot.Routes[message.Text]
	if ok {
		f(message)
		return
	}
	if message.PeerID < 2000000000 {
		pf, ok := bot.PrivateRoutes[message.Text]
		if ok {
			pf(message)
			return
		}
	} else {
		cf, ok := bot.ChatRoutes[message.Text]
		if ok {
			cf(message)
			return
		}
	}
	if bot.AllHandler != nil {
		bot.AllHandler(message)
	}
}

func (bot *Bot) OnMessage(text string, f func(object.NewMessage)) {
	bot.Routes[text] = f
}

func (bot *Bot) OnPrivateMessage(text string, f func(object.NewMessage)) {
	bot.PrivateRoutes[text] = f
}

func (bot *Bot) OnChatMessage(text string, f func(object.NewMessage)) {
	bot.ChatRoutes[text] = f
}
