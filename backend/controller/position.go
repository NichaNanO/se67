package controller

import (
   "net/http"
   "example.com/project/config"
	"example.com/project/entity"
   "github.com/gin-gonic/gin"
)

func GetPositions(c *gin.Context) {
   var positions []entity.Position
   
   db := config.DB()
   db.Find(&positions)
   c.JSON(http.StatusOK, &positions)
}