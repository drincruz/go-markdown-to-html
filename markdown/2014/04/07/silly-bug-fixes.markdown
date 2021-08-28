![Silly SQL bug](Screenshot-04072014_011622.png)

Yes, this was a bug fix. "A one line bug?" you ask? Yes, sad but true.

## Now, for some background...

A bug was sent in about pagination not completely working for a web application. It was consistently showing only four pages. This is obviously not expected behavior especially if the user knows that there should have been 10, 20, 50 pages.

## Okay, now the history of that code

So it just so happens that the portion of the site was quickly rolled out to return only 100 rows from the database as a shortcut. As part of technical debt, a fix for pagination was then added. Unfortunately, sometimes working in a fast pace, you miss some details in code and bugs like this slip through.

Doh!

## Silly, but let's fix this

So, I know I'm not the only engineer to have commits like this. I mean, remember the `goto fail;` [bug](https://www.imperialviolet.org/2014/02/22/applebug.html)? To remedy this, we've already started doing code review on this project. Why haven't we been doing code review from the beginning? Well, if you remember my post on a [quick site launch](/2014/02/01/from-initial-commit-to-launch-the-story-of-a-quick-site-launch.html), I'm sure you are aware that this is for a site that was bootstrap'd fast to ship a functional product in a short amount of time.

Unfortunately, things like this was expected due to the lack of time. But here's to finding a humorous tone in my mistakes. Cheers!
