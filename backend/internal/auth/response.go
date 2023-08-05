package auth

type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
	Token         string `json:"token"`
}

type GoogleLoginCallbackResponse200 struct {
	StatusCode int  `json:"status_code"`
	User       User `json:"user"`
}

type UserEmail struct {
	Email string `json:"email"`
}
