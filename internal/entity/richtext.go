package entity

type Color string

const ColorDefault Color = "default"

func (c Color) String() string {
	return string(c)
}

func (c Color) MarshalText() ([]byte, error) {
	if c == "" {
		return []byte(ColorDefault), nil
	}

	return []byte(c), nil
}

type Annotations struct {
	Bold          bool  `json:"bold"`
	Italic        bool  `json:"italic"`
	Strikethrough bool  `json:"strikethrough"`
	Underline     bool  `json:"underline"`
	Code          bool  `json:"code"`
	Color         Color `json:"color"`
}

type Link struct {
	Url string `json:"url,omitempty"`
}

type Text struct {
	Content string `json:"content"`
	Link    *Link  `json:"link,omitempty"`
}

type RichText struct {
	Type        ObjectType   `json:"type,omitempty"`
	Text        Text         `json:"text"`
	Annotations *Annotations `json:"annotations,omitempty"`
	PlainText   string       `json:"plain_text,omitempty"`
	Href        *string      `json:"href,omitempty"`
}
