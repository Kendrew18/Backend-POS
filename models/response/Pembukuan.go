package response

type Read_Pembukuan_Response struct {
	Kode_pembukuan        string                           `json:"kode_pembukuan"`
	Tanggal               string                           `json:"tanggal"`
	Kode_nota             string                           `json:"kode_nota"`
	Kode_jenis_pembayaran string                           `json:"kode_jenis_pembayaran"`
	Nama_jenis_pembayaran string                           `json:"nama_jenis_pembayaran"`
	Diskon                int64                            `json:"diskon"`
	Total_harga           int64                            `json:"total_harga"`
	Total_barang          float64                          `json:"total_barang"`
	Barang_pembukuan      []Read_Barang_Pembukuan_Response `json:"barang_pembukuan" gorm:"-"`
}
type Read_Barang_Pembukuan_Response struct {
	Kode_stock string  `json:"kode_stock"`
	Nama_stock string  `json:"nama_stock"`
	Jumlah     float64 `json:"jumlah"`
	Harga      int64   `json:"harga"`
	Satuan     string  `json:"satuan"`
	Sub_total  string  `json:"sub_total"`
}
