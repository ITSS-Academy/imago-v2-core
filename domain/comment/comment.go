package comment

type Comment struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	CreatorID string `json:"creator_id"`
	PostID    string `json:"post_id"`
	UpdatedAt int64  `json:"updated_at"`
	CreatedAt int64  `json:"created_at"`
}
