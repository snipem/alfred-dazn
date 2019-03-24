package main

// run: alfred_workflow_data=workflow alfred_workflow_cache=/tmp/alfred alfred_workflow_bundleid=mk_testing go run alfred-dazn.go

import (
	"log"

	aw "github.com/deanishe/awgo"
	"github.com/gocolly/colly"
)

// aw.Workflow is the main API
var wf *aw.Workflow

func init() {
	// Create a new *Workflow using default configuration
	// (workflow settings are read from the environment variables
	// set by Alfred)
	wf = aw.New()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}

func getDAZNSchedule() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Allow requests only to store.xkcd.com
		colly.AllowedDomains("www.spox.com"),
	)

	// Extract product details
	c.OnHTML("div.formatdate", func(e *colly.HTMLElement) {

		// whole := e.Attr("date-search")
		// whole := e.Text
		dataID := e.Attr("data-id")
		fixture := e.ChildText(".t")
		competition := e.ChildText(".dz")
		time := e.ChildText(".time")
		url := "https://www.dazn.com/de-DE/home/" + dataID

		wf.NewItem(fixture + " (" + competition + ") " + time).
			// Subtitle(album.Album.Title).
			Valid(true).
			// Icon(&icon).
			Arg(url).
			Quicklook(url).
			UID(dataID)
	})

	c.Visit("http://www.spox.com/daznpic/daznprogram.html?c=spoxschedule")
	// Display collector's statistics
	log.Println(c)
	wf.SendFeedback()
}

func run() {
	getDAZNSchedule()
}

func outputWorkflowData() {

	// for _, album := range albums.Data {
	// 	var icon aw.Icon
	// 	icon.Value = album.CoverSmall

	// 	id := strconv.Itoa(album.ID)
	// 	url := "https://www.deezer.com/en/album/" + id

	// 	wf.NewItem(album.Artist.Name + " - " + album.Title).
	// 		// Subtitle(album.Album.Title).
	// 		Valid(true).
	// 		// Icon(&icon).
	// 		Arg(url).
	// 		Quicklook(url).
	// 		UID("album" + id).
	// 		NewModifier("cmd").
	// 		Subtitle("Open in Deezer App").
	// 		Arg(getLocalURL(url))
	// }

	// // And send the results to Alfred
	// wf.SendFeedback()
}
