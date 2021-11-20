package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	if expirationTime != 24 {
		t.Error("expiration time should be 24 hours")
	}
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	token := GetNewAccessToken()
	assert.False(t, token.IsExpired(), "brand new token should not be expired")
	assert.EqualValues(t, "", token.AccessToken, "new access token should not have defined access token id")
	assert.Truef(t, token.UserId == 0, "new access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	token := AccessToken{}
	assert.True(t, token.IsExpired(), "empty access token should be expired")
	token.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, token.IsExpired(), "access token expiring three hours from now should NOT be expired")
}