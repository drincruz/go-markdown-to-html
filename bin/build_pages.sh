#!/usr/bin/env bash

# build_pages.sh

MARKDOWN_TO_HTML="go-markdown-to-html"

function to_html () {
  echo "./${MARKDOWN_TO_HTML} $@"
  ./${MARKDOWN_TO_HTML} "$@"
}

# One-off for favicon
cp markdown/android-chrome-192x192.png dist/
cp markdown/android-chrome-512x512.png dist/
cp markdown/apple-touch-icon.png dist/
cp markdown/favicon-16x16.png dist/
cp markdown/favicon-32x32.png dist/
cp markdown/favicon.ico dist/
cp markdown/site.webmanifest dist/

to_html markdown/test.markdown 'This is a title' 'And a subtitle' dist/test.html
to_html markdown/error.markdown 'Error' 'Uh-oh, something went wrong' dist/error.html
to_html markdown/about.markdown 'About Me' 'So who am I?' dist/about.html
# Write every post listed in the CCYY.json files
to_html write_posts

# Write the yearly archive pages
to_html write_year_archives
# Write robots.txt and sitemap.xml
to_html write_seo_files
to_html markdown/archive.markdown 'Archive' 'Posts from the past' dist/archive.html
# Write the Index page last
to_html write_index
