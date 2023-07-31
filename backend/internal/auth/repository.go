package auth

import (
	"fmt"

	"github.com/ak9024/okr-generator/config"
	"github.com/go-resty/resty/v2"
)

func (a *auth) InsertUser(m UserModel) (*resty.Response, error) {
	endpoint := fmt.Sprintf("%s/users", a.Config.GetString("supabase.url"))

	resp, err := DoRequest(a.Config).
		EnableTrace().
		SetBody(m).
		Post(endpoint)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *auth) ViewUserFilterByEmail(email string) (*UserEmail, error) {
	endpoint := fmt.Sprintf("%s/users?email=eq.%s&select=email", a.Config.GetString("supabase.url"), email)

	var results []UserEmail

	_, err := DoRequest(a.Config).
		EnableTrace().
		SetResult(&results).
		Get(endpoint)
	if err != nil {
		return nil, err
	}

	if len(results) > 0 {
		return &results[0], nil
	}

	return nil, nil
}

func DoRequest(cfg config.Provider) *resty.Request {
	client := resty.New()
	return client.R().
		SetHeaders(map[string]string{
			"Content-Type":  "application/json",
			"apikey":        cfg.GetString("supabase.key"),
			"Authorization": fmt.Sprintf("Bearer %s", cfg.GetString("supabase.key")),
		})
}
