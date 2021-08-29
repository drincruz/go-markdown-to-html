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
to_html markdown/2013/12/10/how-i-pull-down-my-latest-social-posts-feedmesocial.markdown \
  'How I Pull Down My Latest Social Posts' \
  'feedmesocial' \
  dist/2013/12/10/how-i-pull-down-my-latest-social-posts-feedmesocial.html
mkdir_dist_path dist/2013/12/17
to_html markdown/2013/12/17/tries-in-python.markdown \
  'Tries in Python' \
  'trie as in Retrieve' \
  dist/2013/12/17/tries-in-python.html
mkdir_dist_path dist/2014/01/08
to_html markdown/2014/01/08/reading-xml-in-python-with-elementtree.markdown \
  'Reading XML in Python with ElementTree' \
  '' \
  dist/2014/01/08/reading-xml-in-python-with-elementtree.html
mkdir_dist_path dist/2014/02/01
to_html markdown/2014/02/01/from-initial-commit-to-launch-the-story-of-a-quick-site-launch.markdown \
  'From Initial Commit to Launch' \
  'The Story of a Quick Site Launch' \
  dist/2014/02/01/from-initial-commit-to-launch-the-story-of-a-quick-site-launch.html
mkdir_dist_path dist/2014/04/07
to_html markdown/2014/04/07/silly-bug-fixes.markdown \
  'Silly Bug Fixes' \
  '' \
  dist/2014/04/07/silly-bug-fixes.html
# One-off for media, but TODO to streamline this process
cp markdown/2014/04/07/*.png dist/2014/04/07/

mkdir_dist_path dist/2014/05/07
to_html markdown/2014/05/07/py-closure-a-python-client-for-closure.markdown \
  'Py-Closure' \
  'A Python Client for Closure' \
  dist/2014/05/07/py-closure-a-python-client-for-closure.html

mkdir_dist_path dist/2014/07/18
to_html markdown/2014/07/18/stand-ups-and-bikesheds.markdown \
  'Stand-ups and Bikesheds' \
  '' \
  dist/2014/07/18/stand-ups-and-bikesheds.html

mkdir_dist_path dist/2014/08/06
to_html markdown/2014/08/06/julython-2014-recap.markdown \
  'Julython 2014 Recap' \
  '' \
  dist/2014/08/06/julython-2014-recap.html
# One-off for media, but TODO to streamline this process
cp markdown/2014/08/06/*.png dist/2014/08/06/

mkdir_dist_path dist/2014/11/03
to_html markdown/2014/11/03/fighting-procrastination-with-time-management.markdown \
  'Fighting Procrastination With Time Management' \
  '' \
  dist/2014/11/03/fighting-procrastination-with-time-management.html

mkdir_dist_path dist/2015/04/06
to_html markdown/2015/04/06/squashing-git-commits-is-great-code-etiquette.markdown \
  'Squashing git Commits is Great Code is Great Code Etiquette' \
  '' \
  dist/2015/04/06/squashing-git-commits-is-great-code-etiquette.html
