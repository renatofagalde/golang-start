package response

type CustomResponse struct {
	ID       string `json:id`
	Custom   string `json:"custom"`
	FullName string `json:"full_name"`
	Email    string `json: email`
}
