package main

import (
	"lightup/src/common/db"
	"lightup/src/router"
)

func main() {
	db.InitDB()
	router.InitRouter()
}
