#!/usr/bin/env bash

# build_pages.sh

MARKDOWN_TO_HTML="go-markdown-to-html"

function to_html () {
  echo "./${MARKDOWN_TO_HTML} $1 '$2' '$3' > $4";
  ./${MARKDOWN_TO_HTML} $1 "$2" "$3" > $4
}

function mkdir_dist_path () {
  echo "mkdir -p $1";
  mkdir -p $1
}

to_html markdown/test.markdown 'This is a title' 'And a subtitle' dist/test.html
to_html markdown/index.markdown 'Greetings!' 'My name is Adrian' dist/index.html
to_html markdown/about.markdown 'About Me' 'So who am I?' dist/about.html
mkdir_dist_path dist/2013/10/10
to_html markdown/2013/10/10/yes-this-is-bootstrap.markdown \
  'Yes, This is Bootstrap' \
  'Updated, Bootstrap (still)' \
  dist/2013/10/10/yes-this-is-bootstrap.html
