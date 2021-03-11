package apis

import (
	"database/sql"
	"golang_mysql/database"
	"golang_mysql/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetCompany(c *gin.Context) {
	var arr []models.ArrComp
	db := database.InitialDB()
	arr = GetDataCompany(db)
	defer db.Close()
	c.JSON(http.StatusOK, gin.H{"result": arr})
}

func GetDataCompany(db *sql.DB) []models.ArrComp {
	var CompData []models.ArrComp
	results, err := db.Query("select * from company_tb where id_comp=?", 1)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag models.TagCompany
		err = results.Scan(&tag.Id_comp, &tag.Name_comp_th, &tag.Name_comp_en, &tag.Noaddress, &tag.Soi, &tag.Road, &tag.District, &tag.Amphur, &tag.Province, &tag.Zipcode, &tag.Tel_call, &tag.Tel_fax, &tag.Ss_number, &tag.Tax_number, &tag.Bank_number, &tag.Co_number, &tag.Fund_number, &tag.Status_comp)

		if err != nil {
			panic(err.Error())
		}

		CompData = append(CompData, models.ArrComp{
			Id_comp:      tag.Id_comp,
			Name_comp_th: tag.Name_comp_th,
			Name_comp_en: tag.Name_comp_en,
			Noaddress:    tag.Noaddress,
			Soi:          tag.Soi,
			Road:         tag.Road,
			District:     tag.District,
			Amphur:       tag.Amphur,
			Province:     tag.Province,
			Zipcode:      tag.Zipcode,
			Tel_call:     tag.Tel_call,
			Tel_fax:      tag.Tel_fax,
			Ss_number:    tag.Ss_number,
			Tax_number:   tag.Tax_number,
			Bank_number:  tag.Bank_number,
			Co_number:    tag.Co_number,
			Fund_number:  tag.Fund_number,
			Status_comp:  tag.Status_comp,
		})
	}
	return CompData
}
