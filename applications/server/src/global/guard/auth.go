package guard

import (
	"lightup/src/common/http"
	app_model "lightup/src/common/model"
	"lightup/src/modules/user/bl"
)

var missingTokenError = http.Error{
	StatusCode: 401,
	Message:    "Token is missing",
}

type AuthGuard struct {
	authBl bl.AuthBl
}

func NewAuthGuard() *AuthGuard {
	return &AuthGuard{
		authBl: bl.NewAuth(),
	}
}

func (g *AuthGuard) IsActive(c *app_model.ReqContext) error {
	token := c.GetHeader("X-Lightup-Token")

	if token == "" {
		return missingTokenError
	}

	user, err := g.authBl.GetUserByToken(token)

	if err != nil {
		return err
	}

	if user == nil {
		return http.Error{
			StatusCode: 409,
			Message:    "Invalid Token",
		}
	}

	c.User = user
	c.AccountID = user.AccountID

	return nil
}
