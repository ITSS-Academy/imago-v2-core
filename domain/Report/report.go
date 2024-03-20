package Report

type Report struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Reason    string `json:"reason"`
	Status    bool   `json:"status"`
	Content   string `json:"content"`
	CreatorID string `json:"creator_id"`
	UpdatedAt int64  `json:"updated_at"`
	CreatedAt int64  `json:"created_at"`
	TypeID    string `json:"type_id"`
}
