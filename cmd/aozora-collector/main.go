package main

import (
	"fmt"
	"log"
	"net/url"
	"path"
	"regexp"
	"strings"

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
		// fmt.Println(token) // [../cards/001086/card4311.html 001086 4311]
		if len(token) != 3 { // 正規表現にマッチしない場合
			return
		}
		pageURL := fmt.Sprintf("https://www.aozora.gr.jp/cards/%s/card%s.html",
			token[1], token[2])
		author, zipURL := findAuthorAndZIP(pageURL) // 作者とZIPファイルのURLを得る
		println(author, zipURL)

	})
	
	return nil, nil // とりま、nil
}

// 作者とZIPファイルのURLを得る
func findAuthorAndZIP(siteURL string) (string, string) {
	log.Println("query", siteURL)
	doc, err := goquery.NewDocument(siteURL)
	if err != nil {
		return "", ""
	}

	author := doc.Find("table[summary=作家データ] tr:nth-child(1) td:nth-child(2)").Text()

	zipURL := ""
	doc.Find("table.download a").Each(func(n int, elem *goquery.Selection) {
		href := elem.AttrOr("href", "")
		if strings.HasSuffix(href, ".zip") {
			zipURL = href
		}
	})

	if zipURL == "" {
		return author, ""
	}
	// zipURLの始まりが "http://" or "https://" なら true
	if strings.HasPrefix(zipURL, "http://") || strings.HasPrefix(zipURL, "https://") {
		return author, zipURL
	}

	u, err := url.Parse(siteURL)
	if err != nil {
		return author, ""
	}
	u.Path = path.Join(path.Dir(u.Path), zipURL)

	return author, u.String()
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
