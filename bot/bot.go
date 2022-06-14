package bot

import (
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
		TextRoutes: NewRoutes(),
	}
	bot.Routes.MessageNew = bot.messageHandler

	return bot, nil
}

func (bot *Bot) RunSync() {
	bot.Longpoll.ListenNewEvents()
}
