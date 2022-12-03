package pkg

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *sql.DB
}
type (
	FamilyHandler interface {
		GetUsers(c *gin.Context)
		CreateFamily(c *gin.Context)
		DeleteById(c *gin.Context)
		UpdateFamily(c *gin.Context)
		GetByID(c *gin.Context)
	}
)

func NewHandler(db *sql.DB) FamilyHandler {
	return &Handler{
		db: db,
	}
}
