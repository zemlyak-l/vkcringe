package object

type Video struct {
	AccessKey                string             `json:"access_key"`
	AddingDate               int                `json:"adding_date"`
	ReleaseDate              int                `json:"release_date"`
	CanAdd                   Bool               `json:"can_add"`
	CanAddToFaves            Bool               `json:"can_add_to_faves"`
	CanComment               Bool               `json:"can_comment"`
	CanEdit                  Bool               `json:"can_edit"`
	CanLike                  Bool               `json:"can_like"`
	CanDownload              Bool               `json:"can_download"`
	CanRepost                Bool               `json:"can_repost"`
	CanSubscribe             Bool               `json:"can_subscribe"`
	CanAttachLink            Bool               `json:"can_attach_link"`
	IsFavorite               Bool               `json:"is_favorite"`
	IsPrivate                Bool               `json:"is_private"`
	IsExplicit               Bool               `json:"is_explicit"`
	IsSubscribed             Bool               `json:"is_subscribed"`
	Added                    Bool               `json:"added"`
	Repeat                   Bool               `json:"repeat"` // Information whether the video is repeated
	ContentRestricted        int                `json:"content_restricted"`
	Live                     Bool               `json:"live"` // Returns if the video is a live stream
	Upcoming                 Bool               `json:"upcoming"`
	Comments                 int                `json:"comments"`    // Number of comments
	Date                     int                `json:"date"`        // Date when video has been uploaded in Unixtime
	Description              string             `json:"description"` // Video description
	Duration                 int                `json:"duration"`    // Video duration in seconds
	Files                    VideoFiles         `json:"files"`
	Trailer                  VideoFiles         `json:"trailer,omitempty"`
	FirstFrame               []VideoImage       `json:"first_frame"`
	Image                    []VideoImage       `json:"image"`
	Height                   int                `json:"height"`   // Video height
	ID                       int                `json:"id"`       // Video ID
	OwnerID                  int                `json:"owner_id"` // Video owner ID
	UserID                   int                `json:"user_id"`
	Photo130                 string             `json:"photo_130"`  // URL of the preview image with 130 px in width
	Photo320                 string             `json:"photo_320"`  // URL of the preview image with 320 px in width
	Photo640                 string             `json:"photo_640"`  // URL of the preview image with 640 px in width
	Photo800                 string             `json:"photo_800"`  // URL of the preview image with 800 px in width
	Photo1280                string             `json:"photo_1280"` // URL of the preview image with 1280 px in width
	Player                   string             `json:"player"`
	Processing               int                `json:"processing"` // Returns if the video is processing
	Title                    string             `json:"title"`      // Video title
	Subtitle                 string             `json:"subtitle"`   // Video subtitle
	Type                     string             `json:"type"`
	Views                    int                `json:"views"` // Number of views
	Width                    int                `json:"width"` // Video width
	Platform                 string             `json:"platform"`
	LocalViews               int                `json:"local_views"`
	Likes                    LikesInfo          `json:"likes"`   // Count of likes
	Reposts                  RepostsInfo        `json:"reposts"` // Count of views
	TrackCode                string             `json:"track_code"`
	PrivacyView              Privacy            `json:"privacy_view"`
	PrivacyComment           Privacy            `json:"privacy_comment"`
	ActionButton             VideoActionButton  `json:"action_button"`
	Restriction              VideoRestriction   `json:"restriction"`
	ContentRestrictedMessage string             `json:"content_restricted_message"`
	MainArtists              []AudioAudioArtist `json:"main_artists"`
	FeaturedArtists          []AudioAudioArtist `json:"featured_artists"`
	Genres                   []ObjectWithName   `json:"genres"`
	OvID                     string             `json:"ov_id,omitempty"`
}

type VideoImage struct {
	BaseImage
	WithPadding Bool `json:"with_padding"`
}

type VideoFiles struct {
	External     string `json:"external,omitempty"` // URL of the external player
	Mp4_1080     string `json:"mp4_1080,omitempty"` // URL of the mpeg4 file with 1080p quality
	Mp4_1440     string `json:"mp4_1440,omitempty"` // URL of the mpeg4 file with 2k quality
	Mp4_2160     string `json:"mp4_2160,omitempty"` // URL of the mpeg4 file with 4k quality
	Mp4_240      string `json:"mp4_240,omitempty"`  // URL of the mpeg4 file with 240p quality
	Mp4_360      string `json:"mp4_360,omitempty"`  // URL of the mpeg4 file with 360p quality
	Mp4_480      string `json:"mp4_480,omitempty"`  // URL of the mpeg4 file with 480p quality
	Mp4_720      string `json:"mp4_720,omitempty"`  // URL of the mpeg4 file with 720p quality
	Live         string `json:"live,omitempty"`
	HLS          string `json:"hls,omitempty"`
	DashUni      string `json:"dash_uni,omitempty"`
	DashSep      string `json:"dash_sep,omitempty"`
	DashWebm     string `json:"dash_webm,omitempty"`
	FailoverHost string `json:"failover_host,omitempty"`
}

type VideoActionButton struct {
	ID      string       `json:"id"`
	Type    string       `json:"type"`
	URL     string       `json:"url"`
	Snippet VideoSnippet `json:"snippet"`
}

type VideoSnippet struct {
	Description string      `json:"description"`
	OpenTitle   string      `json:"open_title"`
	Title       string      `json:"title"`
	TypeName    string      `json:"type_name"`
	Date        int         `json:"date"`
	Image       []BaseImage `json:"image"`
}

type VideoRestriction struct {
	Title          string      `json:"title"`
	Text           string      `json:"text"`
	AlwaysShown    Bool        `json:"always_shown"`
	Blur           Bool        `json:"blur"`
	CanPlay        Bool        `json:"can_play"`
	CanPreview     Bool        `json:"can_preview"`
	CardIcon       []BaseImage `json:"card_icon"`
	ListIcon       []BaseImage `json:"list_icon"`
	DisclaimerType int         `json:"disclaimer_type"`
}

type VideoNewObject Video

type VideoCommentNewObject WallWallComment

type VideoCommentEditObject WallWallComment

type VideoCommentRestoreObject WallWallComment
