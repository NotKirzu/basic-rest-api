package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restApi/constants"
	"strconv"
)

func DeleteUser(users *[]constants.User) func(*gin.Context) {
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
			constants.RemoveIndex(users, index)
			c.IndentedJSON(http.StatusNoContent, nil)
		} else {
			c.IndentedJSON(http.StatusNotFound, nil)
		}
	}
}
