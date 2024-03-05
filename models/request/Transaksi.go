package request

type Input_Transaksi_Request struct {
	Co                    int     `json:"co"`
	Kode_transaksi        string  `json:"kode_pembayaran"`
	Kode_nota             string  `json:"kode_nota"`
	Tanggal               string  `json:"tanggal"`
	Kode_jenis_pembayaran string  `json:"kode_jenis_pembayaran"`
	Jumlah_total          float64 `json:"jumlah_total"`
	Total_harga           int64   `json:"total_harga"`
	Diskon                int64   `json:"diskon"`
	Kode_user             string  `json:"kode_user"`
	Uuid_session          string  `json:"uuid_session"`
}

type Input_Barang_Transaksi_Request struct {
	Co                              int    `json:"co"`
	Kode_transaksi                  string `json:"kode_transaksi"`
	Kode_inventory                  string `json:"kode_inventory"`
	Kode_barang_transaksi_inventory string `json:"kode_barang_transaksi_inventory"`
	Jumlah_barang                   int    `json:"jumlah_barang"`
	Nama_satuan                     string `json:"nama_satuan"`
	Harga                           int64  `json:"harga"`
	Sub_total                       int64  `json:"sub_total"`
}

type Body_Input_Transaksi_Request struct {
	Input_transaksi        Input_Transaksi_Request        `json:"input_transaksi"`
	Input_barang_transaksi Input_Barang_Transaksi_Request `json:"input_barang_transaksi"`
}
