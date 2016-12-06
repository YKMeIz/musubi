package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// checkErr check if err variable is empty. If error occurs, let Gin response
// error status and Json body.
func checkErr(err error, c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"Status": err.Error()})
		c.AbortWithStatus(http.StatusNotImplemented)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status": "ok!"})
}
