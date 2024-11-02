**unmarked 🎯**
---

Similar to [`harpoon`](https://github.com/ThePrimeagen/harpoon), unmarked
is the keyboard user's tool for switching desktop windows using just their
marks.

If you are familiar with [vim/neovim's concept of marks](https://vim.fandom.com/wiki/Using_marks#Setting_marks) - unmarked does
the same for desktop windows.

```shell
unmarked mark f     # Mark the currently active window with the letter 'f'
# Move around to other windows in the desktop environment, etc
unmarked summon f   # Switch back to and focus the window marked 'f'
```

[yabai](https://github.com/koekeishiya/yabai) and [skhd](https://github.com/koekeishiya/skhd) are required to complete functionality. Works only on MacOS currently.

**Setup ⚙️**
---

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
You are free to use any letter now to mark (and jump between) windows.

Let's say you use wezterm a lot in your workflow and want to mark it - you
would press `ctrl-alt-cmd-t` to mark it with the letter `t`. (`t` being
a mnemonic for terminal - but you would choose any letter of your liking).

Now, let's say you've switched windows and are doing something else and want
to move back to the wezterm window quickly - simply press `ctrl-alt-t`. Voila!

No need to `alt-tab` or reach for the mouse - Win! 🏆

**why? 💡**
---

Most developers' workflow usually involves making some code edits in the
terminal, switching to a browser to test, moving to some custom app to making
some comments, moving to slack to make an announcement, moving back to the
terminal to pick up coding work, etc.

`alt-tabbing` your way through many open windows is a tad bit tedious that the
tab key starts to develop a shine. For the few windows that make up the
current context, it should be super quick to switch to/between them and hence
the `ctrl-alt-<mnemonic>` to keep you in flow state.

**Building 🛠️**
---

Requires `go` >= 1.19, yabai, skhd

```shell
make build
cp unmarked-darwin* ~/.bin/unmarked  # or some other dir in $PATH
export PATH="$HOME/.bin:$PATH"

unmarked help  # Testing installation
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
