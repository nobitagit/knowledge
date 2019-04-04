## Nginx

### Does Nginx run processes as root?

No, it does start as root, but spawns child processes as a non root user (the one you have configured a user in the nginx.conf, generally it's www-data).

