package entity

type Date struct {
	Start    *string `json:"start"`
	End      *string `json:"end,omitempty"`
	TimeZone *string `json:"time_zone,omitempty"`
}
