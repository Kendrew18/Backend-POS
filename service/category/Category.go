package category

import (
	"Bakend-POS/db"
	"Bakend-POS/models/response"
	"net/http"
)

func Read_Category() (response.Response, error) {
	var res response.Response
	var arr_invent []response.Read_Category_Response

	con := db.CreateConGorm()

	err := con.Table("category").Select("kode_category", "nama_category").Order("nama_category ASC").Scan(&arr_invent)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		return res, err.Error
	}

	if arr_invent == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_invent
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_invent
	}

	return res, nil
}
