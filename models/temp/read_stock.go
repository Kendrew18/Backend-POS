package _struct

type Read_Stock struct {
	Kode_stock    string  `json:"kode_inventory"`
	Nama_barang   string  `json:"nama_barang"`
	Jumlah_barang float64 `json:"jumlah_barang"`
	Satuan_barang string  `json:"satuan_barang"`
	Harga_barang  int     `json:"harga_barang"`
}
