package main

// run: alfred_workflow_data=workflow alfred_workflow_cache=/tmp/alfred alfred_workflow_bundleid=mk_testing go run alfred-dazn.go

import (
	aw "github.com/deanishe/awgo"
	"github.com/gocolly/colly"
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
		time := e.ChildText(".time")
		url := "https://www.dazn.com/de-DE/home/" + dataID

		wf.NewItem(fixture + " (" + competition + ") " + time).
			Valid(true).
			Arg(url).
			Quicklook(url).
			UID(dataID)
	})

	c.Visit("http://www.spox.com/daznpic/daznprogram.html?c=spoxschedule")
	wf.SendFeedback()
}
