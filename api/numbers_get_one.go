package api

import (
	"go-rest-api/db/slice_db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NumbersGetOne(db *slice_db.Db) gin.HandlerFunc {
	return func(c *gin.Context) {
		numString := c.Param("int")
		num, _ := strconv.Atoi(numString)
		result := db.GetOne(num)
		if result == "Number is not found" {
			c.JSON(http.StatusNotFound, result)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}
