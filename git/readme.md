## Git

## Set author for current repo

```sh
git config user.email "email@example.com"
```

## Get file from other branch

```sh
git checkout branch-name -- file-name.py
```

## What do we really mean by branch?

See [here](https://stackoverflow.com/questions/25068543/what-exactly-do-we-mean-by-branch) and [here](https://stackoverflow.com/a/58020183/1446845).

To help clear up all the confusion, let's start with something more basic, though. When you use Git, there's more than one repository involved. You have _your_ copy, I have _mine_, Fred has _his_, Alice has _hers_, and so on. Every one of these Git repositories has branches and commits and stuff like that. How can we keep them all straight?

There are only a few absolutes, in Git. The main one is those hash IDs you see in your `git log` output:

>     commit 1cfee7bc5c292c09a108e0319ddcec8ab3608887

and:

>     commit a2d1acfe855899e7e9562a16b692aa5d1f44d5dd

> These hash IDs are truly universal. If _you_ have commit `a2d1acfe855899e7e9562a16b692aa5d1f44d5dd`, your Git calls it `a2d1acfe855899e7e9562a16b692aa5d1f44d5dd`. If Alice has it, _her_ Git calls it `a2d1acfe855899e7e9562a16b692aa5d1f44d5dd` too. _Everyone_ who has _that_ commit, calls it that same big ugly hash ID. They either have that, or they don't have anything with that big ugly hash ID.
>
> Mere humans, of course, are terrible at dealing with these hash IDs. But we don't _have_ to deal with them directly, most of the time: we have a _computer_. Why not have _it_ remember the hash IDs? So Git does that—which is where branch names come in—but each Git repository has _its own_ branch names. Your `master` need not match Bob's `master`. Your name `master` will hold some big ugly hash ID, but it might be different from someone else's.
>
> To deal with that, we have our Gits remember each other's branches. When you have your Git call up some other Git that you call `origin`, your Git has their Git list out its branch names and hash IDs. Their Git may say _my `master` is `a2d1acfe855899e7e9562a16b692aa5d1f44d5dd`_. If so, your Git makes sure you have that commit (by hash ID), getting it from their Git if not, and then your Git sets _your_ `refs/remotes/origin/master` _name_ to `a2d1acfe855899e7e9562a16b692aa5d1f44d5dd`.
>
> What this means in the end is that you can tell what their Git's `master` was, the last time your Git talked to their Git, by checking to see what _your_ Git's `origin/master` is now. You can change which hash ID your `master` is; your Git will leave your `origin/master`—full name `refs/remotes/origin/master`—alone, until you have your Git call up their Git again and find out where their `master` is.
>
> Your _reflogs_ remember, for you, what your Git put in your references. Every time your Git updates your `master`, your Git stores a new entry in your Git's `refs/heads/master` reflog. This works for your remote-tracking names too: every time your Git updates your `refs/remotes/origin/master`, your memory of `origin`'s master, your Git stores a new entry in your Git's `refs/remotes/origin/master` reflog. So you _can_ look in your reflogs, but that's to see what you used to have in your own refs. That does not look at any actual _commits_, at least not right away; these reflog entries hold one hash ID each, and you can use that to find commits.
>
> Reflog entries are private to each Git. Your Git will not get reflog entries from their Git. You cannot see what they did to their refs over time, not this way anyway. You can only see what your Git did to _your_ refs over time.

## Other

- [some aliases](https://github.com/Bash-it/bash-it/blob/master/aliases/available/git.aliases.bash)

- You can have a different, perhaps more accurate diff by doing [`git diff --patience`](https://stackoverflow.com/a/36551123/1446845)

- making a few mistakes while typing? [`git config --global help.autocorrect -1`](https://twitter.com/kuizinas/status/1155862466489999362)

- print current branch: `git rev-parse --abbrev-ref HEAD`
