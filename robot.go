package main

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"gooly/models"
	"gooly/services"
	"io"
	"net/http"
)

func main() {
	db := services.Database{}.Conn()

	url := "https://kenh14.vn/"

	title, desc, thumb, icon := get(url)
	fmt.Println(title)
	fmt.Println(desc)
	fmt.Println(thumb)
	fmt.Println(icon)

	site := models.Site{
		URL:         url,
		Title:       title,
		Description: desc,
		Thubnail:    thumb,
		Icon:        icon,
	}

	result := db.Create(&site)

	fmt.Println(site.ID)
	fmt.Println(result.RowsAffected)
}

func get(url string) (string, string, string, string) {
	req, _ := http.Get(url)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36\n")

	body, _ := io.ReadAll(req.Body)

	doc := soup.HTMLParse(string(body))
	title := doc.Find("title").Text()

	descf := doc.Find("meta", "name", "Description")
	var desc string
	if descf.Error != nil {
		descf = doc.Find("meta", "name", "description")
	}
	if descf.Error == nil {
		desc = descf.Attrs()["content"]
	}

	thumbf := doc.Find("meta", "property", "og:image")
	var thumb string
	if thumbf.Error == nil {
		thumb = thumbf.Attrs()["content"]
	}

	iconf := doc.Find("link", "rel", "icon")
	var icon string
	if iconf.Error == nil {
		icon = iconf.Attrs()["href"]
	}

	return title, desc, thumb, icon
}
