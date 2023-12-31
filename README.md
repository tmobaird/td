# TD (Todo CLI)

This is a fun project I am working on to help learn Golang.
Feature ideas and feedback on the Golang code is welcome 🤝.

### Installation

```bash
export TD_DOWNLOAD_VERSION="0.0.1"
export TD_DOWNLOAD_URL="https://github.com/tmobaird/td/releases/download/$TD_DOWNLOAD_VERSION/td.tar.gz"
export BINARIES_PATH="$HOME/bin"
mkdir -p /tmp/td-download
wget -O - "$TD_DOWNLOAD_URL" | tar -xz -C /tmp/td-download
mkdir -p $BINARIES_PATH
mv /tmp/td-download/build/td $BINARIES_PATH
chmod +x "$BINARIES_PATH/td"
```

### Usage

TD is a CLI allows for local todo list management. The interface
documentation can be seen below.

```
Usage: td [options] [command] [arguments]
Options:
  -h, --help     Print usage
  -v, --verbose  Print verbose output
Commands:
  a,  add <name>                      Add a new todo
  ls, list                            List all todos
  d,  delete <index|uuid>             Delete a todo
  do, done   <index|uuid>             Mark a todo as done
  un, undo   <index|uuid>             Mark a todo as not done
  e,  edit   <index|uuid> <new name>  Edit a todo
  r,  rank   <index|uuid> <new rank>  Rerank a todo
```

### Roadmap

- [x] Add tests for commands
- [x] Add reporter interface to print output to console
- Flags (https://pkg.go.dev/flag@go1.20.6)
    - [x] Support verbose list
    - [x] Help command
    - [ ] Test flags interface
- [x] Support marking as completed
- [x] Support marking as not completed
- [x] Support editing of todos
- [x] Support batch delete todos
- [x] Support batch add todos
- [x] Perform actions on todos by index or uuid (edit, delete, done, undo)
    - [x] delete
    - [x] edit
    - [x] done
    - [x] undo
- [x] Command shorthands (a for add, l for list, d for delete)
- [x] Prioritize items in list
- [x] Catch errors when not enough inputs for command
- [x] Change persistence directory to ~/.td
- [x] Release 1
- [ ] Restructure project directories
- [ ] Persistent configurations
    - [x] Dont show completed
    - [x] Default config does not show completed
    - [ ] Multiple todo lists
    - [ ] Switch ranking to be done through UUID so that indexes can be remove completed indexes
- [ ] Subtodos
