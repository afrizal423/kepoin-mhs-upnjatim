package main

import (
	"fmt"

	"github.com/afrizal423/kepoin-mhs-upnjatim/kepoin"
)

func main() {
	f := kepoin.CariMahasiswa{}
	f.FindMhsByNPM("19081010026")
	fmt.Println()
	fmt.Println(f.Nama, f.Npm, f.Prodi)
	fmt.Println("heko")
}
