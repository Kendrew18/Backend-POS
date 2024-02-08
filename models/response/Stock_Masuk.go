package response

type Read_Stock_Masuk_Response struct {
	Kode_stock_masuk   string                             `json:"kode_stock_masuk"`
	Tanggal_masuk      string                             `json:"Tanggal_masuk"`
	Kode_supplier      string                             `json:"kode_supplier"`
	Nama_supplier      string                             `json:"nama_supplier"`
	Total_barang       float64                            `json:"total_barang"`
	Total_harga        int64                              `json:"total_harga"`
	Barang_stock_masuk []Read_Barang_Stock_Masuk_Response `json:"barang_stock_masuk"`
}

type Read_Barang_Stock_Masuk_Response struct {
	Kode_barang_stock_masuk string  `json:"kode_barang_stock_masuk"`
	Kode_stock              string  `json:"kode_stock"`
	Nama_barang             string  `json:"nama_barang"`
	Jumlah                  float64 `json:"jumlah"`
	Harga                   int64   `json:"harga"`
	Sub_total               int64   `json:"sub_total"`
}
