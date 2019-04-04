## Nginx

### Does Nginx run processes as root?

No, it does start as root, but spawns child processes as a non root user (the one you have configured a user in the nginx.conf, generally it's www-data).

You can test this by starting nginx and then `ps aux | grep nginx`. You will see something like:

```
root     29757  0.0  0.1 140628  1504 ?        Ss   11:21   0:00 nginx: master process /usr/sbin/nginx -g daemon on; master_process on;
www-data 29759  0.0  0.6 143300  6188 ?        S    11:21   0:00 nginx: worker process
```

