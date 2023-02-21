package domain

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=6"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponseData struct {
	UserId string `json:"user_id,omitempty"`
	User   string `json:"user,omitempty"`
	Name   string `json:"name,omitempty"`
}

type LoginResponse struct {
	IsEmpty bool              `json:"is_empty"`
	Payload LoginResponseData `json:"payload"`
}

type Options struct {
	Token string `json:"token,omitempty"`
}

type LoginResponsePayload struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message,omitempty"`
	Token      string `json:"token,omitempty"`
}
