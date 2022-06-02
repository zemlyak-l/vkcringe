package api

import (
	"github.com/zemlyak-l/vkgottle/object"
)

func (api *Api) MessagesSend(message *object.Message) error {
	return api.Request("messages.send", message, nil)
}
