package main

import (
	"fmt"
	"log"
	"os"
)

// Pages that are published but should not be crawled.
var robotsDisallowed = []string{
	"/test.html",
	"/error.html",
}

func buildRobots() string {
	robots := "User-agent: *\nAllow: /\n"
	for _, path := range robotsDisallowed {
		robots += fmt.Sprintf("Disallow: %s\n", path)
	}
	robots += fmt.Sprintf("\nSitemap: %s/sitemap.xml\n", BaseURL())
	return robots
}

func writeRobots(outPath string) {
	check(os.WriteFile(outPath, []byte(buildRobots()), 0o644), "Failed to write robots.txt")
	log.Printf("[INFO][writeRobots] Wrote %s", outPath)
}
