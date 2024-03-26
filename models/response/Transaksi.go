package response

type Read_Transaksi_Response struct {
	Kode_transaksi        string                           `json:"kode_transaksi"`
	Kode_nota             string                           `json:"kode_nota"`
	Tanggal               string                           `json:"tanggal"`
	Nama_customer         string                           `json:"nama_customer"`
	Alamat_customer       string                           `json:"alamat_customer"`
	Nomer_telp_customer   string                           `json:"nomer_telp_customer"`
	Kode_jenis_pembayaran string                           `json:"kode_jenis_pembayaran"`
	Nama_jenis_pembayaran string                           `json:"nama_jenis_pembayaran"`
	Jumlah_total          float64                          `json:"jumlah_total"`
	Total_harga           int64                            `json:"total_harga"`
	Diskon                int64                            `json:"diskon"`
	Barang_transaksi      []Read_Barang_Transaksi_Response `json:"barang_transaksi" gorm:"-"`
}

type Read_Barang_Transaksi_Response struct {
	Kode_barang_transaksi           string  `json:"kode_barang_transaksi"`
	Kode_barang_transaksi_inventory string  `json:"kode_barang_transaksi_inventory"`
	Kode_inventory                  string  `json:"kode_inventory"`
	Nama_barang                     string  `json:"nama_barang"`
	Jumlah_barang                   float64 `json:"jumlah_barang"`
	Nama_satuan                     string  `json:"nama_satuan"`
	Harga                           int64   `json:"harga"`
	Sub_total                       int64   `json:"sub_total"`
}

type Read_Jumlah_Invent_Response struct {
	Kode_inventory string  `json:"kode_inventory"`
	Jumlah_barang  float64 `json:"jumlah_barang"`
}
