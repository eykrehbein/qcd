<div align="center">
  
  # Quickcd
  
  [![asciicast](https://asciinema.org/a/C8zuyXBaLo2HjTYtw5DZZODiE.svg)](https://asciinema.org/a/C8zuyXBaLo2HjTYtw5DZZODiE)
  
  Terminal Navigation. Faster than ever before.
</div>

## Table of Contents

- [Introduction](#Introduction)
- [Features](#Features)
- [Installation](#Installation)
- [Usage](#Usage)
- [Updating](#Updating)

## Introduction

For everyone who is upset by typing endless directory paths to change directories. Define shortcuts,
with which you can navigate quickly between directories and projects. Using the `qcd back` command, you can easily switch
back to the directory you came from.

## Features

- Create shortcuts for any directory path & navigate to them quickly
- Go back to the directory you came from

## Installation

1. Download and unpack the latest [release](https://github.com/eykrehbein/quickcd/releases) for your specific platform

_By now, only MacOS and Linux are supported_

2. Run the setup

```bash
sudo sh setup.sh
```

3. Restart the terminal

You're done! Use `qcd help` for a list of all available commands

## Usage

It's pretty simple.

Navigation:

- `qcd <name>` Navigate to this QuickLink
- `qcd back` Get back to the directory you came from

Managing QuickLinks:

- `qcd help|h` A list of all commands
- `qcd add <name> <relative_or_absolute_path>` Add a QuickLink (shortcut)
- `qcd remove|rm <name>` Remove a QuickLink
- `qcd list` List all of your QuickLinks

## Updating

Follow these steps to update to a newer release:

1. Download and unpack the latest [release](https://github.com/eykrehbein/quickcd/releases) for your specific platform

2. Run `sudo sh update_dev.sh`

You're done! Use `qcd help` for a list of all available commands
