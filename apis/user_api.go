package apis

import (
	"database/sql"
	"golang_mysql/database"
	"golang_mysql/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func ExecuteAllUser(c *gin.Context) {
	var arr []models.ArrUser
	db := database.InitialDB()
	arr = GetAllUser(db)
	defer db.Close()
	c.JSON(http.StatusOK, gin.H{"result": arr})
}

func ExecuteRowsUser(c *gin.Context) {
	var arr []models.ArrUser
	id := c.Param("id") // For GET METHOD
	db := database.InitialDB()
	results, err := db.Query("SELECT id_emp , id_no_emp , name_th FROM employee_tb WHERE id_emp=?", id)

	for results.Next() {
		var tag models.Tag
		err = results.Scan(&tag.ID, &tag.IDem, &tag.Name)
		if err != nil {
			panic(err.Error())
		}
		arr = append(arr, models.ArrUser{ID: tag.ID, IDEmp: tag.IDem, NameEmp: tag.Name})
	}
	/*
		c.JSON(200, gin.H{
			"status": "success",
			"body":   arr,
		})
	*/
	c.JSON(http.StatusOK, gin.H{"result": arr})
}

func UpdateUser(c *gin.Context) {
	db := database.InitialDB()
	id := c.Param("id")
	name := c.PostForm("name")

	rows, err1 := db.Query("SELECT COUNT(*) FROM employee_tb WHERE name_th = ? AND id_emp != ?", name, id)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		count += 1
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{"result": "Name is " + name + " is already exists."})
		return
	} else {
		updateQry, err := db.Prepare("UPDATE employee_tb SET name_th = ?  WHERE id_emp = ?")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not found update"})
			return
		}
		updateQry.Exec(name, id) // prepare param bypass
		defer db.Close()
		c.JSON(http.StatusOK, gin.H{"result": "Update Successfully"})
	}
}

func GetAllUser(db *sql.DB) []models.ArrUser {
	var Users []models.ArrUser
	results, err := db.Query("SELECT id_emp, id_no_emp , name_th FROM employee_tb")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag models.Tag
		err = results.Scan(&tag.ID, &tag.IDem, &tag.Name)

		if err != nil {
			panic(err.Error())
		}

		Users = append(Users, models.ArrUser{ID: tag.ID, IDEmp: tag.IDem, NameEmp: tag.Name})
	}
	return Users
}
