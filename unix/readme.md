## Unix

- You can search `man` pages, since they're generally displayed using `less`. Use `/` and search for term + enter. `n` to go next, `N` to go back.

- Easily cut a string.

```sh
my_var="1234my string"
${my_var:4} # "my string"
${my_var:4:2} # "my"
```

- See all colors in bash `for (( i = 0; i < 17; i++ )); do echo "$(tput setaf $i)This is ($i) $(tput sgr0)"; done`

### Searching/finding stuff

To order `ls` by modified date (asc/desc):
```
ls -t
ls -tr 
```
