package auth

type LoginResponse200 struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Token      interface{} `json:"token"`
}
