package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nubesFilius/bselling-go-oauth-api.git/src/client/cassandra"
	"github.com/nubesFilius/bselling-go-oauth-api.git/src/domain/access_token"
	"github.com/nubesFilius/bselling-go-oauth-api.git/src/http"
	"github.com/nubesFilius/bselling-go-oauth-api.git/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session , err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	session.Close()

	tokenHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", tokenHandler.GetById)
	router.POST("/oauth/access_token", tokenHandler.Create)

	router.Run(":8080")
}