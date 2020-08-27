package api

import (
	"go-rest-api/db/slice_db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NumbersGet(db *slice_db.Db) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := db.GetAll()
		c.JSON(http.StatusOK, result)
	}
}
