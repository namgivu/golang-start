package main

import (
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
)

/*ref. https://jonathanmh.com/web-scraping-golang-goquery/*/

func postScrape() {
	//doc, err := goquery.NewDocument("http://jonathanmh.com")
	doc, err := goquery.NewDocument("https://jonathanmh.com")
	if err != nil {
		log.Fatal(err)
	}

	// use CSS selector found with the browser inspector
	// for each, use index and item
	selector := "#main article .entry-title"
	doc.Find(selector).Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		fmt.Printf("Post #%d: %s - %s\n", index, title, link)
	})
}

func postScrape02() {
	rollOutTypes := []string{"staging", "live"}
	ventures := []string{"ID", "MY", "PH", "SG", "TH", "VN"}

	for _, rollOutType := range rollOutTypes {
		for _, venture := range ventures {
			soaPage := fmt.Sprintf("https://soa-manager.lzd.co/#/%s/etcd/%s/instances/", rollOutType, venture)
			fmt.Printf("%s %s %s\n", rollOutType, venture, soaPage)

			/*region scape the page*/
			doc, err := goquery.NewDocument(soaPage)
			if err != nil { log.Fatal(err) }
			//selector := "body > div > div > ng-include > loader > div:nth-child(1) > ng-transclude > div > table > tbody > tr:nth-child(2) > td.ng-binding > h4"
			//selector := "table[data-qa-locator=etcd_nav_intances_browse_v2_link] tr:nth-child(2) h4"
			//selector := "table > tbody > tr:nth-child(2) > td.ng-binding > h4"
			selector := "table"
			doc.Find(selector).Each(func(index int, item *goquery.Selection) {
				instanceName := item.Text()
				//linkTag := item.Find("a")
				//link, _ := linkTag.Attr("href")
				//fmt.Printf("Post #%d: %s - %s\n", index, title, link)
				fmt.Printf("%q", item)
				fmt.Printf("%s\n", instanceName)
			})
			/*endregion scape the page*/
		}
	}
}

func main() {
	postScrape()
	//postScrape02()
}