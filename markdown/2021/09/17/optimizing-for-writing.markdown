Welcome back to ... _me_! I haven't written a blog post in over a year and I have no excuses. To be honest, I got stuck in the mindset of "how can I create the _coolest_ personal, website architecture ever?". But let's be honest here, if my main objective here is to write personal content, why not just optimize for writing the way I am most comfortable with?

## HTML gets the job done

So, let's take a step back and think about the web; and more importantly, let's focus on consuming content on the web. As of now, I have no need for dynamic content. The basic workflow for me is to write up some content, it gets published, and that's it. There's no interactive component that needs custom coding. There's no dynamic content that gets built with AI, data pipelines, nor custom algorithms. What this site is now is exactly its main purpose: *serve up HTML to deliver content for reading*.

## Why the change?

Maybe you happen to remember what this site looked like before its current facelift, but if not, well it worked! It served its purpose of also just delivering content. But it also came with an expense of maintenance: Python, Django, Postgresql, compute instance fees, etc.
So a good friend of mine asked if I _really_ needed all of this setup just for my own personal use. I mean, sure there's a little bit of bragging rights into hosting your own systems, but there's also over-engineering. I thought about it and agreed, I needed to go back to basics. Bye bye, old dynamic site! Hello static HTML site!

## There is still _some_ engineering going on

While this site is just another static site, it does have some neat engineering work that I do need to still support. For the record, I am more than happy to support this. :)

### Markdown is already in my day-to-day

If you know how I operate, I take notes for _everything_. No, really, *everything*. But to optimize how I can write notes, I picked up a habit of writing notes in markdown. Markdown is nice for writing notes, especially as an engineer, because you can get markdown renderers to display beautiful content of your notes very easily.
For example, in my work notes, I may write `SomeSortOfRelevantCode.forMeToRemember()` with details on its behaviour, just so I can reference it in the future.
Or if I am looking into a defect and want to write down several steps to reproduce something, I may make a code block.

```
$ mkdir -p markdown/2021/09/21
$ touch optimizing-for-writin.markdown
$ # ... etc.
```

The use of markdown is useful for a lot of my note taking and _now_ my blogging.

### Converting markdown to HTML

The funner bits I worked on were writing simple tools to convert markdown to HTML. This ended up being a project I called, [go-markdown-to-html](https://github.com/drincruz/go-markdown-to-html). It is exactly as it sounds, it is a project written in [Golang](https://golang.org/), it takes markdown as its source, and it writes HTML files. Simple, yet effective.

## What now?

As of now, `go-markdown-to-html` is honestly still very much a work in progress (WIP). But like most things, I wanted to prioritize my main goal of *getting back to writing*.
I plan on engineering `go-markdown-to-html` a bit more, but also wanted to make sure that maintenance wasn't taking up most of my time.
The fun thing about running a static site now is that if I _really_ wanted to just write HTML and publish those files, I can totally do so! But as of now, my process to publish is pretty streamlined for markdown to HTML and I am content.
