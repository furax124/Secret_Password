
# Secret_Password

This program, written in Go, automatically types custom text using a combination of user-defined keybinds.The default keybind is `Ctrl + Shift + F3

Note: I created this for entering master keys keepass because I couldn't remember a 64-character master key. It also served as a learning experience with Go.

## Features

- Secure (using Garble obfuscation)
- Simple to use
- Works anywhere (make sur you start the program as an administrator)

## Environment Variables

If you want to enable control flow obfuscation in Garble, you need to set the following environment variable:

`GARBLE_EXPERIMENTAL_CONTROLFLOW=1`

## Installation

1. Download and install Go 1.20.2 or later (obviously)
2. Download this repository and navigate to the folder
3. Run Start.bat

## Todo

- [ ] Add Trustedinstaller option (to make it work in secure dekstop or anywhere)
- [x] Add xor Encryption
- [x] Builder
