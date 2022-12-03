package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) DeleteById(c *gin.Context) {
	id := c.Param("id")
	rows, err := h.db.Query("DELETE FROM family WHERE id = $1", id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "users not found"})
	} else {
		c.IndentedJSON(http.StatusOK, rows)
	}

}
