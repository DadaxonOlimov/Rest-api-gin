package pkg

import (
	"Dadaxon-Olimov/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) UpdateFamily(c *gin.Context) {
	var Update model.Family
	Update.Id = c.Param("id")
	if err := c.BindJSON(&Update); err != nil {
		fmt.Sprintf("invalid JSON body: %v", err)
		panic(err)
	}
	_, err := h.db.Exec("UPDATE family SET mother =$1, brother = $2, prise = $3 WHERE id =$4",
		Update.Mother, Update.Brother, Update.Prise, Update.Id)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, Update)

}
