package response

type Read_Kasir_Response struct {
	Kode_inventory   string                           `json:"kode_inventory"`
	Nama_barang      string                           `json:"nama_barang"`
	Jumlah_barang    float64                          `json:"jumlah_barang"`
	Satuan_barang    string                           `json:"satuan_barang"`
	Harga_jual       int64                            `json:"harga_jual"`
	Path_photo       string                           `json:"path_photo"`
	Detail_inventory []Read_Detail_Inventory_Response `json:"detail_inventory" gorm:"-"`
}

type Read_Detail_Inventory_Response struct {
	Kode_barang_transaksi_inventory string  `json:"kode_barang_transaksi_inventory"`
	Kode_inventory                  string  `json:"kode_inventory"`
	Tanggal                         string  `json:"tanggal"`
	Jumlah                          float64 `json:"jumlah"`
}
