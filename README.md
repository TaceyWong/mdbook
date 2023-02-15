# mdbook

**mdbook** is a utility to create modern online books from Markdown files.

Command-Line-Interface:

```text
NAME:
   mdbook - Creates a book from markdown files

USAGE:
   mdbook [global options] command [command options] [arguments...]

VERSION:
   dev-0.0.1

AUTHOR:
   Tacey Wong <xinyong.wang@qq.com>

COMMANDS:
   init         Creates the boilerplate structure and files for a new book
   build        Builds a book from its markdown files
   test         Tests that a book's Rust code samples compile
   clean        Deletes a built book
   completions  Generate shell completions for your shell to stdout
   watch        Watches a book's files and rebuilds it on changes
   serve        Serves a book at http://localhost:3000, and rebuilds it on changes
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

COPYRIGHT:
   © 2023 TaceyWong under GPL3
```

## Acknowledgements

+ [fsnotify](https://github.com/fsnotify/fsnotify) for file system notifications
+ [cli](https://github.com/urfave/cli) for building command line interface
+ [survey](https://github.com/go-survey/survey) for building interactive and accessible prompts
+ [goldmark](https://github.com/yuin/goldmark) for parsering and rendering Markdown
+ [go-socket.io](https://github.com/googollee/go-socket.io) for updating web page realtime on md-files changes  
+ [semver]https://github.com/Masterminds/semver for parsing semantic versions
+ more in `go.mod`
