package global

import (
	"lightup/src/common/config"
	"lightup/src/modules/feature_flag/controller"
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
		controller.Init(v1)
	}

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.SetTrustedProxies([]string{})
	route.Run(":" + strconv.Itoa(conf.Port))
}
