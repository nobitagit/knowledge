## Unix

### How to list all active connections

```sh
sudo netstat -plunt
```

-p, --program Show the PID and name of the program to which each socket belongs.
-l, --listening Show only listening sockets.
-u UDP
-n numeric - instead of names, like ssh, http
-t TCP

## git

- rebasing against remote branches is faster than keeping a local up to date copy. Just `git pull --rebase origin master`. See https://twitter.com/hybridcattt/status/1114596481359581185?s=19
