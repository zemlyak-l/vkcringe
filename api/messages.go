package api

import (
	"github.com/zemlyak-l/vkcringe/object"
)

func (api *Api) MessagesSend(message *object.Message) error {
	return api.Request("messages.send", message, nil)
}
