package entity

type UserID string

func (u UserID) String() string {
	return string(u)
}

type UserType string

type Person struct {
	Email string `json:"email"`
}

type Bot struct{}

type User struct {
	Object    ObjectType `json:"object"`
	ID        UserID     `json:"id"`
	Type      UserType   `json:"type"`
	Name      string     `json:"name"`
	AvatarURL string     `json:"avatar_url"`
	Person    *Person    `json:"person"`
	Bot       *Bot       `json:"bot"`
}
