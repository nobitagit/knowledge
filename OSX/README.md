## OSX

- You can tell Spotlight to ignore all the content of a directory just by inserting a file named `.metadata_never_index`. All the content from that directory and all the subdirectories will be ignored. This is very useful to ignore generated files (see `node_modules`)

- To convert MOV files to mp4 `ffmpeg -i ~/in.mov ~/out.mp4`. [>](https://apple.stackexchange.com/a/276557/44487)

- To easily screen capture in various formats use [Kap](https://getkap.co/).

- A browser in terminal https://www.brow.sh/docs/installation/

- To delete a site from Brave/Chrome autocomplete, type it, highlight it and press `SHIFT + FN + BACKSPACE`.

- To kill and restart the touchbar process when it's stuck

```
sudo pkill TouchBarServer;
sudo killall “ControlStrip”;
```
