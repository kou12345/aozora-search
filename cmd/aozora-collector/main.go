package main

import (
	"fmt"
	"log"
	"regexp"

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

	// 詳細ページへのURLを正規表現で抜き出す
	pat := regexp.MustCompile(`.*/cards/([0-9]+)/card([0-9]+).html$`)
	// URL一覧を取得
	doc.Find("ol li a").Each(func(n int, elem *goquery.Selection) {
		token := pat.FindStringSubmatch(elem.AttrOr("href", ""))
		if len(token) != 3 {
			return
		}
		pageURL := fmt.Sprintf("https://www.aozora.gr.jp/cards/%s/card%s.html",
			token[1], token[2])
		println(pageURL)

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
