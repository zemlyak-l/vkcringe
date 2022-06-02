package bot

import (
	"github.com/zemlyak-l/vkgottle/api"
)

type Bot struct {
	Api *api.Api
}

func NewBot(token string) *Bot {
	return &Bot{
		Api: api.NewApi(token),
	}
}
