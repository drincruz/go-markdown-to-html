package main

import (
	"strings"
	"testing"
)

func TestBuildRobots(t *testing.T) {
	robots := buildRobots()
	wantLines := []string{
		"User-agent: *",
		"Allow: /",
		"Disallow: /error.html",
		"Sitemap: https://www.drincruz.com/sitemap.xml",
	}
	for _, line := range wantLines {
		if !strings.Contains(robots, line+"\n") {
			t.Errorf("robots.txt is missing %q:\n%s", line, robots)
		}
	}
	if strings.Contains(robots, "/test.html") {
		t.Errorf("robots.txt should not mention the removed test page:\n%s", robots)
	}
	if !strings.HasPrefix(robots, "User-agent: *\n") {
		t.Errorf("robots.txt should start with the User-agent line:\n%s", robots)
	}
}
