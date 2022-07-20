package entity

import "encoding/json"

type Record struct {
	// fields returned by the server
	Role string `json:"role"`
	// polymorphic value of the record, which we decode into Block, Space etc.
	Value json.RawMessage `json:"value"`

	// fields calculated from Value based on type
	ID    string `json:"-"`
	Table string `json:"-"`
	//Activity       *Activity       `json:"-"`
	//Block          *Block          `json:"-"`
	//Space          *Space          `json:"-"`
	//NotionUser     *NotionUser     `json:"-"`
	//UserRoot       *UserRoot       `json:"-"`
	//UserSettings   *UserSettings   `json:"-"`
	//Collection     *Collection     `json:"-"`
	//CollectionView *CollectionView `json:"-"`
	//Comment        *Comment        `json:"-"`
	//Discussion     *Discussion     `json:"-"`
	// TODO: add more types
}

type RecordMap struct {
	Version         int                `json:"__version__"`
	Activities      map[string]*Record `json:"activity"`
	Blocks          map[string]*Record `json:"block"`
	Spaces          map[string]*Record `json:"space"`
	NotionUsers     map[string]*Record `json:"notion_user"`
	UsersRoot       map[string]*Record `json:"user_root"`
	UserSettings    map[string]*Record `json:"user_setting"`
	Collections     map[string]*Record `json:"collection"`
	CollectionViews map[string]*Record `json:"collection_view"`
	Comments        map[string]*Record `json:"comment"`
	Discussions     map[string]*Record `json:"discussion"`
}

// TODO: reference: https://github.dev/kjk/notionapi/blob/master/api_loadCachedPageChunk.go
