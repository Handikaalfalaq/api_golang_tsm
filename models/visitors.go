package models

type Visitors struct {
	ID          int    `json:"id"`
	Nama        string `json:"nama"`
	NoKTP       string `json:"no_ktp"`
	TglLahir    string `json:"tgl_lahir"`
	NamaGedung  string `json:"nama_gedung"`
	Suhu        string `json:"suhu"`
	JamMasuk    string `json:"jam_masuk"`
	JamKeluar   string `json:"jam_keluar"`
	PhoneNumber string `json:"phone_number"`
}
