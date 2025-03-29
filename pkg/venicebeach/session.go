package venicebeach

import (
	"resty.dev/v3"
	"runtime"
)

type Session struct {
	authToken string
	client    *resty.Client
}

func NewSession(token string) *Session {
	client := resty.New()

	client.SetAuthScheme("Bearer")
	client.SetAuthToken(token)
	client.SetBaseURL(apiBaseUrl)

	session := &Session{
		authToken: token,
		client:    client,
	}

	runtime.AddCleanup(session, func(client *resty.Client) {
		client.Close()
	}, session.client)

	return session
}

func (s *Session) GetToken() string {
	return s.authToken
}
