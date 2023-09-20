package types

type Post struct {
	UserId   int       `json:"userId"`
	PostID   int       `json:"id,omitempty"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	PostId    int    `json:"postId,omitempty"`
	CommentID int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Body      string `json:"body"`
}
