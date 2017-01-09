# pocket

<img align="right" src="/pocket.svg" height="166">

[![license](https://img.shields.io/github/license/libeclipse/pocket.svg)](https://raw.githubusercontent.com/libeclipse/pocket/master/LICENSE) [![Build Status](https://travis-ci.org/libeclipse/pocket.svg?branch=master)](https://travis-ci.org/libeclipse/pocket) [![Go Report Card](https://goreportcard.com/badge/github.com/libeclipse/pocket)](https://goreportcard.com/report/github.com/libeclipse/pocket)

***Note: Still in alpha stages. Should not (yet) be used seriously.***

Protect super secret passwords and sketchy snippets - even in the case of your password being leaked.

## Features

***All of these features have not yet been implemented.***

* ***Multi-layer security*** - the password alone isn't enough to compromise your secrets.
* ***Multiple password support*** - you're free to use different passwords for different entries, and no one would ever know that you did.
* ***Deniability*** - *pocket* will randomly add decoy entries so in the event of rubber-hose cryptanalysis, you can claim that some/all the entries are decoys.
* ***Hidden entry identifiers*** - the entry identifiers are hashed so that an attacker cannot even tell what type of data is stored. There have been many cases where users have encrypted their data, but file names have still given them away. In *pocket*, this is mitigated.
* ***Hidden data length*** - every entry is padded to a fixed length so that it is impossible to determine the length of the secret.

## Installation

Simply run:

`~ >> go get github.com/libeclipse/pocket`

This will fetch, compile, and install *pocket* automatically. An added bonus is that it should now be in your PATH so you can call the program from anywhere with a simple:

`~ >> pocket`

## Credits

- [@dotcppfile](https://github.com/dotcppfile) - Came up with ideas and always there to brainstorm and bounce thoughts off. Truly invaluable to the project.
