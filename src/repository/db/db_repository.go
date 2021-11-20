package db

import (
	"github.com/nubesFilius/bselling-go-oauth-api.git/src/client/cassandra"
	"github.com/nubesFilius/bselling-go-oauth-api.git/src/domain/access_token"
	"github.com/nubesFilius/bselling-go-oauth-api.git/src/utils/errors"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires from access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?, ?, ?, ?);"
	queryUpdateExpires = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(token access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {

}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session , err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	var token access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(&token.AccessToken, &token.UserId, &token.ClientId, &token.Expires); err != nil {
		if err.Error() == "not found" {
			return nil, errors.NewNotFoundError("no access token found wht given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &token, nil
}

func (r *dbRepository) Create(token access_token.AccessToken) *errors.RestErr {
	session , err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryCreateAccessToken,
		&token.AccessToken,
		&token.UserId,
		&token.ClientId,
		&token.Expires).Exec();
		err != nil {
			return errors.NewInternalServerError(err.Error())
		}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Closed()
	if err := session.Query(queryUpdateExpires, token.Expires, token.AccessToken).Exec();
		err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}