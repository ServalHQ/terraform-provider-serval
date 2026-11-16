package client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const tokenExpiryMargin = 10 * time.Second

// tokenSource caches a client-credentials access token for one Client.
type tokenSource struct {
	clientID     string
	clientSecret string

	mu        sync.Mutex
	value     string
	expiresAt time.Time
}

func (s *tokenSource) token(ctx context.Context, httpClient *http.Client, baseURL *url.URL) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.value != "" && time.Now().Before(s.expiresAt.Add(-tokenExpiryMargin)) {
		return s.value, nil
	}
	s.value, s.expiresAt = "", time.Time{}

	tokenURL, err := baseURL.Parse(tokenPath)
	if err != nil {
		return "", fmt.Errorf("client: invalid OAuth2 token url: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL.String(), strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		return "", fmt.Errorf("client: build OAuth2 token request: %w", err)
	}
	creds := base64.StdEncoding.EncodeToString([]byte(s.clientID + ":" + s.clientSecret))
	req.Header.Set("Authorization", "Basic "+creds)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("client: OAuth2 token request failed: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("client: OAuth2 token request returned status %d, expected %d", res.StatusCode, http.StatusOK)
	}
	contents, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("client: read OAuth2 token response: %w", err)
	}
	var body struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err := json.Unmarshal(contents, &body); err != nil {
		return "", fmt.Errorf("client: parse OAuth2 token response: %w", err)
	}
	if body.AccessToken == "" {
		return "", fmt.Errorf("client: OAuth2 token response missing access_token")
	}
	s.value = body.AccessToken
	s.expiresAt = time.Now().Add(time.Duration(body.ExpiresIn) * time.Second)
	return s.value, nil
}
