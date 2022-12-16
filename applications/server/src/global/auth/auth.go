package guard

import (
	"lightup/src/common/http"
	app_model "lightup/src/common/model"
	"lightup/src/modules/user/bl"
)

var missingToken = http.Error{
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
		return missingToken
	}

	user, err := g.authBl.GetUserByToken(token)

	if err != nil {
		return err
	}

	c.SetUser(user)

	return nil
}
