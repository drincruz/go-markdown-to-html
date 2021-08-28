Here is a tale of commits starting from the first initial commit, up to the production release commit. Full disclosure: these are **not** complete commit logs, I have tailored out any work-specific text. These are also **not** all of the commits in the log, but only the highlights.

## Initial Commit and Commuter Commits In One Week

So, here's the problem, I've been tasked to build a web application. They need it "A.S.A.P." aka "we need it yesterday". I'm up for the challenge, I do my research, get the requirements, pick my tools, and off we go.

```
commit fc230bbd88e488224145e2dcb8d3522aebbd706d
Author: Adrian Cruz
Date:   Tue Jan 7 11:20:48 2014 -0500

    initial commit
```

It's a Python [Django](https://www.djangoproject.com/) application. My second place choice was to use [Play Framework](https://www.playframework.com/), and honestly, I would have loved to use it, but I would just be learning Scala and the stress to have the application ASAP resonated. I know Python, I've built web applications in Django before. Let's do this.

```
commit e65ef58f6f32c39a7b345f137c899479485cdc04
Author: Adrian Cruz
Date:   Wed Jan 8 09:44:32 2014 -0500

    working on the bus commit
```

I've been having issues with [Thrift](https://thrift.apache.org/). We use Thrift to provide service API calls for cross-functional teams to use. Thrift is not documented very well. Learning how to utilize Thrift to pull data and store data is taking up too much of my time. It's a blocking item, pushing all of my other tasks further back.

```
commit 61954bbfb2738eba93c8b270fd2dadcce4c99207
Author: Adrian Cruz
Date:   Fri Jan 10 03:42:20 2014 -0500

    working on page detail view, not quite working (404'ing)
     - time for sleep

commit d62a212b3d44836258d82fa783f7d9388ee27493
Author: Adrian Cruz
Date:   Thu Jan 9 03:18:19 2014 -0500

    specific pages are loading in a table now.
    -- i should sleep
```

I want to get this app done. I've got some strength in me to pull some late-night coding sessions.

```
commit 926607874dfe5ad62064a990e29a27c7d3087367
Author: Adrian Cruz
Date:   Mon Jan 13 20:08:45 2014 +0000

     - updated settings and models for running on beta box
```

I've been going back and forth with another developer. I have downtime while he is working on fixing things on his end. I set up a beta server to host my application instead of running this on my local machine. My git logs now reflect GMT time since I am not working on my local machine any longer. _Note_: convert the time to be in GMT -5 to see the actual times I've been working.

```
commit bd4ce1a3029fd412a3ae9e4465645a77405d729d
Author: Adrian Cruz
Date:   Fri Jan 17 19:00:01 2014 +0000

    quick lunch commit -- loading data from thrift successfully

commit dd8d1d209a483bfa32e51203969bc9a83c2a7903
Author: Adrian Cruz
Date:   Fri Jan 17 09:07:41 2014 +0000

    falling asleep commit -- working on pulling data in from thrift now
```

Things are getting better. I'm getting the hang of Thrift finally.

```
commit c6671bbc0ad5aafd41f84096fcdacd59c2466366
Author: Adrian  Cruz
Date:   Tue Jan 21 19:30:03 2014 +0000

    v0.1-beta-rc1
     - functioning; saves and pulls data
```

Looks like I have a working beta version. I had to do a lot of AJAX work from this weekend up until January 21.

## Snowstorms May Stop Traffic, But Not Development

```
commit 01695b45c3db9ac521952f6ed9467907e7ad5675
Author: Adrian Cruz
Date:   Tue Jan 21 22:54:33 2014 +0000

    on the bus commute! wow this snow is horrible! i'm _still_ not home yet.
```

Working on the commute to or from the office has become a normal routine. But this time, I really had to since it was the only thing keeping me sane. Total commute time due to snow: **4 hours 20 minutes**.

## On a Weekend Where Staying In Was On the Agenda: Weekend Work

```
commit 721aa0881e8039925825049e8d3ad2907226171b
Author: Adrian Cruz
Date:   Sun Jan 26 05:49:53 2014 +0000

    WORKING LDAP authentication!
     - need to make it pretty and add @login_required decorators everywhere
```

One of the requirements was to be able to snap into our company's LDAP server so users won't need to remember _another_ password. If anything, I could roll out authentication using Django just for the time being. But luckily, just one weekend worth of hacking, I was able to write a custom LDAP authentication module for Django.

```
commit 01318875cb0f0adaebe7d0d03fd6d38ca8a36593
Author: Adrian Cruz
Date:   Mon Jan 27 08:13:32 2014 +0000

    i should go to sleep - fixed data models
```

```
commit a161c8b476c33dc2330d357143d9a2a8c94dbb79
Author: Adrian Cruz
Date:   Tue Jan 28 12:06:23 2014 +0000

    Thrift saving seems to be working properly now... it's 7am...
```

Having Thrift issues certainly did slow down the process.

## Production Live Beta

```
commit 767f9c0796804d3fbca38db0fd0cac2c1bc746b2
Author: Adrian Cruz
Date:   Tue Jan 28 22:11:10 2014 +0000

    production settings.py
```

```
commit 3cbb95e1fa3f7a38e9de8d8d4f392944e5e0e52a
Author: Adrian Cruz
Date:   Wed Jan 29 19:57:58 2014 +0000

    optimized view for pages
     - LEFT JOIN
     - limited results to 100 until pagination
```

We did a live beta to demo to our end users. I like these since I get to hear user feedback. But on the other hand, **I really got to hear user feedback**! A flood of changes were requested and some of them were done. Others, entered in as a feature request for the next iteration.

```
commit 8a10107c9d3b9ce4ea2ed2d4255b2f6113ae5c7c
Author: Adrian Cruz
Date:   Thu Jan 30 23:51:02 2014 +0000

    publishing properly
     - TODO carousel json
```

I now started working on publishing output. I can see the checkered flag.

## Production: Rollout!

```
commit b8ab872dbcb464c63f4b6d260d5a81e9448d1e4f
Author: Adrian Cruz
Date:   Fri Jan 31 01:13:18 2014 +0000

    v0.1 - publishing seems to be working properly
```

I made it. I've deployed everything I need to production. It's a barebones application that functions. Any further refinements will be handled with maintenance releases.

```
commit 53554dffeac07a86e9d62791178ca15126877887
Author: Adrian Cruz
Date:   Fri Jan 31 17:21:03 2014 +0000

    v0.1.1 - maintenance release - bad sql query
     - it was late, i was rushing, everything all looked the same
```

I'm technically off on holiday. I checked in on email though almost every half hour. I got an email about a couple of bugs. I was able to send over a solution via email for one bug. I VPN'd in and made some changes for the second.

## Bullet Points

 - Being able to pivot is key
 - User feedback can sometimes get in the way
 - Quick coding is not quality coding

 ## This is NOT The End

I have obviously got some more work to be done. The code base for a fairly well-designed system is now in place. All that remains is to make some fancy user design changes, optimizations, and of course some general code clean up!
