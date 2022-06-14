package object

import (
	"bytes"
	"errors"
)

type Bool bool

func (b *Bool) UnmarshalJSON(data []byte) (err error) {
	switch {
	case bytes.Equal(data, []byte("1")) || bytes.Equal(data, []byte("true")):
		*b = true

	case bytes.Equal(data, []byte("0")) || bytes.Equal(data, []byte("false")):
		*b = false

	default:
		return errors.New("Not a Bool")
	}

	return
}

type LikesInfo struct {
	CanLike    Bool `json:"can_like"`
	CanPublish Bool `json:"can_publish"`
	UserLikes  Bool `json:"user_likes"`
	Count      int  `json:"count"`
}

type Link struct {
	Application  LinkApplication `json:"application"`
	Button       LinkButton      `json:"button"`
	ButtonText   string          `json:"button_text"`
	ButtonAction string          `json:"button_action"`
	Caption      string          `json:"caption"`
	Description  string          `json:"description"`
	Photo        Photo           `json:"photo"`
	Video        Video           `json:"video"`
	PreviewPage  string          `json:"preview_page"`
	PreviewURL   string          `json:"preview_url"`
	Product      LinkProduct     `json:"product"`
	Rating       LinkRating      `json:"rating"`
	Title        string          `json:"title"`
	Target       string          `json:"target"`
	URL          string          `json:"url"`
	IsFavorite   Bool            `json:"is_favorite"`
}

type LinkProduct struct {
	Price       MarketPrice `json:"price"`
	Merchant    string      `json:"merchant"`
	OrdersCount int         `json:"orders_count"`
}

type LinkRating struct {
	ReviewsCount int     `json:"reviews_count"`
	Stars        float64 `json:"stars"`
}

type MarketPrice struct {
	Amount        string         `json:"amount"` // Amount
	Currency      MarketCurrency `json:"currency"`
	DiscountRate  int            `json:"discount_rate"`
	OldAmount     string         `json:"old_amount"`
	Text          string         `json:"text"` // Text
	OldAmountText string         `json:"old_amount_text"`
}

type MarketCurrency struct {
	ID    int    `json:"id"`    // Currency ID
	Name  string `json:"name"`  // Currency sign
	Title string `json:"title"` // Currency Title
}

type LinkApplication struct {
	AppID float64              `json:"app_id"`
	Store LinkApplicationStore `json:"store"`
}

type LinkApplicationStore struct {
	ID   float64 `json:"id"`
	Name string  `json:"name"`
}

type LinkButton struct {
	Action LinkButtonAction `json:"action"`
	Title  string           `json:"title"`
}

type LinkButtonAction struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type RepostsInfo struct {
	Count        int `json:"count"`
	WallCount    int `json:"wall_count"`
	MailCount    int `json:"mail_count"`
	UserReposted int `json:"user_reposted"`
}

type Privacy struct {
	Category string `json:"category,omitempty"`
	Lists    struct {
		Allowed  []int `json:"allowed"`
		Excluded []int `json:"excluded"`
	} `json:"lists,omitempty"`
	Owners struct {
		Allowed  []int `json:"allowed"`
		Excluded []int `json:"excluded"`
	} `json:"owners,omitempty"`
}

type ObjectWithName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Sticker struct {
	Images               []BaseImage `json:"images"`
	ImagesWithBackground []BaseImage `json:"images_with_background"`
	ProductID            int         `json:"product_id"`
	StickerID            int         `json:"sticker_id"`
	IsAllowed            Bool        `json:"is_allowed"`
	AnimationURL         string      `json:"animation_url"`
}

type CommentsInfo struct {
	Count         int  `json:"count"`
	CanPost       Bool `json:"can_post"`
	GroupsCanPost Bool `json:"groups_can_post"`
	CanClose      Bool `json:"can_close"`
	CanOpen       Bool `json:"can_open"`
}

type Geo struct {
	Coordinates string `json:"coordinates"`
	Place       Place  `json:"place"`
	Showmap     int    `json:"showmap"`
	Type        string `json:"type"`
}

type Place struct {
	Address        string         `json:"address"`
	Checkins       int            `json:"checkins"`
	City           interface{}    `json:"city"` // BUG(VK): https://github.com/VKCOM/vk-api-schema/issues/143
	Country        interface{}    `json:"country"`
	Created        int            `json:"created"`
	ID             int            `json:"id"`
	Icon           string         `json:"icon"`
	Latitude       float64        `json:"latitude"`
	Longitude      float64        `json:"longitude"`
	Title          string         `json:"title"`
	Type           string         `json:"type"`
	IsDeleted      Bool           `json:"is_deleted"`
	TotalCheckins  int            `json:"total_checkins"`
	Updated        int            `json:"updated"`
	CategoryObject CategoryObject `json:"category_object"`
}

type CategoryObject struct {
	ID    int         `json:"id"`
	Title string      `json:"title"`
	Icons []BaseImage `json:"icons"`
}

type Object struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type EventsEventAttach struct {
	Address      string `json:"address,omitempty"`       // address of event
	ButtonText   string `json:"button_text"`             // text of attach
	Friends      []int  `json:"friends"`                 // array of friends ids
	ID           int    `json:"id"`                      // event ID
	IsFavorite   Bool   `json:"is_favorite"`             // is favorite
	MemberStatus int    `json:"member_status,omitempty"` // Current user's member status
	Text         string `json:"text"`                    // text of attach
	Time         int    `json:"time,omitempty"`          // event start time
}

type Country struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
