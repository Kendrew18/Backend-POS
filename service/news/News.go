package news

import (
	"Bakend-POS/db"
	"Bakend-POS/models/request"
	"Bakend-POS/models/response"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Input_News(Request request.Input_News_Request, writer http.ResponseWriter, request *http.Request) (response.Response, error) {
	var res response.Response

	con := db.CreateConGorm()
	co := 0

	err := con.Table("news").Select("co").Order("co DESC").Limit(1).Scan(&co)

	Request.Co = co + 1
	Request.Kode_news = "NW-" + strconv.Itoa(Request.Co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	date, _ := time.Parse("02-01-2006", Request.Date)
	Request.Date = date.Format("2006-01-02")

	request.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err2 := request.FormFile("photo")

	if file != nil {

		defer file.Close()

		fmt.Println("File Info")
		fmt.Println("File Name : ", handler.Filename)
		fmt.Println("File Size : ", handler.Size)
		fmt.Println("File Type : ", handler.Header.Get("Content-Type"))

		var tempFile *os.File
		path := ""

		if strings.Contains(handler.Filename, "jpg") {
			path = "uploads/foto_inventory/" + Request.Kode_news + ".jpg"
			tempFile, err2 = ioutil.TempFile("uploads/foto_news/", "Read"+"*.jpg")
		}
		if strings.Contains(handler.Filename, "jpeg") {
			path = "uploads/foto_inventory/" + Request.Kode_news + ".jpeg"
			tempFile, err2 = ioutil.TempFile("uploads/foto_inventory/", "Read"+"*.jpeg")
		}
		if strings.Contains(handler.Filename, "png") {
			path = "uploads/foto_inventory/" + Request.Kode_news + ".png"
			tempFile, err2 = ioutil.TempFile("uploads/foto_inventory/", "Read"+"*.png")
		}

		if err2 != nil {
			return res, err2
		}

		fileBytes, err2 := ioutil.ReadAll(file)
		if err2 != nil {
			return res, err2
		}

		_, err2 = tempFile.Write(fileBytes)
		if err2 != nil {
			return res, err2
		}

		fmt.Println("Success!!")
		fmt.Println(tempFile.Name())
		tempFile.Close()

		err2 = os.Rename(tempFile.Name(), path)
		if err2 != nil {
			fmt.Println(err)
		}

		defer tempFile.Close()

		fmt.Println("new path:", tempFile.Name())

		Request.Image_path = path
	} else {
		Request.Image_path = "uploads/foto_news/box.jpg"
	}

	err = con.Table("news").Select("co", "kode_news", "date", "title", "image_path").Create(&Request)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	err = con.Table("content").Select("co").Limit(1).Order("co DESC").Scan(&co)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	for i := 0; i < len(Request.Content); i++ {
		Request.Content[i].Co = co + 1 + i
		Request.Content[i].Kode_content = "CT-" + strconv.Itoa(Request.Content[i].Co)
		Request.Content[i].Kode_news = Request.Kode_news
	}

	err = con.Table("content").Select("co", "kode_news", "kode_content", "content").Create(&Request.Content)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	} else {
		res.Status = http.StatusOK
		res.Message = "Suksess"
		res.Data = map[string]int64{
			"rows": err.RowsAffected,
		}
	}

	return res, nil
}

func Read_News(Request request.Read_News_Request) (response.Response, error) {
	var res response.Response
	var arr_invent []response.Read_News_Response

	con := db.CreateConGorm()

	tanggal := time.Now()
	tanggal_sql := tanggal.Format("2006-01-02")

	err := con.Table("news").Select("kode_news", "DATE_FORMAT(date, '%d-%m-%Y') AS date", "image_path", "title").Where("date <= ?", tanggal_sql).Scan(&arr_invent)

	if err.Error != nil {
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = Request
		return res, err.Error
	}

	for i := 0; i < len(arr_invent); i++ {
		var temp []response.Read_Content_Response
		err := con.Table("content").Select("kode_content", "content").Where("kode_news = ?", arr_invent[i].Kode_news).Order("co ASC").Scan(&temp)

		content := ``

		for i := 0; i < len(temp); i++ {
			content += temp[i].Content

			if i < len(temp)-1 {
				content += "\\n"
			}
		}

		arr_invent[i].Content = content

		fmt.Println(arr_invent[i].Content)

		if err.Error != nil {
			res.Status = http.StatusNotFound
			res.Message = "Status Not Found"
			res.Data = Request
			return res, err.Error
		}
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
