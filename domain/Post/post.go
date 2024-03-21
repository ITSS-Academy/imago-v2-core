package Post

type Post struct {
	ID         string   `json:"id"`
	Content    string   `json:"content"`
	CreatorID  string   `json:"creator_id"`
	PhotoUrl   string   `json:"photo_url"`
	UpdatedAt  int64    `json:"updated_at"`
	CreatedAt  int64    `json:"created_at"`
	CategoryID []string `json:"category_id"`
	Comment    []string `json:"comment"`
	Like       []string `json:"like"`
	HashTag    []string `json:"hash_tag"`
	Share      []string `json:"share"`
	Mention    []string `json:"mention"`
	Status     bool     `json:"status"`
}
