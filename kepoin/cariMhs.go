package kepoin

type CariMahasiswa struct {
	Npm   string
	Nama  string
	Prodi string
}

func (findMhs *CariMahasiswa) FindMhsByNPM(npm string) {
	findMhs.Npm = npm
	f := PdDikti{}
	foundDikti := f.FindMhsByNPM(npm)
	findMhs.Nama = foundDikti.Dataumum.NmPd
	findMhs.Prodi = foundDikti.Dataumum.Namaprodi
	sc := ScrapeMahasiswa{
		Npm: npm,
	}
	sc.ScrapeSiamik(findMhs.Npm[2:5])
}
