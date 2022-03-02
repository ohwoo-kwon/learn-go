package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id string
	title string
	location string
	salary string
	summary string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"
var LastpageURL string = "https://kr.indeed.com/jobs?q=python&limit=50&start=9999"

func main() {
	totalPages := getPages()
	
	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")

	searchCards.Each(func(i int, card *goquery.Selection) {
		id, _ := card.Attr("data-jk")
		title := card.Find("h2 > span").Text()
		location := card.Find(".companyLocation").Text()
		salary := card.Find(".salary-snippet > span").Text()
		summary :=  card.Find(".job-snippet").Text()
		fmt.Println(id)
		fmt.Println(title)
		fmt.Println(location)
		fmt.Println(salary)
		fmt.Println(summary)
	})
}

func getPages() int {
	pages := 0
	res, err := http.Get(LastpageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages, _ = strconv.Atoi(s.Find("b").Text())
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}