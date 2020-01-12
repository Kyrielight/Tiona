package main

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
)

var mangadexURL string = "https://mangadex.cc"

func rss(w http.ResponseWriter, req *http.Request) {

	resp, _ := http.Get(mangadexURL + req.URL.String())

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	//feedReg := regexp.MustCompile(`(\t<mangaLink>.+<\/mangaLink>\r\n)`)
	//feed := feedReg.ReplaceAllString(buf.String(), "")

	feedReg := regexp.MustCompile(`<mangaLink>\D+(\d+)<\/mangaLink>`)
	image := `<image><title>Cover</title><link>https://mangadex.cc</link><url>https://mangadex.cc/images/manga/$1.jpg</url></image>`
	feed := feedReg.ReplaceAllString(buf.String(), image)
	fmt.Fprintf(w, feed)
}

func main() {

	http.HandleFunc("/", rss)
	http.ListenAndServe(":8080", nil)

}
