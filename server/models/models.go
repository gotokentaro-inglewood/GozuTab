package models

// 大文字で始めることで、外部（main.goなど）から参照可能になります
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IconURL string `json:"icon_url"`
}

type Tab struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Content  string `json:"content"`
	AudioURL string `json:"audio_url"`
	Status   string `json:"status"`
}

type PageData struct {
	Users []User `json:"users"`
	Tabs  []Tab  `json:"tabs"`
}