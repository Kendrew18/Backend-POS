package _struct

type Detail_Stock_Masuk struct {
	Kode_stock    string `json:"kode_stock"`
	Jumlah_barang int    `json:"jumlah_barang"`
	Harga_barang  int    `json:"harga_barang"`
}

type Detail_Stock_Masuk_String struct {
	Kode_stock    string `json:"kode_stock"`
	Jumlah_barang string `json:"jumlah_barang"`
	Harga_barang  string `json:"harga_barang"`
}
