package auth

type GoogleLoginCallbackResponse200 struct {
	StatusCode int         `json:"status_code"`
	UserInfo   interface{} `json:"user_info"`
	Token      interface{} `json:"token"`
}
