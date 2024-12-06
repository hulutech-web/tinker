package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

func Auth(ctx http.Context) models.User {
	var user models.User
	facades.Auth(ctx).User(&user)
	return user
}

func First() models.User {
	var user models.User
	facades.Orm().Query().First(&user)
	return user
}
