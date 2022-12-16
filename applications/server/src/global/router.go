package global

import (
	"lightup/src/common/config"
	account_ctrl "lightup/src/modules/account/ctrl"
	ff_ctrl "lightup/src/modules/feature_flag/ctrl"
	serving_ctrl "lightup/src/modules/serving/ctrl"
	tag_ctrl "lightup/src/modules/tag/ctrl"
	user_ctrl "lightup/src/modules/user/ctrl"
	"strconv"

	"github.com/gin-gonic/gin"
)

type routerConfig struct {
	Mode string `mapstructure:"mode"`
	Port int    `mapstructure:"port"`
}

var getConfig = config.UnmarshalKey
var conf = &routerConfig{}

func InitRouter() {
	getConfig("router", conf)
	route := gin.Default()

	if conf.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	v1 := route.Group("api/v1")
	{
		ff_ctrl.New().Init(v1)
		tag_ctrl.New().Init(v1)
		account_ctrl.New().Init(v1)
		user_ctrl.New().Init(v1)
		serving_ctrl.New().Init(v1)
	}

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	route.Use(gin.Recovery())
	route.SetTrustedProxies([]string{})
	route.Run(":" + strconv.Itoa(conf.Port))
}
