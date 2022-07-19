package entity

type TableOfContents interface {
}

type Bookmark struct {
	URL     string      `json:"url,omitempty"`
	Caption *[]RichText `json:"caption,omitempty"`
}

type Heading struct {
	Title *[]RichText `json:"rich_text,omitempty"`
	Color *Color      `json:"color,omitempty"`
}

type TextBlock struct {
	Title *[]RichText `json:"rich_text,omitempty"`
	Color *Color      `json:"color,omitempty"`
}

type Block struct {
	HasChildren     bool             `json:"has_children"`
	Archived        bool             `json:"archived"`
	Type            string           `json:"type"`
	Paragraph       *Paragraph       `json:"paragraph,omitempty"`
	Code            *Code            `json:"code,omitempty"`
	Bookmark        *Bookmark        `json:"bookmark,omitempty"`
	Heading1        *Heading         `json:"heading_1,omitempty"`
	Heading2        *Heading         `json:"heading_2,omitempty"`
	Heading3        *Heading         `json:"heading_3,omitempty"`
	TableOfContents *TableOfContents `json:"table_of_contents,omitempty"`
	Quote           *TextBlock       `json:"quote,omitempty"`
}
