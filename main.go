package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"restApi/constants"
	"restApi/routes"
)

var users []constants.User
var host = "localhost:8080"

func main() {
	router := gin.Default()
	if err := router.SetTrustedProxies([]string{"127.0.0.1", "192.168.3.6"}); err != nil {
		fmt.Println(err)
		return
	}

	router.GET("/users", routes.GetUser(&users))
	router.GET("/users/:id", routes.GetUserById(&users))
	router.POST("/users", routes.AppendUser(&users))
	router.DELETE("/users/:id", routes.DeleteUser(&users))
	router.PATCH("/users/:id", routes.ModifyUser(&users))

	if err := router.Run(host); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Server running at %s\nPress Ctrl-C to exit...", host)
}
