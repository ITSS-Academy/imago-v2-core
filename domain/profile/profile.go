package profile

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	ID        string   `json:"id" gorm:"primaryKey"`
	UserName  string   `json:"username"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Bio       string   `json:"bio"`
	Email     string   `json:"email"`
	PhotoURL  string   `json:"photo_url"`
	Category  []string `json:"category"`
	Followers []string `json:"followers"`
	Following []string `json:"following"`
}

type ProfileRepository interface {
}
