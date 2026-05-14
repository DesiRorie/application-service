package dto

type CreateApplicationRequest struct {
	Company  string `json:"company" binding:"required"`
	Position string `json:"position" binding:"required"`
	Status   string `json:"status" binding:"required"`
	Notes    string `json:"notes"`
}
