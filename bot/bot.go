package bot

import (
	"strings"

	"github.com/zemlyak-l/vkgottle/api"
	"github.com/zemlyak-l/vkgottle/object"
	"github.com/zemlyak-l/vkgottle/polling"
)

type Bot struct {
	Api        *api.Api
	Longpoll   *polling.Longpoll
	Routes     *polling.Routes
	TextRoutes *TextRoutes
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
		Api:        api,
		Longpoll:   lp,
		Routes:     lp.Routes,
		TextRoutes: NewTextRoutes(),
	}
	bot.Routes.MessageNew = bot.messageHandler

	return bot, nil
}

func (bot *Bot) RunSync() {
	bot.Longpoll.ListenNewEvents()
}

func (bot *Bot) messageHandler(message object.NewMessage) {
	cmdArgs := strings.Split(message.Text, " ")
	cmdName := cmdArgs[0]
	if len(cmdArgs) != 1 {
		message.CmdArgs = cmdArgs[1:]
	}
	f, ok := bot.TextRoutes.AllRoutes[cmdName]
	if ok {
		f(message)
		return
	}
	if message.PeerID < 2000000000 {
		pf, ok := bot.TextRoutes.PrivateRoutes[cmdName]
		if ok {
			pf(message)
			return
		}
	} else {
		cf, ok := bot.TextRoutes.ChatRoutes[cmdName]
		if ok {
			cf(message)
			return
		}
	}
	if bot.TextRoutes.AllHandler != nil {
		bot.TextRoutes.AllHandler(message)
	}
}
