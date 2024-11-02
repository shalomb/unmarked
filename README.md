**unmarked 🎯**
---

Similar to [`harpoon`](https://github.com/ThePrimeagen/harpoon), unmarked is the keyboard user's tool for switching to windows
just using their marks.

```shell
unmarked mark f     # Give the currently active window a mark of f
# Move around to other windows in the desktop environment, etc
unmarked summon f   # Switch back to and focus the window with mark f
```

[yabai](https://github.com/koekeishiya/yabai) and [skhd](https://github.com/koekeishiya/skhd) are required to complete functionality.

With a `~/.config/skhd/skhdrc` file as follows

```skhd
ctrl + alt - a : ~/.bin/unmarked summon a
ctrl + alt - b : ~/.bin/unmarked summon b
ctrl + alt - c : ~/.bin/unmarked summon c
...
ctrl + alt - x : ~/.bin/unmarked summon x
ctrl + alt - y : ~/.bin/unmarked summon y
ctrl + alt - z : ~/.bin/unmarked summon z

ctrl + alt + cmd - a : ~/.bin/unmarked mark a
ctrl + alt + cmd - b : ~/.bin/unmarked mark b
ctrl + alt + cmd - c : ~/.bin/unmarked mark c
...
ctrl + alt + cmd - x : ~/.bin/unmarked mark x
ctrl + alt + cmd - y : ~/.bin/unmarked mark y
ctrl + alt + cmd - z : ~/.bin/unmarked mark z
```
You can press `ctrl-alt-cmd-t` when over a wezterm window to give it the mark
`t`, press `ctrl-alt-cmd-f` when over a firefox window to mark it with `f`,
etc, etc.

At any point later, press `ctrl-alt-t` to raise/focus the wezterm
terminal, press `ctrl-alt-f` to focus firefox, etc.

No need to `alt-tab` or reach for the mouse - Win! 🏆

**why? 💡**
---

My workflow usually involves making some code edits in neovim in wezterm,
switching to firefox to test, moving to jira to making some comments, moving
to teams to make an annoucement, moving back to neovim, etc. `alt-tabbing` my
way through these is a tad bit tedious.

My working set is usually 2-3 windows - and I want these to be quick to summon at the
speed of thought.
With unmarked, simply pressing `ctrl-alt-<mnemonic>` is enough to get me back
into the app and back on track.

**Building 🛠️**
---

Requires `go` >= 1.19, yabai, skhd

```shell
make build
cp unmarked-darwin ~/.bin/unmarked
export PATH="$HOME/.bin:$PATH"

unmarked help
```
**Debugging 🐞**
---

Ensure `unmarked` is installed into a directory of `$PATH
`
```shell
unmarked version
```

Run `skhd` in debugging mode and test keyboard input

```shell
pkill skhd
skhd -V

# when complete with debugging, restart the skhd service
skhd --start-service
```

Refer to [`skhd`'s documentation](https://github.com/koekeishiya/skhd/issues/1) on how to discover keycodes.
