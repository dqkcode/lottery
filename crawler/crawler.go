package crawler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// https://www.minhngoc.net.vn/tra-cuu-ket-qua-xo-so-tinh.html?tinh=14&ngay=1&thang=3&nam=2017
func GetResultFromWebsite(numbers []string, chanel string, day string, month string, year string) {
	// url := "https://www.minhngoc.net.vn/tra-cuu-ket-qua-xo-so-tinh.html?tinh=" + "14" + "&ngay=" + day + "&thang=" + month + "&nam=" + year
	url := "https://www.minhngoc.net.vn/tra-cuu-ket-qua-xo-so-tinh.html?tinh=" + "tinh=46&ngay=10&thang=10&nam" + "=2019"

	res, err := http.Get(url)
	// res, err := http.Get("https://www.minhngoc.net.vn/tra-cuu-ket-qua-xo-so-tinh.html?tinh=46&ngay=10&thang=10&nam=2019")
	// res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	// doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
	doc.Find(".box_kqxs_content tbody tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		text := strings.TrimSpace(s.Find("td .giaidb").Text())
		fmt.Printf("i=%v %v ", i, text)
	})
}
