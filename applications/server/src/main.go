package main

import (
	"lightup/src/common/db"
	"lightup/src/global"
)

func main() {
	db.InitDB()
	global.InitRouter()
}
