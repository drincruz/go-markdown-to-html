#!/usr/bin/env bash

# build_pages.sh

MARKDOWN_TO_HTML="go-markdown-to-html"

function to_html () {
  echo "./${MARKDOWN_TO_HTML} $1 '$2' '$3' > $4";
  ./${MARKDOWN_TO_HTML} $1 "$2" "$3" > $4
}

to_html markdown/test.markdown 'This is a title' 'And a subtitle' dist/test.html
to_html markdown/index.markdown 'Greetings!' 'My name is Adrian' dist/index.html
to_html markdown/about.markdown 'About Me' 'So who am I?' dist/about.html
