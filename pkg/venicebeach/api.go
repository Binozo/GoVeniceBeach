package venicebeach

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"resty.dev/v3"
)

const authBaseUrl = "https://oauth.actinate.com"
const apiBaseUrl = "https://venicebeach.api.actinate.com/gym/v1"

func Login(email, password, clientSecret string) (*Session, error) {
	client := resty.New()
	client.SetBaseURL(authBaseUrl)
	defer client.Close()

	res, err := client.
		R().
		SetQueryParam("lang", "en").
		SetFormData(map[string]string{
			"grant_type":    "password",
			"username":      email,
			"password":      password,
			"client_id":     "venicebeachV2",
			"client_secret": clientSecret,
			"app_version":   "7.0.0",
			"os_version":    "36",
			"os":            "android",
		}).
		Post("/v6/oauth/token")

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}
	if err = json.Unmarshal(body, &jsonData); err != nil {
		return nil, err
	}

	if jsonData["error"] != nil {
		errorString := jsonData["error"].(string)
		switch errorString {
		case "user_not_found":
			return nil, ErrUserNotFound
		case "wrong_password":
			return nil, ErrInvalidPassword
		default:
			return nil, errors.New(errorString)
		}
	}

	accessToken := jsonData["access_token"].(string)
	_ = jsonData["refresh_token"].(string) // Unused

	return NewSession(accessToken), nil
}

func (s *Session) Ping() error {
	_, err := s.client.
		R().
		Head(fmt.Sprintf("%s/reachability", authBaseUrl))

	if err != nil {
		return err
	}

	return nil
}

func (s *Session) UserInfo() (UserInfo, error) {
	res, err := s.client.
		R().
		Get("/users/me?lang=en")

	if err != nil {
		return UserInfo{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return UserInfo{}, err
	}

	var userInfo UserInfo
	if err = json.Unmarshal(body, &userInfo); err != nil {
		return UserInfo{}, err
	}

	return userInfo, nil
}

func (s *Session) Timeline() (Timeline, error) {
	res, err := s.client.
		R().
		SetQueryParam("limit", "6").
		SetQueryParam("numberOfChallenges", "4").
		SetQueryParam("numberOfLibraryVideos", "6").
		SetQueryParam("categoryId", "all").
		SetQueryParam("scheduledPosts", "0").
		SetQueryParam("numberOfAppointmentBookings", "0").
		SetQueryParam("lang", "en").
		Get("/timeline")

	if err != nil {
		return Timeline{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Timeline{}, err
	}

	var timeline Timeline
	if err = json.Unmarshal(body, &timeline); err != nil {
		return Timeline{}, err
	}

	return timeline, nil
}

func (s *Session) GetStudioInfo(studioId int) (Studio, error) {
	res, err := s.client.
		R().
		Get(fmt.Sprintf("/studios/%d", studioId))

	if err != nil {
		return Studio{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Studio{}, err
	}

	var rawStudioInfo struct {
		Studio Studio `json:"studio"`
	}
	if err = json.Unmarshal(body, &rawStudioInfo); err != nil {
		return Studio{}, err
	}

	return rawStudioInfo.Studio, nil
}
