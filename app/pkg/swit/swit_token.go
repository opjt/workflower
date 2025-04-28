package swit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gom/app/lib"
	"io"
	"net/http"
	"time"
)

type SwitGateway struct {
	logger     lib.Logger
	tokenStore TokenStore
	env        lib.Env
}

// TokenStore 구조체: 액세스 토큰과 리프레시 토큰을 저장
type TokenStore struct {
	AccessToken  string
	RefreshToken string
	Expiry       time.Time
}

// TokenResponse 구조체: API에서 받은 토큰 응답을 저장
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

// SwitGateway 생성자
func NewSwitGateway(logger lib.Logger) *SwitGateway {
	return &SwitGateway{
		logger:     logger,
		tokenStore: TokenStore{},
		env:        lib.NewEnv(),
	}
}

// GetToken 함수: 액세스 토큰을 얻거나 리프레시 토큰을 사용하여 새 토큰을 얻는 함수
func (g *SwitGateway) GetToken(code string) (string, error) {
	// 만약 액세스 토큰이 존재하고 만료되지 않았다면 기존 토큰을 반환
	if g.tokenStore.AccessToken != "" && time.Now().Before(g.tokenStore.Expiry) {
		return g.tokenStore.AccessToken, nil
	}

	// 액세스 토큰이 없거나 만료되었다면 새로 얻어야 함
	var tokenResponse TokenResponse
	var err error

	// 액세스 토큰이 만료되었으면 리프레시 토큰을 사용하여 새로운 토큰을 요청
	if g.tokenStore.RefreshToken != "" {
		tokenResponse, err = g.refreshToken()
	} else {
		tokenResponse, err = g.requestToken(code)
	}

	if err != nil {
		return "", err
	}

	// 새로 얻은 토큰을 메모리에 저장
	g.tokenStore.AccessToken = tokenResponse.AccessToken
	g.tokenStore.RefreshToken = tokenResponse.RefreshToken
	g.tokenStore.Expiry = time.Now().Add(time.Second * time.Duration(tokenResponse.ExpiresIn))

	return g.tokenStore.AccessToken, nil
}

// requestToken 함수: authorization code로 처음 액세스 토큰을 요청하는 함수
func (g *SwitGateway) requestToken(code string) (TokenResponse, error) {
	tokenURL := "https://openapi.swit.io/oauth/token"
	tokenData := fmt.Sprintf(
		"grant_type=authorization_code&client_id=%s&client_secret=%s&redirect_uri=%s/api/v1/oauth/callback&code=%s",
		g.env.ClientId, g.env.ClientSecret, g.env.ServerUrl, code,
	)

	resp, err := http.Post(tokenURL, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(tokenData)))
	if err != nil {
		return TokenResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TokenResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return TokenResponse{}, errors.New(string(body))
	}

	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return TokenResponse{}, err
	}

	return tokenResponse, nil
}

// refreshToken 함수: 리프레시 토큰으로 새로운 액세스 토큰을 요청하는 함수
func (g *SwitGateway) refreshToken() (TokenResponse, error) {
	tokenURL := "https://openapi.swit.io/oauth/token"
	tokenData := fmt.Sprintf(
		"grant_type=refresh_token&client_id=%s&client_secret=%s&refresh_token=%s",
		g.env.ClientId, g.env.ClientSecret, g.tokenStore.RefreshToken,
	)

	resp, err := http.Post(tokenURL, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(tokenData)))
	if err != nil {
		return TokenResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return TokenResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return TokenResponse{}, errors.New(string(body))
	}

	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return TokenResponse{}, err
	}

	return tokenResponse, nil
}

// ApiCall 함수: Swit API를 호출하는 함수
func (g *SwitGateway) ApiCall(url string, body map[string]interface{}) error {
	accessToken := g.tokenStore.AccessToken
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	if accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+accessToken)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		fmt.Println("Token expired or invalid, refreshing token...")
		newAccessToken, err := g.GetToken("some_code_here")
		if err != nil {
			return err
		}

		req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bearer "+newAccessToken)
		req.Header.Set("Content-Type", "application/json")

		resp, err = client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		bodyResp, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
	}

	fmt.Println("API Response:", string(bodyResp))
	return nil
}
