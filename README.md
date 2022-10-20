<p align="right">
بِسْــــــــــــــمِ اللَّهِ الرَّحْمَنِ الرَّحِيم 
</p>

# Kepoin Matkul Mahasiswa Yupien JaTim
Program untuk mengetahui mahasiswa mengambil mata kuliah reguler, mata kuliah MKDU <b>(soon)</b>, dan Mata kuliah Merdeka Belajar Permata Out <b>(soon)</b>.

## Install
```
go get github.com/afrizal423/kepoin-mhs-upnjatim/kepoin
```

## Example

```go
package main

import (
	"fmt"

	"github.com/afrizal423/kepoin-mhs-upnjatim/kepoin"
)

func main() {
	f := kepoin.CariMahasiswa{}
	f.FindMhsByNPM("xxxxxxxxx")
}
```

### Jalankan
```sh
go run main.go
```
### Output
Hasil keluaran <b>hanya sebatas teks console</b>.
```
======================Biodata========================
Nama Mahasiswa : xyz
NPM : xyz
Prodi: Informatika
=====================================================
=======================Matkul========================
Kode Matkul : IF221123
Mata Kuliah : RISET INFORMATIKA
Jumlah SKS: 3
Kelas :  C
=====================================================
=======================Matkul========================
Kode Matkul : IF221124
Mata Kuliah : MANAJEMEN PROYEK
Jumlah SKS: 3
Kelas :  A
=====================================================

```

## Catatan
Proses akan membutuhkan waktu karena untuk menjelajahi setiap mata kuliah yang tersedia.
Silahkan jika ingin mengembangkannya lagi.