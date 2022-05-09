package types

type GetUserByIDRequest struct {
	UserID int64 `uri:"user_id"`
}

type GetUserByIDResponse struct {
	Code  int       `json:"code"`
	Data  *UserData `json:"data"`
	Error string    `json:"error"`
}

type UserData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
