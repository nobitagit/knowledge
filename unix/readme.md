## Unix

- You can search `man` pages, since they're generally displayed using `less`. Use `/` and search for term + enter. `n` to go next, `N` to go back.

- Easily cut a string.

```sh
my_var="1234my string"
${my_var:4} # "my string"
${My_var:4:2} # "my"
```
