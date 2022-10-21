package kepoin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type PdDikti struct {
	Npm  string
	Nama string
}

func (findMhs *PdDikti) FindMhsByNPM(npm string) FoundOnPDDikti {
	url1 := "https://api-frontend.kemdikbud.go.id/hit_mhs/" + npm
	spaceClient := http.Client{
		Timeout: time.Second * 30, // Timeout after 30 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url1, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var foundWebsiteLink FindInPDDikti
	if err := json.Unmarshal(body, &foundWebsiteLink); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	split := strings.Split(foundWebsiteLink.Mahasiswa[0].WebsiteLink, "/")
	// detail mhs
	url2 := "https://api-frontend.kemdikbud.go.id/detail_mhs/" + split[2]

	req2, err := http.NewRequest(http.MethodGet, url2, nil)
	if err != nil {
		log.Fatal(err)
	}

	res2, getErr := spaceClient.Do(req2)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res2.Body != nil {
		defer res2.Body.Close()
	}

	body2, readErr := ioutil.ReadAll(res2.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var mhsDitemukan FoundOnPDDikti
	if err := json.Unmarshal(body2, &mhsDitemukan); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	return mhsDitemukan
}
