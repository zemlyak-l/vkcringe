package object

type Doc struct {
	AccessKey  string      `json:"access_key"` // Access key for the document
	Date       int         `json:"date"`       // Date when file has been uploaded in Unixtime
	Ext        string      `json:"ext"`        // File extension
	ID         int         `json:"id"`         // Document ID
	IsLicensed Bool        `json:"is_licensed"`
	OwnerID    int         `json:"owner_id"` // Document owner ID
	Preview    DocsPreview `json:"preview"`
	Size       int         `json:"size"`  // File size in bites
	Title      string      `json:"title"` // Document title
	Type       int         `json:"type"`  // Document type
	URL        string      `json:"url"`   // File URL
	DocsPreviewAudioMessage
	DocsPreviewGraffiti
}

type DocsPreview struct {
	Photo        DocsPreviewPhoto        `json:"photo"`
	Graffiti     DocsPreviewGraffiti     `json:"graffiti"`
	Video        DocsPreviewVideo        `json:"video"`
	AudioMessage DocsPreviewAudioMessage `json:"audio_message"`
}

type DocsPreviewAudioMessage struct {
	Duration        int    `json:"duration"`
	Waveform        []int  `json:"waveform"`
	LinkOgg         string `json:"link_ogg"`
	LinkMp3         string `json:"link_mp3"`
	Transcript      string `json:"transcript"`
	TranscriptState string `json:"transcript_state"`
}

type DocsPreviewGraffiti struct {
	Src    string `json:"src"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type DocsPreviewVideo struct {
	FileSize int    `json:"file_size"` // Video file size in bites
	Height   int    `json:"height"`    // Video's height in pixels
	Src      string `json:"src"`       // Video URL
	Width    int    `json:"width"`     // Video's width in pixels
}

type DocsPreviewPhoto struct {
	Sizes []DocsPreviewPhotoSizes `json:"sizes"`
}

type DocsPreviewPhotoSizes struct {
	// BUG(VK): json: cannot unmarshal number 162.000000 into Go struct field
	// DocsPreviewPhotoSizes.doc.preview.photo.sizes.height of type Int
	Height float64 `json:"height"` // Height in px
	Src    string  `json:"src"`    // URL of the image
	Type   string  `json:"type"`
	Width  float64 `json:"width"` // Width in px
}
