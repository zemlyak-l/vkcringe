package object

type MarketCommentNewObject WallWallComment

type MarketCommentEditObject WallWallComment

type MarketCommentRestoreObject WallWallComment

type MarketOrderNewObject MarketOrder

type MarketOrderEditObject MarketOrder

type MarketOrder struct {
	ID                int                  `json:"id"`
	GroupID           int                  `json:"group_id"`
	UserID            int                  `json:"user_id"`
	Date              int                  `json:"date"`
	Status            int                  `json:"status"`
	ItemsCount        int                  `json:"items_count"`
	TotalPrice        MarketPrice          `json:"total_price"`
	DisplayOrderID    string               `json:"display_order_id"`
	Comment           string               `json:"comment"`
	PreviewOrderItems []MarketOrderItem    `json:"preview_order_items"`
	PriceDetails      []MarketPriceDetail  `json:"price_details"`
	Delivery          MarketOrderDelivery  `json:"delivery"`
	Recipient         MarketOrderRecipient `json:"recipient"`
}

type MarketOrderDelivery struct {
	TrackNumber   string              `json:"track_number"`
	TrackLink     string              `json:"track_link"`
	Address       string              `json:"address"`
	Type          string              `json:"type"`
	DeliveryPoint MarketDeliveryPoint `json:"delivery_point,omitempty"`
}

type MarketDeliveryPoint struct {
	ID           int                        `json:"id"`
	ExternalID   string                     `json:"external_id"`
	OutpostOnly  Bool                       `json:"outpost_only"`
	CashOnly     Bool                       `json:"cash_only"`
	Address      MarketDeliveryPointAddress `json:"address"`
	DisplayTitle string                     `json:"display_title"`
	ServiceID    int                        `json:"service_id"`
}

type MarketDeliveryPointAddress struct {
	ID             int     `json:"id"`
	Address        string  `json:"address"`
	CityID         int     `json:"city_id"`
	CountryID      int     `json:"country_id"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Phone          string  `json:"phone"`
	Title          string  `json:"title"`
	WorkInfoStatus string  `json:"work_info_status"`
}

type MarketOrderRecipient struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	DisplayText string `json:"display_text"`
}

type MarketMarketItem struct {
	AccessKey    string               `json:"access_key"`   // Access key for the market item
	Availability int                  `json:"availability"` // Information whether the item is available
	Category     MarketMarketCategory `json:"category"`

	// Date when the item has been created in Unixtime.
	Date               int                        `json:"date,omitempty"`
	Description        string                     `json:"description"` // Item description
	ID                 int                        `json:"id"`          // Item ID
	OwnerID            int                        `json:"owner_id"`    // Item owner's ID
	Price              MarketPrice                `json:"price"`
	ThumbPhoto         string                     `json:"thumb_photo"` // URL of the preview image
	Title              string                     `json:"title"`       // Item title
	CanComment         Bool                       `json:"can_comment"`
	CanRepost          Bool                       `json:"can_repost"`
	IsFavorite         Bool                       `json:"is_favorite"`
	IsMainVariant      Bool                       `json:"is_main_variant"`
	AlbumsIDs          []int                      `json:"albums_ids"`
	Photos             []Photo                    `json:"photos"`
	Likes              LikesInfo                  `json:"likes"`
	Reposts            RepostsInfo                `json:"reposts"`
	ViewsCount         int                        `json:"views_count,omitempty"`
	URL                string                     `json:"url"` // URL to item
	ButtonTitle        string                     `json:"button_title"`
	ExternalID         string                     `json:"external_id"`
	Dimensions         MarketDimensions           `json:"dimensions"`
	Weight             int                        `json:"weight"`
	VariantsGroupingID int                        `json:"variants_grouping_id"`
	PropertyValues     []MarketMarketItemProperty `json:"property_values"`
	CartQuantity       int                        `json:"cart_quantity"`
}

type MarketMarketCategory struct {
	ID      int           `json:"id"`   // Category ID
	Name    string        `json:"name"` // Category name
	Section MarketSection `json:"section"`
}

type MarketSection struct {
	ID   int    `json:"id"`   // Section ID
	Name string `json:"name"` // Section name
}

type MarketDimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Length int `json:"length"`
}

type MarketMarketItemProperty struct {
	VariantID    int    `json:"variant_id"`
	VariantName  string `json:"variant_name"`
	PropertyName string `json:"property_name"`
}

type MarketOrderItem struct {
	OwnerID  int              `json:"owner_id"`
	ItemID   int              `json:"item_id"`
	Price    MarketPrice      `json:"price"`
	Quantity int              `json:"quantity"`
	Item     MarketMarketItem `json:"item"`
	Title    string           `json:"title"`
	Photo    Photo            `json:"photo"`
	Variants []string         `json:"variants"`
}

type MarketPriceDetail struct {
	Title    string      `json:"title"`
	Price    MarketPrice `json:"price"`
	IsAccent Bool        `json:"is_accent,omitempty"`
}
