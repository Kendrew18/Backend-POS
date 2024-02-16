package request

type Input_Transaksi_Inventory_Request struct {
	Co                       int     `json:"co"`
	Kode_transaksi_inventory string  `json:"kode_transaksi_inventory"`
	Tanggal                  string  `json:"tanggal"`
	Kode_transaksi           string  `json:"kode_transaksi"`
	Kode_jenis_pembayaran    string  `json:"kode_jenis_pembayaran"`
	Harga_ongkos_kirim       int64   `json:"harga_ongkos_kirim"`
	Ppn                      float64 `json:"ppn"`
	Kode_supplier            string  `json:"kode_supplier"`
	Kode_user                string  `json:"kode_user"`
	Jenis_transaksi          string  `json:"jenis_transaksi"`
}

type Input_Barang_Transaksi_Inventory_Request struct {
	Kode_inventory string `json:"kode_inventory"`
	Jumlah         string `json:"jumlah"`
	Harga          string `json:"harga"`
}

type Input_Barang_Transaksi_Inventory_V2_Request struct {
	Co                              int     `json:"co"`
	Kode_barang_transaksi_inventory string  `json:"kode_barang_transaksi_inventory"`
	Kode_transaksi_inventory        string  `json:"kode_transaksi_inventory"`
	Kode_inventory                  string  `json:"kode_inventory"`
	Jumlah                          float64 `json:"jumlah"`
	Harga                           int64   `json:"harga"`
	Sub_total                       int64   `json:"sub_total"`
}

type Read_Transaksi_Inventory_Request struct {
	Kode_user string `json:"kode_user"`
}

type Read_Transaksi_Inventory_Filter_Request struct {
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
