#!/usr/bin/env bash

# build_pages.sh

MARKDOWN_TO_HTML="go-markdown-to-html"

function to_html () {
  echo "./${MARKDOWN_TO_HTML} $1 '$2' '$3' '$4'";
  ./${MARKDOWN_TO_HTML} $1 "$2" "$3" "$4"
}

function mkdir_dist_path () {
  echo "mkdir -p $1";
  mkdir -p $1
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
  'Squashing git Commits is Great Code Etiquette' \
  '' \
  dist/2015/04/06/squashing-git-commits-is-great-code-etiquette.html

mkdir_dist_path dist/2015/05/31
to_html markdown/2015/05/31/intro-to-building-out-data-pipelines-with-python-and-luigi.markdown \
  'Intro to Building Out Data Pipelines' \
  'With Python and Luigi' \
  dist/2015/05/31/intro-to-building-out-data-pipelines-with-python-and-luigi.html

mkdir_dist_path dist/2016/03/21
to_html markdown/2016/03/21/writing-out-files-python-unicodeencodeerror-woes.markdown \
  'Writing Out Files & Python UnicodeEncodeError Woes' \
  '' \
  dist/2016/03/21/writing-out-files-python-unicodeencodeerror-woes.html

mkdir_dist_path dist/2016/07/31
to_html markdown/2016/07/31/always-a-student-always-learning.markdown \
  'Always a Student, Always Learning' \
  '' \
  dist/2016/07/31/always-a-student-always-learning.html

mkdir_dist_path dist/2017/03/14
to_html markdown/2017/03/14/elixir-conditionals-with-function-signatures.markdown \
  'Elixir Conditionals with Function Signatures' \
  '' \
  dist/2017/03/14/elixir-conditionals-with-function-signatures.html

mkdir_dist_path dist/2017/08/19
to_html markdown/2017/08/19/what-does-fibonacci-look-like-in-elixir.markdown \
  'What Does Fibonacci Look Like in Elixir?' \
  '' \
  dist/2017/08/19/what-does-fibonacci-look-like-in-elixir.html

mkdir_dist_path dist/2018/08/04
to_html markdown/2018/08/04/three-years-of-searching.markdown \
  'Three Years of Searching' \
  '' \
  dist/2018/08/04/three-years-of-searching.html

mkdir_dist_path dist/2019/03/30
to_html markdown/2019/03/30/distance-running-and-software-engineering.markdown \
  'Distance Running and Software Engineering' \
  '' \
  dist/2019/03/30/distance-running-and-software-engineering.html
# One-off for media, but TODO to streamline this process
cp markdown/2019/03/30/*.png dist/2019/03/30/

mkdir_dist_path dist/2019/05/27
to_html markdown/2019/05/27/soft-skills-for-software-engineers.markdown \
  'Soft Skills for Software Engineers' \
  '' \
  dist/2019/05/27/soft-skills-for-software-engineers.html

mkdir_dist_path dist/2021/09/17
to_html markdown/2021/09/17/optimizing-for-writing.markdown \
  'Optimizing for Writing' \
  'Writing content, that is!' \
  dist/2021/09/17/optimizing-for-writing.html

mkdir_dist_path dist/2022/10/23
to_html markdown/2022/10/23/weekly-writing-prompts.markdown \
  'Weekly Writing Prompts' \
  'A Journal for Self-reflection' \
  dist/2022/10/23/weekly-writing-prompts.html

mkdir_dist_path dist/2022/12/21
to_html markdown/2022/12/21/asking-for-help.markdown \
  'Asking for Help' \
  'Everyone Needs a Little Help Now and Then' \
  dist/2022/12/21/asking-for-help.html
cp markdown/2022/12/21/*.jpg dist/2022/12/21/

mkdir_dist_path dist/2023/08/27
to_html markdown/2023/08/27/webflow-a-year-in-review.markdown \
  'Webflow: A Year in Review' \
  'Looking back at my first year of learnings' \
  dist/2023/08/27/webflow-a-year-in-review.html
cp markdown/2023/08/27/*.jpg dist/2023/08/27/

mkdir_dist_path dist/2023/10/21
to_html markdown/2023/10/21/engineering-guilds-at-work.markdown \
  'Engineering Guilds at Work' \
  'A look back at what worked well and what did not work quite well' \
  dist/2023/10/21/engineering-guilds-at-work.html
cp markdown/2023/10/21/*.jpg dist/2023/10/21/

mkdir_dist_path dist/2023/12/10
to_html markdown/2023/12/10/opportunities-for-growth-in-engineering.markdown \
  'Opportunities for Growth in Engineering' \
  'Why opportunities are important for engineering' \
  dist/2023/12/10/opportunities-for-growth-in-engineering.html
cp markdown/2023/12/10/*.jpg dist/2023/12/10/

# Write the yearly archive pages
to_html write_year_archives
to_html markdown/archive.markdown 'Archive' 'Posts from the past' dist/archive.html
# Write the Index page last
to_html write_index
