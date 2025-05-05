package swit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"workflower/app/lib"
)

type SwitGateway struct {
	logger     lib.Logger
	tokenStore TokenStore
	env        lib.Env
	client     *http.Client
}

// TokenStore 구조체: 액세스 토큰과 리프레시 토큰을 저장
type TokenStore struct {
	AccessToken  string
	RefreshToken string
}

// TokenResponse 구조체: API에서 받은 토큰 응답을 저장
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (t TokenStore) String() string {
	return fmt.Sprintf("AccessToken: %s\nRefreshToken: %s", t.AccessToken, t.RefreshToken)
}

// SwitGateway 생성자
func NewSwitGateway(logger lib.Logger) *SwitGateway {
	env := lib.NewEnv()
	return &SwitGateway{
		logger:     logger,
		tokenStore: TokenStore{AccessToken: env.Swit.AccessToken, RefreshToken: env.Swit.RefreshToken},
		env:        env,
		client:     &http.Client{},
	}
}

func (g *SwitGateway) GetToken(code string) (TokenStore, error) {

	if g.tokenStore.AccessToken != "" {
		return g.tokenStore, nil
	}

	var tokenResponse TokenResponse
	var err error

	switch {
	case code != "":
		tokenResponse, err = g.requestToken(code)
	case g.tokenStore.RefreshToken != "":
		tokenResponse, err = g.refreshToken()
	default:
		return TokenStore{}, errors.New("no refresh token or code available to get new token")
	}

	if err != nil {
		return TokenStore{}, err
	}

	// 새로 얻은 토큰을 메모리에 저장
	g.tokenStore.AccessToken = tokenResponse.AccessToken
	g.tokenStore.RefreshToken = tokenResponse.RefreshToken

	g.logger.Info("Generate Token\n", g.tokenStore)
	return g.tokenStore, nil
}

// requestToken 함수: authorization code로 처음 액세스 토큰을 요청하는 함수
func (g *SwitGateway) requestToken(code string) (TokenResponse, error) {
	tokenURL := "https://openapi.swit.io/oauth/token"
	tokenData := fmt.Sprintf(
		"grant_type=authorization_code&client_id=%s&client_secret=%s&redirect_uri=%s/api/v1/oauth/callback&code=%s",
		g.env.Swit.ClientId, g.env.Swit.ClientSecret, g.env.Server.Url, code,
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
	g.logger.Info("Token expired or invalid, refreshing token...")
	tokenURL := "https://openapi.swit.io/oauth/token"
	tokenData := fmt.Sprintf(
		"grant_type=refresh_token&client_id=%s&client_secret=%s&refresh_token=%s",
		g.env.Swit.ClientId, g.env.Swit.ClientSecret, g.tokenStore.RefreshToken,
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

func (g *SwitGateway) ApiCall(httpMethod, targetUrl string, body map[string]any) ([]byte, error) {
	var req *http.Request
	var err error

	if httpMethod == "GET" {
		// GET 쿼리 파라미터 추가
		if body != nil {
			params := url.Values{}
			for key, value := range body {
				params.Add(key, fmt.Sprintf("%v", value))
			}
			targetUrl += "?" + params.Encode()
		}
		req, err = http.NewRequest("GET", targetUrl, nil)
	} else {
		// POST/PUT 등일 경우 JSON body 추가
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(httpMethod, targetUrl, bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
	}

	if err != nil {
		return nil, err
	}

	// Authorization 헤더 설정
	if accessToken := g.tokenStore.AccessToken; accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+accessToken)
	}

	// 요청 전송
	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 인증 만료 시 토큰 갱신 후 재요청
	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		g.tokenStore.AccessToken = ""
		newToken, err := g.GetToken("")
		if err != nil {
			return nil, fmt.Errorf("failed to refresh token: %w", err)
		}
		g.tokenStore.AccessToken = newToken.AccessToken
		g.tokenStore.RefreshToken = newToken.RefreshToken

		// 재시도
		return g.ApiCall(httpMethod, targetUrl, body)
	}

	return bodyResp, nil
}
