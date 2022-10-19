package kepoin

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type ScrapeMahasiswa struct {
	Npm string
}

func (sc *ScrapeMahasiswa) ScrapeSiamik(kodeprodi string) {
	sc.scrapeListProdi(kodeprodi)
}

func (sc *ScrapeMahasiswa) scrapeListPesertaMatkul(nextUrl string) {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("siamik.upnjatim.ac.id"),
	)

	c.OnHTML("h7 p", func(e *colly.HTMLElement) {
		if e.Index == 1 {
			split := strings.Split(e.Text, "Mata Kuliah : ")
			fmt.Println("Mata Kuliah :", split[1])
		}
		// fmt.Printf("Matkul found: %s\n", e.Index)
	})

	// On every a element which has href attribute call callback
	c.OnHTML("table tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			if row.Index != 0 && row.Index != 1 {

				row.ForEach("td", func(_ int, td *colly.HTMLElement) {

					if td.Index == 1 && td.Text == sc.Npm {
						fmt.Println(td.Text, td.Index, sc.Npm, "KKKKKKKKKKKKKKKKKKKKKKKKEEEEEEEEEEEEEEEETEEEEEEEEEEEEEEMUUUUUUUUUUUUUU")
					}
					// fmt.Println(td.Text, td.Index, sc.Npm)
				})
			}

		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println()
		// fmt.Println("Visiting ListPesertaMatkul", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(nextUrl)
}

func (sc *ScrapeMahasiswa) scrapeListMatkul(nextUrl string) {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("siamik.upnjatim.ac.id"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("table tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {

			if row.Index != 0 {
				row.ForEach("td", func(_ int, td *colly.HTMLElement) {
					// fmt.Println(td.Text, td.Index)
					if td.Index == 1 {
						row.ForEach("a[href]", func(_ int, d *colly.HTMLElement) {
							link := d.Attr("href")
							// Print link
							// fmt.Printf("Link found: %q -> %s\n", d.Text, link)
							// fmt.Println(e.Request.AbsoluteURL(link))

							sc.scrapeListPesertaMatkul(e.Request.AbsoluteURL(link))
						})
					}

					if td.Index == 4 {
						fmt.Println("Kelas : ", td.Text)
					}

				})
			}

		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println()
		fmt.Println("Visiting Matkul", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(nextUrl)
}

func (sc *ScrapeMahasiswa) scrapeListProdi(kodprod string) {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("siamik.upnjatim.ac.id"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("table tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {

			if row.Index != 0 {
				row.ForEach("td", func(_ int, td *colly.HTMLElement) {
					if td.Index == 2 && td.Text == kodprod {
						fmt.Println(td.Text, td.Index)
						row.ForEach("a[href]", func(_ int, d *colly.HTMLElement) {
							link := d.Attr("href")
							// Print link
							// fmt.Printf("Link found: %q -> %s\n", d.Text, link)
							// fmt.Println(e.Request.AbsoluteURL(link))

							sc.scrapeListMatkul(e.Request.AbsoluteURL(link))
						})
					}
				})
			}
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://siamik.upnjatim.ac.id/html/siamik/daftarPesertaKuliah.asp")
}
