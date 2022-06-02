package object

type NewGetByID struct {
	GroupIDs string `schema:"group_ids,omitempty"`
	GroupID  int    `schema:"group_id,omitempty"`
	Fields   string `schema:"fields,omitempty"`
}

type GroupResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ScreenName  string `json:"screen_name"`
	IsClosed    int    `json:"is_closed"`
	Deactivated string `json:"deactivated"`
	IsAdmin     int    `json:"is_admin"`
	AdminLevel  int    `json:"admin_level"`
	IsMember    int    `json:"is_member"`
	IsAdvestier int    `json:"is_advertiser"`
	InvitedBy   int    `json:"invited_by"`
	Type        string `json:"type"`
	Photo50     string `json:"photo_50"`
	Photo100    string `json:"photo_100"`
	Photo200    string `json:"photo_200"`
}

type GroupResponseSlice struct {
	Response []GroupResponse `json:"response"`
}
