package profile

type Profile struct {
	UID       string   `json:"uid"`
	UserName  string   `json:"username"`
	Email     string   `json:"email"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Bio       string   `json:"bio"`
	PhotoUrl  string   `json:"photo_url"`
	Category  []string `json:"category"`
	Followers []string `json:"followers"`
	Following []string `json:"following"`
	UpdatedAt int64    `json:"updated_at"`
	CreatedAt int64    `json:"created_at"`
}
