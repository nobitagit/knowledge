# VScode

:boom: - you should never forget this one

See all open files
<kbd>CTRL</kbd> + <kbd>p</kbd>

Show command palette
<kbd>CMD</kbd> + <kbd>p</kbd>

Surround with text. Put this into your `keybindings.json`

```
  {
    "key": "alt+s",
    "command": "editor.action.insertSnippet",
    "when": "editorTextFocus",
    "args": {
      "snippet": "${1}${TM_SELECTED_TEXT}${1}"
    }
```

And then:
<kbd>OPTION</kbd> + <kbd>s<kbd>

Show suggestions (intellisense style)
<kbd>CTRL</kbd> + <kbd>SPACE</kbd> either on a new line or by highlighting a word

Select whole line
<kbd>CMD</kbd> + <kbd>i</kbd>

Move all highlighted up or down
<kbd>ALT</kbd> + <kbd>ARROW UP/DOWN</kbd>

Open integrated terminal
<kbd>CTRL</kbd> + <kbd>`</kbd>

Focus terminal
<kbd>CMD</kbd> + <kbd>§</kbd>:boom:

Focus tab number 1
<kbd>CMD</kbd> + <kbd>1</kbd>

Show Project Explorer
<kbd>CMD</kbd> + <kbd>SHIFT</kbd> + <kbd>E</kbd>

Go to line
<kbd>CTRL</kbd> + <kbd>g</kbd>
or (VIM style) <kbd>ESC</kbd> + <kbd>:</kbd>

Go to Symbol :boom:
<kbd>CMD</kbd> + <kbd>SHIFT</kbd> + <kbd>o</kbd>
Add a `:` and group them by type

Go to Symbol within workspace :boom:
<kbd>CMD</kbd> + <kbd>t</kbd>

Go to previous cursor location :boom:
<kbd>Ctrl</kbd> + <kbd>-</kbd>

List recent files within the tab :boom:
<kbd>CTRL</kbd> + <kbd>SHIFT</kbd>

## VIM mode only

Delete the char under the cursor
<kbd>x</kbd>

Back to start of line
<kbd>0</kbd>

Back to start of line and insert mode
<kbd>I</kbd>

To end of line and enter edit mode
<kbd>A</kbd>

Back to start of file
<kbd>gg</kbd>

Back to end of file
<kbd>G</kbd>

Find matching parens <kbd>%</kbd>

substitute 'new' for 'old' where g is globally `:s/old/new/g`

jump to definition of method :boom:
<kbd>gd</kbd>

jump to another tab (see more [here](https://stackoverflow.com/a/25254470/1446845))
<kbd>CTRL</kbd>+<kbd>ww</kbd>

**d**elete from cursor until (**t**ill) X character :boom:
<kbd>dtX</kbd>
To expand this combination (to prev and exclusive/inclusive see https://askubuntu.com/a/781012/289123)

**d**elete or **c**hange a block **i**nside (surrounded by) a character. :boom:
For example to delete everything inside parens `(THIS WILL BE REMOVED) => ()`
<kbd>di(</kbd> or
<kbd>ci(</kbd>
See https://leanpub.com/boostyourcodingfuwithvscodeandvim/read#leanpub-auto-editing-like-magic-with-vim-operators for more

## Breadcrumbs

Focus breadcrumbs
<kbd>CMD</kbd> + <kbd>SHIFT</kbd> + <kbd>.</kbd>

Navigate breadcrumbs
<kbd>ALT</kbd> + <kbd>LEFT ARROW</kbd> AND <kbd>ALT</kbd> + <kbd>RIGHT ARROW</kbd>

## Searching

- Focus next result in search
  <kbd>ALT</kbd>+<kbd>u</kbd>

- Focus prev result in search
  <kbd>SHIFT</kbd>+<kbd>ALT</kbd>+<kbd>u</kbd>
