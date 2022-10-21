<p align="right">
بِسْــــــــــــــمِ اللَّهِ الرَّحْمَنِ الرَّحِيم 
</p>

# Kepoin Matkul Mahasiswa Yupien JaTim
Program untuk mengetahui mahasiswa mengambil mata kuliah reguler, <s>mata kuliah MKDU <b>(soon)</b>, dan Mata kuliah Merdeka Belajar Permata Out <b>(soon)</b></s>.

## How to use
- Silahkan masuk ke direktori [build](/build)
- Pilih sesuai sistem operasi linux atau windows
- Cara menjalankan program di Linux
	- download [filenya](https://github.com/afrizal423/kepoin-mhs-upnjatim/raw/master/build/linux/kepoin-mhs-upnjatim) terlebih dahulu
	- masuk ke folder file yang sudah didownload tadi
	- buka terminal, ketikkan
		```sh
		./kepoin-mhs-upnjatim <npm>
		```
		contoh
		```sh
		./kepoin-mhs-upnjatim 180810100xx
		```
- Cara menjalankan program di WIndows
	- download [filenya](https://github.com/afrizal423/kepoin-mhs-upnjatim/raw/master/build/windows/kepoin-mhs-upnjatim.exe) terlebih dahulu
	- masuk ke folder file yang sudah didownload tadi
	- buka command prompt, ketikkan
		```sh
		kepoin-mhs-upnjatim.exe <npm>
		```
		contoh
		```sh
		kepoin-mhs-upnjatim.exe 180810100xx
		```
## Develop
Silahkan dipake sesuai yang diinginkan

### Install
```
go get github.com/afrizal423/kepoin-mhs-upnjatim/kepoin
```

### Example

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
Proses akan membutuhkan waktu karena menjelajahi setiap mata kuliah yang tersedia.
Silahkan jika ingin mengembangkannya lagi.
Mata kuliah MKDU, dan Mata kuliah Merdeka Belajar Permata Out tidak dapat dipastikan <i>bisa</i>, dikarenakan server yupien jatim lemot :hand_over_mouth:.