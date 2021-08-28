So, I was asked how my Twitter and Google+ sidebar works. Nothing, much to it to be honest. But alas, here you have it along with the source that can be found on my [github](https://github.com/drincruz/feedmesocial).

## Static, Yet Beautiful

Yes, it writes out to a static file and then I load it into the sidebar. _Why_? Well, I have no need for it to be real time. I do not post that often that I would benefit from making a real time feed.

It's also a speed decision. I have no desire to have a page load dependency on something I can host myself. There are some things that I don't mind loading from third party sites (like jQuery). But, if I can keep that to a minimum, then great!

## Make a few edits and go!

`FeedMeSocialService.java`, `GooglePlusFeedEntry.java`, and `TwitterTokenFactory.java` are the only files you'll need to edit. --of course, if you want to make edits elsewhere, then fork away!

To get the Twitter information, go to [dev.twitter.com](https://dev.twitter.com) and create an application.

To get the Google+ information, go to [cloud.google.com/console](https://cloud.google.com/console) and create a project with a service account. You will be able to download your service key as well as get the email address for the service account.
