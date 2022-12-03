package main

import (
	"Dadaxon-Olimov/connaction"
	"Dadaxon-Olimov/pkg"
	"Dadaxon-Olimov/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	server, db := gin.Default(), connaction.Connection()

	handler := pkg.NewHandler(db)

	fmt.Println("hasten started ...")
	router.Routers(server, handler)
	server.Run("localhost:1217")
	defer db.Close()
}
