package main

import (
	"lightup/src/common/config"
	"lightup/src/common/db"
	"lightup/src/global"
)

func main() {
	config.Init()
	db.Init()
	global.InitRouter()
}
