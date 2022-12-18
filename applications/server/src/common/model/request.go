package app_model

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lightup/src/modules/user/dal"
)

type ReqContext struct {
	*gin.Context
	AccountID primitive.ObjectID
	User      *dal.UserEntity
}
