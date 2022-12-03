package pkg

import (
	"Dadaxon-Olimov/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) GetByID(c *gin.Context) {
	fmt.Println("________________________________---")
	id := c.Param("id")
	rows, err := h.db.Query("select * from family where id = $1", id)
	fmt.Println("________________________________---")
	if err != nil {
		panic(err)
	}
	prod := model.Family{}
	rows.Next()
	err = rows.Scan(&prod.Id, &prod.Mother, &prod.Brother, &prod.Prise)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"massage": "user not found"})
	} else {
		c.IndentedJSON(http.StatusOK, prod)
	}
}
