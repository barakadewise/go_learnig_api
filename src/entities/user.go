package entities

// User struct represents a user in the system
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name" binding:"required"`
	DOB       string `json:"dob" binding:"required"`
	Age       int    `json:"age" binding:"required"`
	Address   string `json:"address" binding:"required"`
	CompanyID int    `json:"company_id" default:"1"`
}
