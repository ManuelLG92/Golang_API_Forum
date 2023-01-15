package handlers

type Post struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PostUpdateOptions struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Posts []Post



