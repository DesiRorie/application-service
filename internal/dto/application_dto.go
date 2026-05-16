package dto

type CreateApplicationRequest struct {
	Company  string `json:"company" binding:"required"`
	Position string `json:"position" binding:"required"`
	Status   string `json:"status" binding:"required"`
	Notes    string `json:"notes"`
}
type UpdateApplication struct {
	Company  string `json:"company"`
	Position string `json:"position"`
	Status   string `json:"status" `
	Notes    string `json:"notes"`
}
