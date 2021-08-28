#!/usr/bin/env bash

# build_pages.sh

MARKDOWN_TO_HTML="go-markdown-to-html"

function to_html () {
  echo "./${MARKDOWN_TO_HTML} $1 '$2' '$3' '$4' > $4";
  ./${MARKDOWN_TO_HTML} $1 "$2" "$3" "$4" > $4
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
mkdir_dist_path dist/2013/10/11
to_html markdown/2013/10/11/interview-with-a-stranger.markdown \
  'Interview With a Stranger' \
  '' \
  dist/2013/10/11/interview-with-a-stranger.html
mkdir_dist_path dist/2013/10/16
to_html markdown/2013/10/16/today-i-learned-bootstrap-has-built-in-code-blocks.markdown \
  'Today I Learned: Bootstrap Has Built-in Code Blocks' \
  '' \
  dist/2013/10/16/today-i-learned-bootstrap-has-built-in-code-blocks.html
mkdir_dist_path dist/2013/11/15
to_html markdown/2013/11/15/prepare-for-your-interview.markdown \
  'Prepare for your Interview!' \
  '' \
  dist/2013/11/15/prepare-for-your-interview.html
mkdir_dist_path dist/2013/12/02
to_html markdown/2013/12/02/let-a-baby-pick-your-passwords.markdown \
  'Let a Baby Pick Your Passwords' \
  '' \
  dist/2013/12/02/let-a-baby-pick-your-passwords.html
# One-off for media, but TODO to streamline this process
cp markdown/2013/12/02/*.png dist/2013/12/02/
mkdir_dist_path dist/2013/12/09
to_html markdown/2013/12/09/technical-interview-prep-phone-screen-fun.markdown \
  'Technical Interview Prep: Phone Screen Fun' \
  '' \
  dist/2013/12/09/technical-interview-prep-phone-screen-fun.html
mkdir_dist_path dist/2013/12/10
to_html markdown/2013/12/10/how-i-pull-down-my-latest-social-posts-feedmesocial \
  'How I Pull Down My Latest Social Posts' \
  'feedmesocial' \
  dist/2013/12/10/how-i-pull-down-my-latest-social-posts-feedmesocial.html
