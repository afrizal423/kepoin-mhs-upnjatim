package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/afrizal423/kepoin-mhs-upnjatim/kepoin"
)

func connected() (ok bool) {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return false
	}
	return true
}
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")()
	if !connected() {
		panic("Kamu OFFLINE, belum terhubung ke internet!")
	}
	f := kepoin.CariMahasiswa{}
	npm := os.Args[1]
	f.FindMhsByNPM(npm)
	fmt.Println()
}
