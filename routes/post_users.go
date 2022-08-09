package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restApi/constants"
)

func AppendUser(users *[]constants.User) func(*gin.Context) {
	return func(c *gin.Context) {
		var newUser constants.User

		if err := c.BindJSON(&newUser); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
			return
		}

		if len(*users) != 0 {
			newUser.ID = (*users)[len(*users)-1].ID + 1
		} else {
			newUser.ID = 0
		}
		*users = append(*users, newUser)
		c.IndentedJSON(http.StatusOK, newUser)
	}
}
