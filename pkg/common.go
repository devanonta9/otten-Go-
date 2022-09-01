package pkg

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func LookUpHTML(raw io.Reader) (tableData []string, err error) {
	doc, err := goquery.NewDocumentFromReader(raw)
	if err != nil {
		return tableData, err
	}

	doc.Find(".table_style").Children().Each(func(i int, sel *goquery.Selection) {
		rows := sel.Find("tr").Text()
		tableData = append(tableData, rows)
	})

	tableData, err = MatchHTML(tableData)
	if err != nil {
		return tableData, err
	}

	return
}

func MatchHTML(data []string) (result []string, err error) {
	tableRow := strings.SplitN(data[5], "\n", -1)

	for _, val := range tableRow {
		if len(val) > 0 {
			result = append(result, val)
		}
	}

	return
}
