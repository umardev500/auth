package domain

type BasicResponse struct {
	StatusCode int64  `json:"status_code"`
	Message    string `json:"message"`
}
