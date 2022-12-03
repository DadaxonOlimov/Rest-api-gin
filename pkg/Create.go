package pkg

import (
	"Dadaxon-Olimov/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) CreateFamily(c *gin.Context) {
	var newFamily model.Family
	if err := c.BindJSON(&newFamily); err != nil {
		panic(err)
	}
	fmt.Println("-----------", newFamily)
	err := h.db.QueryRow("INSERT INTO family (mother,brother,prise)VALUES ($1,$2,$3) returning id",
		newFamily.Mother, newFamily.Brother, newFamily.Prise).Scan(&newFamily.Id)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, newFamily)
}
