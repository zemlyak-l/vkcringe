package object

type FriendsRequestsMutual struct {
	Count int   `json:"count"` // Total mutual friends number
	Users []int `json:"users"`
}
