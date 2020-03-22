## Git

## Set author for current repo

```sh
git config user.email "email@example.com"
```

## Get file from other branch

```sh
git checkout branch-name -- file-name.py
```

## Other

- [some aliases](https://github.com/Bash-it/bash-it/blob/master/aliases/available/git.aliases.bash)

- You can have a different, perhaps more accurate diff by doing [`git diff --patience`](https://stackoverflow.com/a/36551123/1446845)

- making a few mistakes while typing? [`git config --global help.autocorrect -1`](https://twitter.com/kuizinas/status/1155862466489999362)

- print current branch: `git rev-parse --abbrev-ref HEAD`
