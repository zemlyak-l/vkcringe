package object

type UsersGetObject struct {
	UserIDs  string `schema:"user_ids"`
	Fields   string `schema:"fields"`
	NameCase string `schema:"name_case"`
}

type UsersGetResponse struct {
	UsersList []User `json:"response"`
}
