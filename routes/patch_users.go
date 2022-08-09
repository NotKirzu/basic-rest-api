package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restApi/constants"
	"strconv"
)

func ModifyUser(users *[]constants.User) func(*gin.Context) {
	return func(c *gin.Context) {
		idstr := c.Param("id")
		var id int
		if idstr != "" {
			id, _ = strconv.Atoi(idstr)
		} else {
			c.IndentedJSON(http.StatusBadRequest, nil)
		}

		var newUser constants.User

		if err := c.BindJSON(&newUser); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err)
			return
		}

		index, found := constants.GetIndexFromID(*users, id)

		if found {
			if newUser.FirstName != "" {
				(*users)[index].FirstName = newUser.FirstName
			}
			if newUser.LastName != "" {
				(*users)[index].LastName = newUser.LastName
			}
			if newUser.DNI != "" {
				(*users)[index].DNI = newUser.DNI
			}
			c.IndentedJSON(http.StatusNoContent, newUser)
		} else {
			c.IndentedJSON(http.StatusNotFound, nil)
		}
	}
}
