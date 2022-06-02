package object

type Photo struct {
	AccessKey          string       `json:"access_key"`
	AlbumID            int          `json:"album_id"`
	Date               int          `json:"date"`
	Height             int          `json:"height"`
	ID                 int          `json:"id"`
	Images             []Image      `json:"images"`
	Lat                float64      `json:"lat"`
	Long               float64      `json:"long"`
	OwnerID            int          `json:"owner_id"`
	PostID             int          `json:"post_id"`
	Text               string       `json:"text"`
	UserID             int          `json:"user_id"`
	Width              int          `json:"width"`
	CanUpload          bool         `json:"can_upload"`
	CommentsDisabled   bool         `json:"comments_disabled"`
	ThumbIsLast        bool         `json:"thumb_is_last"`
	UploadByAdminsOnly bool         `json:"upload_by_admins_only"`
	HasTags            bool         `json:"has_tags"`
	Created            int          `json:"created"`
	Description        string       `json:"description"`
	PrivacyComment     []string     `json:"privacy_comment"`
	PrivacyView        []string     `json:"privacy_view"`
	Size               int          `json:"size"`
	Sizes              []PhotoSizes `json:"sizes"`
	ThumbID            int          `json:"thumb_id"`
	ThumbSrc           string       `json:"thumb_src"`
	Title              string       `json:"title"`
	Updated            int          `json:"updated"`
	Color              string       `json:"color"`
}

type PhotoSizes struct {
	BaseImage
}

type Image struct {
	BaseImage
	Type string `json:"type"`
}

type BaseImage struct {
	Height float64 `json:"height"`
	URL    string  `json:"url"`
	Width  float64 `json:"width"`
	Type   string  `json:"type"`
}
