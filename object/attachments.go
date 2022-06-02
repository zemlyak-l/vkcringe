package object

type Attachments struct {
	Photo struct {
		AlbumID int  `json:"album_id"`
		Date    int  `json:"date"`
		HasTags bool `json:"has_tags"`
		ID      int  `json:"id"`
		OwnerID int  `json:"owner_id"`
		Sizes   []struct {
			Height int    `json:"height"`
			Type   string `json:"type"`
			Url    string `json:"url"`
		} `json:"sizes"`
		Text string `json:"text"`
	} `json:"photo"`
	Type string `json:"type"`
}
