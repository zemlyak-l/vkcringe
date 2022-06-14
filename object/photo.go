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
	CanUpload          Bool         `json:"can_upload"`
	CommentsDisabled   Bool         `json:"comments_disabled"`
	ThumbIsLast        Bool         `json:"thumb_is_last"`
	UploadByAdminsOnly Bool         `json:"upload_by_admins_only"`
	HasTags            Bool         `json:"has_tags"`
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

type PhotoAlbum struct {
	Created     int    `json:"created"`     // Date when the album has been created in Unixtime
	Description string `json:"description"` // Photo album description
	ID          int    `json:"id"`          // Photo album ID
	OwnerID     int    `json:"owner_id"`    // Album owner's ID
	Size        int    `json:"size"`        // Photos number
	Thumb       Photo  `json:"thumb"`
	Title       string `json:"title"`   // Photo album title
	Updated     int    `json:"updated"` // Date when the album has been updated last time in Unixtime
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

type PhotoNewObject Photo
type PhotoCommentNewObject WallWallComment
type PhotoCommentEditObject WallWallComment
type PhotoCommentRestoreObject WallWallComment
