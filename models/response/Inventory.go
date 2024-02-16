package response

type Read_Inventory_Response struct {
	Kode_inventory string  `json:"kode_inventory"`
	Nama_barang    string  `json:"nama_barang"`
	Jumlah_barang  float64 `json:"jumlah_barang"`
	Harga_jual     int64   `json:"harga_jual"`
	Satuan_barang  string  `json:"satuan_barang"`
}

type Check_Nama_Inventory_Response struct {
	Nama_barang string `json:"nama_barang"`
	Status      bool   `json:"status"`
}
