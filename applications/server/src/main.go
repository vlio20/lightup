package main

import (
	"lightup/src/db"
	"lightup/src/router"
)

func main() {
	db.InitDB()
	router.InitRouter()
}
