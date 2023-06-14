package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Entry struct {
	AuthorID string
	Author   string
	TitleID  string
	Title    string
	InfoURL  string
	ZipURL   string
}

func findEntries(siteURL string) ([]Entry, error) {
	doc, err := goquery.NewDocument(siteURL)
	if err != nil {
		return nil, err
	}

	doc.Find("ol li a").Each(func(n int, elem *goquery.Selection) {
		println(elem.Text(), elem.AttrOr("href", ""))
	})

	return nil, nil // とりま、nil
}

func main() {
	// 芥川龍之介の作品一覧ページ
	listURL := "https://www.aozora.gr.jp/index_pages/person879.html"

	entires, err := findEntries(listURL)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entires {
		fmt.Println(entry.Title, entry.ZipURL)
	}
}
