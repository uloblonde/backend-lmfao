package authdto

type AuthResponse struct {
	Email string `json:"email" form:"email" validate:"required" `
	Token string `json:"token" gorm:"type"`
}
type LoginResponse struct {
	Email string `json:"email" gorm:"type"`
	Token string `json:"token" gorm:"type"`
	Role  string `gorm:"type: varchar(255)" json:"role"`
}
