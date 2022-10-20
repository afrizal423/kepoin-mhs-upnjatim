package main

import (
	"fmt"

	"github.com/afrizal423/kepoin-mhs-upnjatim/kepoin"
)

func main() {
	f := kepoin.CariMahasiswa{}
	f.FindMhsByNPM("22011010106")
	fmt.Println()
}
