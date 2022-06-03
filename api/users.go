package api

import "github.com/zemlyak-l/vkgottle/object"

func (api *Api) UsersGet(
	user object.UsersGetObject, target *object.UsersGetResponse,
) error {
	return api.Request("users.get", user, target)
}
