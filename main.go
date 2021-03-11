package main

import (
	"fmt"
	"golang_mysql/apis"
	"os"

	"github.com/gin-gonic/gin"
)

type Tag struct {
	ID   int    `json:"id_no_emp"`
	Name string `json:"name_th"`
}

type ArrUser struct {
	IDEmp   int
	NameEmp string
}

func main() {
	r := gin.Default()
	r.GET("/getAllUser", apis.ExecuteAllUser)
	r.GET("/getDataCompany", apis.GetCompany)
	r.GET("/selectUser/:id", apis.ExecuteRowsUser)
	r.POST("/upDateUser/:id", apis.UpdateUser)
	r.Run(":"+getPort())

}

func getPort() string {
	var port = os.Getenv("PORT") // ----> (A)
	if port == "" {
	   port = "8080"
	   fmt.Println("No Port In Heroku" + port)
	}
	return ":" + port // ----> (B)
}
