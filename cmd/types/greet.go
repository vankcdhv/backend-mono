package types

type GreetRequest struct {
	UserID int64 `uri:"user_id" json:"user_id"`
}

type GreetResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
