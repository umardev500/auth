package domain

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=6"`
	Password string `json:"password" validate:"required,min=6"`
}

type Options struct {
	Token string `json:"token,omitempty"`
}

type LoginResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message,omitempty"`
	Token      string `json:"token,omitempty"`
}
