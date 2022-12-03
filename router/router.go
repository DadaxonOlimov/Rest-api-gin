package router

import (
	"Dadaxon-Olimov/pkg"
	"github.com/gin-gonic/gin"
)

func Routers(c *gin.Engine, h pkg.FamilyHandler) {
	c.GET("/family", h.GetUsers)
	c.POST("/family", h.CreateFamily)
	c.DELETE("/family/:id", h.DeleteById)
	c.PUT("/family/:id", h.UpdateFamily)
	c.GET("/family/:id", h.GetByID)
}
