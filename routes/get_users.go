package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restApi/constants"
	"strconv"
)

func GetUser(users *[]constants.User) func(c *gin.Context) {
	return func(c *gin.Context) {
		if len(*users) > 0 {
			c.IndentedJSON(http.StatusOK, *users)
		} else {
			c.IndentedJSON(http.StatusOK, []constants.User{})
		}
	}
}

func GetUserById(users *[]constants.User) func(*gin.Context) {
	return func(c *gin.Context) {
		idstr := c.Param("id")
		var id int
		if idstr != "" {
			id, _ = strconv.Atoi(idstr)
		} else {
			c.IndentedJSON(http.StatusBadRequest, nil)
		}

		index, found := constants.GetIndexFromID(*users, id)

		if found {
			c.IndentedJSON(http.StatusOK, (*users)[index])
		} else {
			c.IndentedJSON(http.StatusNotFound, nil)
		}
	}
}
