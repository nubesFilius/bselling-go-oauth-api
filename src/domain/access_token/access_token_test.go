package access_token

import (
	"testing"
	"time"
)

func TestAcessTokenConstants(t *testing.T) {
	if expirationTime != 24 {
		t.Error("expiration time should be 24 hours")
	}
}

func TestGetNewAccessToken(t *testing.T) {
	token := GetNewAccessToken()
	if token.IsExpired() {
		t.Error("new token should not be expired")
	}

	if token.AccessToken != "" {
		t.Error("new access token should not have defined access token id")
	}

	if token.UserId != 0 {
		t.Error("new access token should not have an associtokened user id")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	token := AccessToken{}
	if !token.IsExpired() {
		t.Error("empty access token should be expired by default")
	}

	token.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if token.IsExpired() {
		t.Error("access token expiring three hours from now should NOT be expired")
	}
}