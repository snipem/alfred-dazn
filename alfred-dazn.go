package main

// run: make
// run: alfred_workflow_data=workflow alfred_workflow_cache=/tmp/alfred alfred_workflow_bundleid=mk_testing go run alfred-dazn.go

import (
	aw "github.com/deanishe/awgo"
	"github.com/gocolly/colly"
	"strconv"
	"time"
)

// aw.Workflow is the main API
var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func main() {
	wf.Run(getDAZNSchedule)
}

func getDAZNSchedule() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.spox.com"),
	)

	// Extract product details
	c.OnHTML("div.formatdate", func(e *colly.HTMLElement) {

		dataID := e.Attr("data-id")
		fixture := e.ChildText(".t")
		competition := e.ChildText(".dz")
		dateTime, _ := strconv.ParseInt(e.Attr("data-date"), 10, 64)
		startTime := time.Unix(dateTime/1000, 0)
		url := "https://www.dazn.com/de-DE/home/" + dataID

		isLive := e.DOM.Find(".live").Length() > 0

		if isLive {
			wf.NewItem(fixture + " (" + competition + ")").
				Subtitle(startTime.String()).
				Valid(true).
				Arg(url).
				Quicklook(url).
				UID(dataID)
		}
	})

	c.Visit("http://www.spox.com/daznpic/daznprogram.html?c=spoxschedule")
	wf.SendFeedback()
}
