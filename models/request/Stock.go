package request

type Input_Stock_Request struct {
	Co            int     `json:"co"`
	Kode_stock    string  `json:"kode_stock"`
	Nama_barang   string  `json:"nama_barang"`
	Jumlah_barang float64 `json:"jumlah_barang"`
	Harga_barang  int64   `json:"harga_barang"`
	Satuan_barang string  `json:"satuan_barang"`
	Kode_user     string  `json:"kode_user"`
}

type Read_Stock_Request struct {
	Kode_user string `json:"kode_user"`
}

type Read_Stock_Filter_Request struct {
	Status_ASC_DESC int    `json:"status_asc_desc"`
	Tanggal_awal    string `json:"tanggal_awal"`
	Tanggal_akhir   string `json:"tanggal_akhir"`
}

type Update_Stock_Request struct {
	Kode_stock    string  `json:"kode_stock"`
	Nama_barang   string  `json:"nama_barang"`
	Jumlah_barang float64 `json:"jumlah_barang"`
	Harga_barang  int64   `json:"harga_barang"`
	Satuan_barang string  `json:"satuan_barang"`
}

type Check_Nama_Stock_Request struct {
	Kode_stock  string `json:"kode_stock"`
	Nama_barang string `json:"nama_barang"`
	Kode_user   string `json:"kode_user"`
}
