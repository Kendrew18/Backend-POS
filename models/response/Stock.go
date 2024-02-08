package response

type Read_Stock_Response struct {
	Kode_stock    string  `json:"kode_stock"`
	Nama_barang   string  `json:"nama_barang"`
	Jumlah_barang float64 `json:"jumlah_barang"`
	Harga_barang  int64   `json:"harga_barang"`
	Satuan_barang string  `json:"satuan_barang"`
}

type Check_Nama_Stock_Response struct {
	Nama_barang string `json:"nama_barang"`
	Status      bool   `json:"status"`
}
