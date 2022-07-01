package entity

import "time"

type ObjectType string

func (ot ObjectType) String() string {
	return string(ot)
}

type ObjectID string

func (oid ObjectID) String() string {
	return string(oid)
}

type Page struct {
	Object         ObjectType `json:"object"`
	ID             ObjectID   `json:"id"`
	CreatedTime    time.Time  `json:"created_time"`
	LastEditedTime time.Time  `json:"last_edited_time"`
	CreatedBy      User       `json:"created_by,omitempty"`
	LastEditedBy   User       `json:"last_edited_by,omitempty"`
	Archived       bool       `json:"archived"`
	Properties     Properties `json:"properties"`
	//Parent         Parent     `json:"parent"`
	URL   string `json:"url"`
	Icon  *Icon  `json:"icon,omitempty"`
	Cover *Image `json:"cover,omitempty"`
}
