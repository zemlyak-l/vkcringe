package object

type Audio struct {
	AccessKey           string             `json:"access_key"`
	ID                  int                `json:"id"`
	OwnerID             int                `json:"owner_id"`
	Artist              string             `json:"artist"`
	Title               string             `json:"title"`
	Duration            int                `json:"duration"`
	Date                int                `json:"date"`
	URL                 string             `json:"url"`
	IsHq                Bool               `json:"is_hq"`
	IsExplicit          Bool               `json:"is_explicit"`
	StoriesAllowed      Bool               `json:"stories_allowed"`
	ShortVideosAllowed  Bool               `json:"short_videos_allowed"`
	IsFocusTrack        Bool               `json:"is_focus_track"`
	IsLicensed          Bool               `json:"is_licensed"`
	StoriesCoverAllowed Bool               `json:"stories_cover_allowed"`
	LyricsID            int                `json:"lyrics_id"`
	AlbumID             int                `json:"album_id"`
	GenreID             int                `json:"genre_id"`
	TrackCode           string             `json:"track_code"`
	NoSearch            int                `json:"no_search"`
	MainArtists         []AudioAudioArtist `json:"main_artists"`
	Ads                 AudioAds           `json:"ads"`
	Subtitle            string             `json:"subtitle"`
}

type AudioAudioArtist struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Domain string `json:"domain"`
}

type AudioAds struct {
	ContentID      string `json:"content_id"`
	Duration       string `json:"duration"`
	AccountAgeType string `json:"account_age_type"`
	PUID1          string `json:"puid1"`
	PUID22         string `json:"puid22"`
}

type AudioNewObject Audio
