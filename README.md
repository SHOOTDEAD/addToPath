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

## Prerequisites
Make sure you have Go version 1.23.2 or later installed. If you are using an older version of Go, please upgrade to the recommended version to ensure compatibility.

## Building from Source
If you'd like to clone the repository and compile the binary yourself, follow these steps:

### Clone the repository:

```bash
git clone https://github.com/SHOOTDEAD/addToPath.git
cd addToPath
```
### Build the binary:

```bash
go build
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
