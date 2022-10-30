package _struct

type Read_Stock struct {
	Kode_stock    string `json:"kode_inventory"`
	Nama_barang   string `json:"nama_barang"`
	Jumlah_barang string `json:"jumlah_barang"`
	Harga_barang  string `json:"harga_barang"`
}
