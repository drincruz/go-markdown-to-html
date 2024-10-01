package main

import "log"

func getMostRecentArray(posts []Blog) []string {
	const LOG_PREFIX = "[getMostRecentArray]"
	if len(posts) < 1 {
		log.Printf("%s No posts found\n", LOG_PREFIX)
		return []string{}
	}
	mostRecent := []string{}
	for _, post := range posts {
		mostRecent = append(mostRecent, getMarkdownFilenameFromHtml(post.Url))
	}
	log.Printf("%s Most recent: %+v\n", LOG_PREFIX, mostRecent)

	return mostRecent
}
