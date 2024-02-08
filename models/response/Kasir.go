package response

type Read_Kasir_Response struct {
	Kode_stock    string  `json:"kode_stock"`
	Nama_barang   string  `json:"nama_barang"`
	Jumlah_barang float64 `json:"jumlah_barang"`
	Satuan_barang string  `json:"satuan_barang"`
	Harga_barang  int64   `json:"harga_barang"`
}
