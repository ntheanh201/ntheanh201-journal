package entity

type ObjectType string

func (ot ObjectType) String() string {
	return string(ot)
}

type ObjectID string

func (oid ObjectID) String() string {
	return string(oid)
}

type Page struct {
	Object     ObjectType `json:"object"`
	ID         ObjectID   `json:"id"`
	Archived   bool       `json:"archived"`
	Properties Properties `json:"properties"`
	//URL        string     `json:"url"`
	Icon  *Icon  `json:"icon,omitempty"`
	Cover *Image `json:"cover,omitempty"`
}
