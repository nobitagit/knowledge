## VIM

### How to add a character at the beginning of multiple lines
Useful for example to comment a big chunk of code

- https://stackoverflow.com/a/253391/1446845

### How to save a file without having permissions to do it
If you open a file as your user (not using sudo), and after having edited you want to save it anyway:

```sh
:w !sudo tee %
```

Where:
- :w – Write a file (actually buffer).
- !sudo – Call shell with sudo command.
- tee – The output of write (vim :w) command redirected using tee.
- % – The % is nothing but current file name. In this example, it is /etc/apache2/conf.d/mediawiki.conf. In other words tee command is run as root and it takes standard input (or the buffer) and write it to a file represented by %. However, this will prompt to reload file again (hit L to load changes in vim itself):
