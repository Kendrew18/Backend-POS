package request

type Input_Transaksi_Request struct {
	Co                    int     `json:"co"`
	Kode_transaksi        string  `json:"kode_transaksi"`
	Kode_nota             string  `json:"kode_nota"`
	Tanggal               string  `json:"tanggal"`
	Kode_jenis_pembayaran string  `json:"kode_jenis_pembayaran"`
	Jumlah_total          float64 `json:"jumlah_total"`
	Total_harga           int64   `json:"total_harga"`
	Diskon                int64   `json:"diskon"`
	Nama_customer         string  `json:"nama_customer"`
	Alamat_customer       string  `json:"alamat_customer"`
	Nomer_telp_customer   string  `json:"nomer_telp_customer"`
	Kode_user             string  `json:"kode_user"`
}

type Input_Barang_Transaksi_Request struct {
	Co                              int     `json:"co"`
	Kode_barang_transaksi           string  `json:"kode_barang_transaksi"`
	Kode_transaksi                  string  `json:"kode_transaksi"`
	Kode_barang_transaksi_inventory string  `json:"kode_barang_transaksi_inventory"`
	Kode_inventory                  string  `json:"kode_inventory"`
	Jumlah_barang                   float64 `json:"jumlah_barang"`
	Nama_satuan                     string  `json:"nama_satuan"`
	Harga                           int64   `json:"harga"`
	Sub_total                       int64   `json:"sub_total"`
}

type Body_Input_Transaksi_Request struct {
	Input_transaksi        Input_Transaksi_Request          `json:"input_transaksi"`
	Input_barang_transaksi []Input_Barang_Transaksi_Request `json:"input_barang_transaksi"`
}

type Body_Read_Transaksi_Request struct {
	Read_transaksi        Read_Transaksi_Request        `json:"read_transaksi"`
	Read_transaksi_filter Read_Transaksi_Filter_Request `json:"read_transaksi_filter"`
}

type Read_Transaksi_Request struct {
	Kode_user string `json:"kode_user"`
}

type Read_Transaksi_Filter_Request struct {
	Tanggal_awal  string `json:"tanggal_awal"`
	Tanggal_akhir string `json:"tanggal_akhir"`
	Nama_customer string `json:"nama_customer"`
}
