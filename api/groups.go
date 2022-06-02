package api

import (
	"github.com/zemlyak-l/vkgottle/object"
)

func (api *Api) GroupsGetLongPollServer(
	message object.GetServerMessage, target *object.ResponseInit,
) error {
	return api.Request("groups.getLongPollServer", message, target)
}
