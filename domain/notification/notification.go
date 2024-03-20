package notification

type Notification struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	CreatedAt  int64  `json:"created_at"`
	CreatorID  string `json:"creator_id"`
	ReceiverID string `json:"receiver_id"`
}
