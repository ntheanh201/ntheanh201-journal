package entity

type Paragraph struct {
	RichText *[]RichText `json:"rich_text"`
	Color    Color       `json:"color"`
}
