package entity

type RichTextBlock struct {
	Text     []RichText `json:"rich_text"`
	Children []Block    `json:"children,omitempty"`
}

type Code struct {
	RichTextBlock
	Caption  *[]RichText `json:"caption,omitempty"`
	Language *string     `json:"language,omitempty"`
}
