package models

// User eporting model for joke
type User struct {
	ID        int    `json:"id" binding:"required"`
	Name      string `json:"name"  binding:"required"`
	Email     string `json:"email" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}
