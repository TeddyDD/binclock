# Binclock

[![HitCount](http://hits.dwyl.com/TeddyDD/binclock.svg)](http://hits.dwyl.com/TeddyDD/binclock) 
![License](https://img.shields.io/github/license/teddydd/binclock?style=flat-square)
![Release](https://img.shields.io/github/v/release/teddydd/binclock?style=flat-square)

Tiny binary clock for your terminal. Written in Go with [tcell] library.

![screenshot](screen.png)

## Installation

```
go get -u github.com/TeddyDD/binclock
```

or download [binary release](https://github.com/TeddyDD/binclock/releases)
for your OS

# Usage

Run `binclock` in terminal. <kbd>Esc</kbd>, <kbd>q</kbd> or <kbd>CTRL</kbd>
<kbd>C</kbd> to quit. You can set characters used by clock to display on
and off bits with `-o` and `-z` flags. `-r` flag reverse colors of active
and inactive bits (useful for white color schemes)

[tcell]: https://github.com/gdamore/tcell
