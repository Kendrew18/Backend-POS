package response

type Read_News_Response struct {
	Kode_news  string `json:"kode_news"`
	Date       string `json:"date"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Image_path string `json:"image_path"`
}
