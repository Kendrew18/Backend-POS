package response

type Read_Inventory_Response struct {
	Kode_inventory   string                     `json:"kode_inventory"`
	Nama_barang      string                     `json:"nama_barang"`
	Jumlah_barang    float64                    `json:"jumlah_barang"`
	Harga_jual       int64                      `json:"harga_jual"`
	Satuan_barang    string                     `json:"satuan_barang"`
	Detail_inventory []Detail_Iventory_Response `json:"detail_inventory" gorm:"-"`
}

type Detail_Iventory_Response struct {
	Kode_barang_transaksi_inventory string  `json:"kode_barang_transaksi_inventory"`
	Kode_transaksi_inventory        string  `json:"kode_transaksi_inventory"`
	Nama_supplier                   string  `json:"nama_supplier"`
	Jumlah                          float64 `json:"jumlah"`
	Harga                           int64   `json:"harga"`
}

type Check_Nama_Inventory_Response struct {
	Nama_barang string `json:"nama_barang"`
	Status      bool   `json:"status"`
}
