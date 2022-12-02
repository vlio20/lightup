package api

import (
	"lightup/src/modules/serving/bl"
)

type ServingApi struct {
	servingBl bl.ServingBl
}

func New() *ServingApi {
	return &ServingApi{
		servingBl: bl.New(),
	}
}
