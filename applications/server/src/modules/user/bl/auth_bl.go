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
	GetUserByToken(token string) (*dal.UserEntity, error)
}

func NewAuth() *AuthImpl {
	return &AuthImpl{
		jwtSecret: []byte("secret"),
		log:       log.GetLogger("AuthBl"),
		userRepo:  dal.New(),
		hasher:    hasher.New(),
	}
}

var invalidTokenError = http.Error{
	StatusCode: 401,
	Message:    "Invalid token",
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

func (impl *AuthImpl) GetUserByToken(tokenStr string) (*dal.UserEntity, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidTokenError
		}

		return impl.jwtSecret, nil
	})

	if err != nil {
		return nil, invalidTokenError
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok == false {
		return nil, invalidTokenError
	}

	userId, ok := claims["userId"].(string)

	if ok == false {
		return nil, invalidTokenError
	}

	return impl.userRepo.FindByStrID(userId)
}

func (impl *AuthImpl) createUserJwtToken(user *dal.UserEntity) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID.Hex(),
		"exp":    time.Now().AddDate(1, 0, 0).Unix(),
	})

	return token.SignedString(impl.jwtSecret)
}
