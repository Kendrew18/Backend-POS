package response

type Read_Transaksi_Inventory_Response struct {
	Kode_transaksi_inventory   string                                     `json:"Kode_transaksi_inventory"`
	Tanggal                    string                                     `json:"tanggal"`
	Kode_nota                  string                                     `json:"kode_nota"`
	Nama_supplier              string                                     `json:"nama_supplier"`
	Nomor_telpon_supplier      string                                     `json:"nomor_telpon_supplier"`
	Harga_ongkos_kirim         int64                                      `json:"harga_ongkos_kirim"`
	Ppn                        float64                                    `json:"ppn"`
	Total_harga                int64                                      `json:"total_harga"`
	Total_barang               float64                                    `json:"total_barang"`
	Status                     int                                        `json:"status"`
	Jenis_transaksi            int                                        `json:"jenis_transaksi"`
	Barang_transaksi_inventory []Read_Barang_Transaksi_Inventory_Response `json:"barang_transaksi_inventory"`
}

type Read_Barang_Transaksi_Inventory_Response struct {
	Kode_barang_transaksi_inventory string  `json:"kode_barang_transaksi_inventory"`
	Kode_inventory                  string  `json:"kode_inventory"`
	Nama_barang                     string  `json:"nama_barang"`
	Jumlah                          float64 `json:"jumlah"`
	Harga                           int64   `json:"harga"`
	Sub_total                       int64   `json:"sub_total"`
}

type Update_Barang_Transaction_Inventory_Response struct {
	Kode_transaksi_inventory string  `json:"kode_transaksi_inventory"`
	Harga_ongkos_kirim       int64   `json:"harga_ongkos_kirim"`
	Ppn                      float64 `json:"ppn"`
	Sub_total                int64   `json:"sub_total"`
	Total_barang             float64 `json:"total_barang"`
	Total_harga              int64   `json:"total_harga"`
}
