package bl

import (
	"github.com/golang-jwt/jwt/v4"
	"lightup/src/common/hasher"
	"lightup/src/common/http"
	"lightup/src/common/log"
	"lightup/src/modules/user/dal"
	"time"
)

type AuthImpl struct {
	jwtSecret []byte
	log       log.Logger
	userRepo  *dal.UserRepo
	hasher    hasher.Hasher
}

type AuthBl interface {
	CreateToken(email string, password string) (string, error)
}

func NewAuth() *AuthImpl {
	return &AuthImpl{
		jwtSecret: []byte("secret"),
		log:       log.GetLogger("AuthBl"),
		userRepo:  dal.New(),
		hasher:    hasher.New(),
	}
}

func (impl *AuthImpl) CreateToken(email string, password string) (string, error) {
	user, err := impl.userRepo.FindByEmail(email)

	if err != nil {
		return "", err
	}

	if user == nil || !impl.hasher.CheckHash(password, user.Password) {
		return "", http.Error{
			StatusCode: 409,
			Message:    "Invalid password",
		}
	}

	token, err := impl.createUserJwtToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (impl *AuthImpl) createUserJwtToken(user *dal.UserEntity) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID.Hex(),
		"exp":    time.Now().AddDate(1, 0, 0).Unix(),
	})

	return token.SignedString(impl.jwtSecret)
}
