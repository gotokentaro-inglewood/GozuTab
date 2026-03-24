package models

// 大文字で始めることで、外部（main.goなど）から参照可能になります
type User struct {
	ID      int
	Name    string
	Email   string
	IconURL string
}

type Tab struct {
	ID       int
	UserID   int
	Title    string
	Artist   string
	Content  string
	AudioURL string
	Status   string
}

type PageData struct {
	Users []User
	Tabs  []Tab
}