package request

type Input_Stock_Masuk_Request struct {
	Co                    int    `json:"co"`
	Kode_stock_masuk      string `json:"kode_stock_masuk"`
	Nama_penanggung_jawab string `json:"nama_penanggung_jawab"`
	Tanggal_masuk         string `json:"tanggal_masuk"`
	Kode_supplier         string `json:"kode_supplier"`
	Kode_user             string `json:"kode_user"`
}

type Input_Barang_Stock_Masuk_Request struct {
	Kode_stock string `json:"kode_stock"`
	Jumlah     string `json:"jumlah"`
	Harga      string `json:"harga"`
}

type Input_Barang_Stock_Masuk_V2_Request struct {
	Co                      int     `json:"co"`
	Kode_barang_stock_masuk string  `json:"kode_barang_stock_masuk"`
	Kode_stock_masuk        string  `json:"kode_stock_masuk"`
	Kode_stock              string  `json:"kode_stock"`
	Jumlah                  float64 `json:"jumlah"`
	Harga                   int64   `json:"harga"`
	Sub_total               int64   `json:"sub_total"`
}

type Read_Stock_Masuk_Request struct {
	Kode_user string `json:"kode_user"`
}

type Read_Stock_Masuk_Filter_Request struct {
	Tanggal_awal  string `json:"tanggal_awal"`
	Tanggal_akhir string `json:"tanggal_akhir"`
	Kode_supplier string `json:"kode_supplier"`
}

type Update_Kode_Barang_Stock_Masuk_Request struct {
	Kode_barang_stock_masuk string `json:"kode_barang_stock_masuk"`
}

type Update_Stock_Masuk_Request struct {
	Kode_stock string  `json:"kode_stock"`
	Jumlah     float64 `json:"jumlah"`
	Harga      int64   `json:"harga"`
	Sub_total  int64   `json:"sub_total"`
}

type Update_Status_Stock_Masuk_Request struct {
	Status int `json:"status"`
}

type Update_Kode_Stock_Masuk_Request struct {
	Kode_stock_masuk string `json:"kdoe_stock_masuk"`
}
