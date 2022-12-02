package bl

import (
	"lightup/src/common/log"
)

type ServingImpl struct {
	log log.Logger
}

type ServingBl interface {
}

func New() ServingBl {
	return &ServingImpl{
		log: log.GetLogger("ServingBl"),
	}
}
