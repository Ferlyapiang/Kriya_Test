package main

import (
	"kriya_Test/routes/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	user.User(r)

	r.Run() //  serve on ("localhost:8080")
}
