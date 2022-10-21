package kepoin

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

type ScrapeMahasiswa struct {
	Npm   string
	Kelas string
}

func (sc *ScrapeMahasiswa) ScrapeSiamik(kodeprodi string) {
	// sc.scrapeListProdi(kodeprodi)
	sc.scrapeListMKDU() // WIP
}

func (sc *ScrapeMahasiswa) scrapeListPesertaMatkul(nextUrl string, kelasnya string, kodematkul string,
	matkul string, sks string) {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("siamik.upnjatim.ac.id"),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5})
	// On every a element which has href attribute call callback
	c.OnHTML("h7 > table > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {

			row.ForEach("td", func(_ int, td *colly.HTMLElement) {

				if td.Index == 1 && td.Text == sc.Npm {
					sksnya, _ := strconv.Atoi(sks)
					split2 := strings.Split(kelasnya, sc.Npm[2:5])
					fmt.Println("=======================Matkul========================")
					fmt.Println("Kode Matkul :", kodematkul)
					fmt.Println("Mata Kuliah :", matkul)
					fmt.Println("Jumlah SKS:", sksnya)
					fmt.Println("Kelas : ", split2[0])
					fmt.Println("=====================================================")
				}
			})

		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println()
		// fmt.Println("Visiting ListPesertaMatkul", r.URL.String())
	})

	c.Visit(nextUrl)
	c.Wait()
}

func (sc *ScrapeMahasiswa) scrapeListMatkul(nextUrl string) {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("siamik.upnjatim.ac.id"),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5})

	// On every a element which has href attribute call callback
	c.OnHTML("table tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {

			if row.Index != 0 {
				kelasnya := row.ChildText("td:nth-child(5)")
				kodematkul := row.ChildText("td:nth-child(2)")
				matkul := row.ChildText("td:nth-child(3)")
				sks := row.ChildText("td:nth-child(4)")
				row.ForEach("td", func(_ int, td *colly.HTMLElement) {

					if td.Index == 1 {
						row.ForEach("a[href]", func(_ int, d *colly.HTMLElement) {
							link := d.Attr("href")
							sc.scrapeListPesertaMatkul(e.Request.AbsoluteURL(link), kelasnya, kodematkul, matkul, sks)
						})
					}
				})
			}

		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println()
		// fmt.Println("Visiting Matkul", r.URL.String())
	})

	c.Visit(nextUrl)
	c.Wait()
}

func (sc *ScrapeMahasiswa) scrapeListProdi(kodprod string) {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("siamik.upnjatim.ac.id"),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5})
	// On every a element which has href attribute call callback
	c.OnHTML("table tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {

			if row.Index != 0 {
				row.ForEach("td", func(_ int, td *colly.HTMLElement) {
					if td.Index == 2 && td.Text == kodprod {
						// fmt.Println(td.Text, td.Index)
						row.ForEach("a[href]", func(_ int, d *colly.HTMLElement) {
							link := d.Attr("href")
							sc.scrapeListMatkul(e.Request.AbsoluteURL(link))
						})
					}
				})
			}
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://siamik.upnjatim.ac.id/html/siamik/daftarPesertaKuliah.asp")

	c.Wait()
}

// MKDU
// WIP
func (sc *ScrapeMahasiswa) scrapeListMKDU() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("siamik.upnjatim.ac.id"),
		colly.Debugger(&debug.LogDebugger{}),
		colly.Async(true),
	)

	// Limit the number of threads started by colly to 5
	// when visiting links which domains' matches "*httpbin.*" glob
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*siamik.upnjatim.*",
		Parallelism: 2,
		RandomDelay: 50 * time.Second,
	})

	// On every a element which has href attribute call callback
	c.OnHTML("h7 > table > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {

			if row.Index != 0 {
				row.ForEach("td", func(_ int, td *colly.HTMLElement) {
					// fmt.Println(td.Text, td.Index)
					row.ForEach("a[href]", func(_ int, d *colly.HTMLElement) {
						link := d.Attr("href")
						// Print link
						fmt.Printf("Link found: %q -> %s\n", d.Text, link)
						// fmt.Println(e.Request.AbsoluteURL(link))

						// sc.scrapeListMatkul(e.Request.AbsoluteURL(link))
					})
				})
			}
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://siamik.upnjatim.ac.id/html/siamik/daftarPesertaKuliahMKDU.asp")
	// Wait until threads are finished
	c.Wait()
}
