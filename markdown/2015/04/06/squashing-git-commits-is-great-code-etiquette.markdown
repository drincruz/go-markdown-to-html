Whether you collaborate with a team of folks or just by yourself, having tidy git commits are useful in being able to search through your own git log.

Take this `git log` example:

```
[acruz@wai-laan tmp_blog]$ git log
commit eeb0dae575e2ce9f055625f367c802060838d584
Author: Adrian
Date: Mon Apr 6 15:50:06 2015 -0400

added comments for decorator

commit 12cb1bdb3278c410cc21842c95da1066b3386b21
Author: Adrian
Date: Mon Apr 6 15:49:25 2015 -0400

added working decorator

commit 690b9668ed0a2f205ac0a535c89f089c4e2cca7b
Author: Adrian
Date: Mon Apr 6 15:33:54 2015 -0400

WIP: testing a decorator

commit 34eeee492343aad2936d1b89eff92ee04b818aef
Author: Adrian
Date: Mon Apr 6 15:27:53 2015 -0400

initial commit with README
```

I want to tidy things up before I put in a pull request, because if I were the person reviewing my code, I wouldn't want to have all of those commits to look through. I also want to keep the git log clean so that when myself or any other engineer comes and looks to see the history, there will be a definitive point where I can say, "okay, this is the commit from this pull request".

The basic steps that I typically do are as follows `git rebase -i {COMMIT_TO_REBASE_AFTER}` and then edit the git commit during the interactive rebase.

So, for my example, I want to rebase after my initial commit: `[acruz@wai-laan tmp_blog]$ git rebase -i 34eeee492343aad2936d1b89eff92ee04b818aef`

A menu comes up similar to the following:

```
pick 690b966 WIP: testing a decorator
pick 12cb1bd added working decorator
pick eeb0dae added comments for decorator

# Rebase 34eeee4..eeb0dae onto 34eeee4
#

# Commands:

# p, pick = use commit
# r, reword = use commit, but edit the commit message
# e, edit = use commit, but stop for amending
# s, squash = use commit, but meld into previous commit
# f, fixup = like "squash", but discard this commit's log message
# x, exec = run command (the rest of the line) using shell
#
```

It is hopefully self explanatory. I am going to `reword` one commit and `fixup` the rest so I can have one nice, detailed commit.

```
r 690b966 WIP: testing a decorator
f 12cb1bd added working decorator
f eeb0dae added comments for decorator
```

And now, my `git log` is nice and tidy!

```
[acruz@wai-laan tmp_blog]$ git log
commit 477232a042c308fff5e771757616130603207d85
Author: Adrian
Date: Mon Apr 6 15:33:54 2015 -0400

Added tmp.py with a decorator

decorator has awesome decorator-functionality


commit 34eeee492343aad2936d1b89eff92ee04b818aef
Author: Adrian
Date: Mon Apr 6 15:27:53 2015 -0400

initial commit with README
```
