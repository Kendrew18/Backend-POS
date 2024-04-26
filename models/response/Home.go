package response

type Home_Response struct {
	Total_pemasukan          int64               `json:"total_pemasukan"`
	Total_pengeluaran        int64               `json:"total_pengeluaran"`
	Total_pembayaran_pending int64               `json:"total_pembayaran_pending"`
	Chart_Pemasukan          []Chart_Pemasukan   `json:"chart_pemasukan" gorm:"-"`
	Chart_Pengeluaran        []Chart_Pengeluaran `json:"chart_pengeluaran" gorm:"-"`
}
type Chart_Pengeluaran struct {
	Tanggal       string `json:"tanggal"`
	Tanggal_tahun string `json:"tanggal_tahun"`
	Value         int64  `json:"Value"`
}

type Chart_Pemasukan struct {
	Tanggal       string `json:"tanggal"`
	Tanggal_tahun string `json:"tanggal_tahun"`
	Value         int64  `json:"Value"`
}
