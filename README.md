# addToPath

A command-line tool for easily adding paths to your Linux user's `.bashrc` file.

## Overview

This CLI tool simplifies the process of managing environment paths in your Linux user account by modifying the `.bashrc` file. It allows users to append new paths quickly and safely, creating a backup of the original `.bashrc` before making any changes. If needed, the tool also provides functionality to reverse the last change, restoring the previous state from the backup. This utility is perfect for developers and users who frequently adjust their environment settings for various applications.

## Features

- **Add Paths**: Append a new path to your `.bashrc` file.
- **Backup**: Automatically create a backup of your `.bashrc` file before making changes.
- **Undo Changes**: Revert the last change if needed.

## Installation

You can install `addToPath` using the `go get` command:

```bash
go install github.com/SHOOTDEAD/addToPath@latest
```

## Usage

### Add a Path
To add a path to your .bashrc, use the following command:

```bash
addToPath --path /your/new/path
```
Reverse the Last Change
To revert the last change made to your .bashrc, use:

```bash
addToPath --reverse
```

## Flags
--path, -p: Specify the path you want to add to your .bashrc.
--reverse, -r: Revert the last change made to .bashrc.

## Example
To add /usr/local/bin to your .bashrc, you would run:

```bash
addToPath --path /usr/local/bin
```

If you want to undo that change, you would run:

```bash
addToPath --reverse
```
