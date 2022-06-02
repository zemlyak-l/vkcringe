package object

type MessageTemplate struct {
	Type     string            `json:"type"`
	Elements []TemplateElement `json:"elements"`
}

type TemplateElement struct {
	ElementCarousel
}

type ElementCarousel struct {
	Title       string           `json:"title,omitempty"`
	Action      CarouselAction   `json:"action,omitempty"`
	Description string           `json:"description,omitempty"`
	Photo       *Photo           `json:"photo,omitempty"`
	PhotoID     string           `json:"photo_id,omitempty"`
	Buttons     []KeyboardButton `json:"buttons,omitempty"`
}

type CarouselAction struct {
	Type string `json:"type"`
	Link string `json:"link,omitempty"`
}
