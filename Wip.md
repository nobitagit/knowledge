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

## JavaScript

### Modules

commonjs:
Implemented by node
Used for the server side when you have modules installed
No runtime/async module loading
import via “require”
export via “module.exports”
When you import you get back an object
No tree shaking, because when you import you get an object
No static analyzing, as you get an object, so property lookup is at runtime
You always get a copy of an object, so no live changes in the module itself
Poor cyclic dependency management
Simple Syntax

ECMAScript Harmony (ES6)

Used for both server/client side
Runtime/static loading of modules supported
When you import, you get back bindings value (actual value)
Import via “import” and export via “export”
Static analyzing — You can determine imports and exports at compile time (statically) — you only have to look at the source code, you don’t have to execute it
Tree shakeable, because of static analyzing supported by ES6
Always get an actual value so live changes in the module itself
Better cyclic dependency management than CommonJS

https://medium.freecodecamp.org/anatomy-of-js-module-systems-and-building-libraries-fadcd8dbd0e
