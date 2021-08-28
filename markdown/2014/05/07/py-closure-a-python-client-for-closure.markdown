I made a little command line utility to assist in minifying javascript files. It uses [Closure Compiler](https://closure-compiler.appspot.com/home) to compile. I've been using the web app and needed a way to automate the minification process; this command line utility is just the beginning of that process.

## What can I do with it?

Well, if you use Closure Compiler already I'm sure you have always wanted a way to streamline your minification process. This is really just a piece of that. For example, you can probably add this to your build process in Jenkins to minify all javascript files within your code repositiory.

## But, there are a bunch of other asset minifiers I can use!

There are, and I actually was looking into using [Grunt](https://gruntjs.com/). Well, I have been [busy with this new project](/2014/02/01/from-initial-commit-to-launch-the-story-of-a-quick-site-launch.html) and just want to keep the ball rolling. So, I'm sure that I will revisit this minifying process again in due time.

## So, where's the code?

Sure, it's [here on my GitHub](https://github.com/drincruz/Py-Closure). It's short and sweet and super simple. I'm actually going to fork this off and use it similarly at work. Anyhow, I just thought I'd share. Cheers!
