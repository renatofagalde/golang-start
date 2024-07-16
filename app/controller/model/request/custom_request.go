package request

type CustomRequest struct {
	Custom string `json:"custom" binding:"required,custom" example:"lorem ipsum"`
	FullName string `json:"full_name" binding:"required,full_name"`
  Email string `json:email binding:"required"`
}
