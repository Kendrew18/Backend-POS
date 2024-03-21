package response

type Home_Response struct {
	Total_pemasukan          int64               `json:"total_pemasukan"`
	Total_pengeluaran        int64               `json:"total_pengeluaran"`
	Total_pembayaran_pending int64               `json:"total_pembayaran_pending"`
	Chart_Pemasukan          []Chart_Pemasukan   `json:"chart_pemasukan" gorm:"-"`
	Chart_Pengeluaran        []Chart_Pengeluaran `json:"chart_pengeluaran" gorm:"-"`
}
type Chart_Pengeluaran struct {
	Tanggal           string `json:"tanggal"`
	Total_pengeluaran int64  `json:"total_pengeluaran"`
}

type Chart_Pemasukan struct {
	Tanggal         string `json:"tanggal"`
	Total_pemasukan int64  `json:"total_pemasukan"`
}
