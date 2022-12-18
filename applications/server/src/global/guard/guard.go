package guard

import (
	app_model "lightup/src/common/model"
)

type Guard interface {
	IsActive(c *app_model.ReqContext) error
}
