package _struct

type Insert_Stock struct {
	Nama_barang   string  `json:"nama_barang"`
	Jumlah_barang float64 `json:"jumlah_barang"`
	Satuan_barang string  `json:"satuan_barang"`
	Harga_barang  int     `json:"harga_barang"`
}
