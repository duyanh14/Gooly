package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"io"
	"net/http"
)

func main() {
	req, _ := http.Get("http://example.com/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36\n")

	body, _ := io.ReadAll(req.Body)

	doc := soup.HTMLParse(string(body))
	links := doc.Find("title").Text()
	fmt.Println(links)
	//fmt.Println(string(body))
}