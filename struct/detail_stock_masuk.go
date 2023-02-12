package _struct

type Detail_Stock_Masuk struct {
	Kode_stock    string  `json:"kode_stock"`
	Nama_barang   string  `json:"nama_barang"`
	Jumlah_barang float64 `json:"jumlah_barang"`
	Satuan_barang string  `json:"satuan_barang"`
	Harga_barang  int     `json:"harga_barang"`
}

type Detail_Stock_Masuk_String struct {
	Kode_stock    string `json:"kode_stock"`
	Nama_barang   string `json:"nama_barang"`
	Jumlah_barang string `json:"jumlah_barang"`
	Satuan_barang string `json:"satuan_barang"`
	Harga_barang  string `json:"harga_barang"`
}
