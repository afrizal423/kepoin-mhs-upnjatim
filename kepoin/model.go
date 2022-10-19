package kepoin

import "time"

type FindInPDDikti struct {
	Mahasiswa []struct {
		Text        string `json:"text"`
		WebsiteLink string `json:"website-link"`
	} `json:"mahasiswa"`
}

type FoundOnPDDikti struct {
	Datastatuskuliah []struct {
		IDSmt     string `json:"id_smt"`
		SksSmt    int    `json:"sks_smt"`
		NmStatMhs string `json:"nm_stat_mhs"`
	} `json:"datastatuskuliah"`
	Datastudi []struct {
		KodeMk     string      `json:"kode_mk"`
		NmMk       string      `json:"nm_mk"`
		SksMk      int         `json:"sks_mk"`
		IDSmt      string      `json:"id_smt"`
		NilaiHuruf interface{} `json:"nilai_huruf"`
	} `json:"datastudi"`
	Dataumum struct {
		NmPd         string      `json:"nm_pd"`
		Jk           string      `json:"jk"`
		Nipd         string      `json:"nipd"`
		Namapt       string      `json:"namapt"`
		Namajenjang  string      `json:"namajenjang"`
		Namaprodi    string      `json:"namaprodi"`
		RegPd        string      `json:"reg_pd"`
		MulaiSmt     string      `json:"mulai_smt"`
		NmJnsDaftar  string      `json:"nm_jns_daftar"`
		NmPtAsal     interface{} `json:"nm_pt_asal"`
		NmProdiAsal  interface{} `json:"nm_prodi_asal"`
		KetKeluar    string      `json:"ket_keluar"`
		TglKeluar    time.Time   `json:"tgl_keluar"`
		NoSeriIjazah string      `json:"no_seri_ijazah"`
		SertProf     interface{} `json:"sert_prof"`
		LinkPt       string      `json:"link_pt"`
		LinkProdi    string      `json:"link_prodi"`
	} `json:"dataumum"`
}
