/* alert_scraper - parses alerts from omnilert.net and writes to csv
* Written by: Brennan Turner @BLTSEC
* Website: https://bltsec.ninja
* GIT: https://github.com/BLTSEC
 */
package main

import (
	"flag"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// Scrape function that requests alerts then parses them into a csv file
func Scrape() {
	urlPtr := flag.String("url", "https://widgets.omnilert.net/-11273", "URL to receive alerts from")
	dstPtr := flag.String("dst", ".", "Destination to save the csv file to")
	flag.Parse()

	doc, err := goquery.NewDocument(*urlPtr)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(*dstPtr + "/alerts.csv")

	defer f.Close()

	// Find the alerts
	doc.Find(".SmartBoard_Table").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the subject, datetime, and message
		subject := s.Find(".SmartBoard_Subject").Text()
		datetime := s.Find(".SmartBoard_DateTime").Text()
		message := s.Find(".SmartBoard_Message").Text()
		// Write to csv file
		f.WriteString(subject + ", " + datetime + ", " + message + "\n")
	})
}

func main() {
	Scrape()
}
