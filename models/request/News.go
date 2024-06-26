package request

type Input_News_Request struct {
	Co         int       `json:"co"`
	Kode_news  string    `json:"kode_news"`
	Date       string    `json:"date"`
	Title      string    `json:"title"`
	Content    []Content `json:"content" gorm:"-"`
	Image_path string    `json:"image_path"`
}

type Read_News_Request struct {
	Kode_user string `json:"kode_user"`
}

type Content struct {
	Co           int    `json:"co"`
	Kode_content string `json:"kode_content"`
	Kode_news    string `json:"kode_news"`
	Content      string `json:"content"`
}

type Delete_News_Request struct {
	Kode_news string `json:"kode_news"`
}
