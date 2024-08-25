
# Secret_Password

This program, written in Go, automatically types custom text using a combination of user-defined keybinds.The default keybind is `Ctrl + Shift + F3

Note: I created this for entering master keys because I couldn't remember a 64-character master key. It also served as a learning experience with Go.

## Features

- Secure (using Garble obfuscation)
- Simple to use
- Works anywhere (make sur you start the program as an administrator)

## Environment Variables

If you want to enable control flow obfuscation in Garble, you need to set the following environment variable:

`GARBLE_EXPERIMENTAL_CONTROLFLOW=1`

## Installation

1. Download and install Go (obviously)
2. Download this repository and navigate to the folder
3. Compile the program using this command (remove the `-H=windowsgui` flag if you want a console window):

```bash
garble -literals -seed=random -tiny build -ldflags="-w -s -H=windowsgui -buildid=" -trimpath