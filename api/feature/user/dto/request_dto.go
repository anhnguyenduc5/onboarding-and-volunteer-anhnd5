package dto

type RequestCreatingDTO struct {
	UserID int `json:"user_id"`
}

type ValidateInputDTO struct {
	Gender *string `json:"gender"`
	Mobile *string `json:"mobile"`
}
