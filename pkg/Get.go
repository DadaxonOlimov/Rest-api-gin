package pkg

import (
	"Dadaxon-Olimov/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) GetUsers(c *gin.Context) {
	rows, err := h.db.Query("SELECT * FROM family")
	if err != nil {
		fmt.Println(";;;;;;;;;;;____________")
		panic(err)
	}
	defer rows.Close()
	fmt.Println(";;;;;;;;;;;____________")
	var family []model.Family
	var f model.Family
	for rows.Next() {
		err := rows.Scan(&f.Id, &f.Mother, &f.Brother, &f.Prise)
		fmt.Println("err")
		if err != nil {
			fmt.Println(err)
			continue
		}
		family = append(family, f)
	}
	c.IndentedJSON(http.StatusOK, family)
}
