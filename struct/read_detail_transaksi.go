package _struct

type Read_Detail_Transaksi struct {
	Kode_stock    string  `json:"kode_stock"`
	Nama_barang   string  `json:"nama_barang"`
	Jumlah_barang float64 `json:"jumlah_barang"`
	Satuan_barang string  `json:"satuan_barang"`
	Harga_barang  int     `json:"harga_barang"`
	Harga_satuan  int     `json:"harga_satuan"`
}

type Read_Detail_Transaksi_String struct {
	Kode_stock    string `json:"kode_stock"`
	Nama_barang   string `json:"nama_barang"`
	Jumlah_barang string `json:"jumlah_barang"`
	Satuan_barang string `json:"satuan_barang"`
	Harga_barang  string `json:"harga_barang"`
	Harga_satuan  string `json:"harga_satuan"`
}
