package access_token

import (
	"github.com/nubesFilius/bselling-go-oauth-api.git/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId int64 `json:"user_id"`
	ClientId int64 `json:"client_id"`
	Expires int64 `json:"expires"`
}

func (token *AccessToken) Validate() *errors.RestErr {
	token.AccessToken = strings.TrimSpace(token.AccessToken)
	if token.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if token.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if token.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if token.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (token AccessToken) IsExpired() bool {
	return time.Unix(token.Expires, 0).Before(time.Now().UTC())
}

// Web FrontEnd - Client-Id: 123
// Android App - Client-Id:234

